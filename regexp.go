package password

import (
	"fmt"
	"io"
	"regexp/syntax"
	"strings"
	"sync"
	"unicode"
	"unicode/utf8"

	"golang.org/x/text/unicode/rangetable"
)

const maxUnboundedRepeatCount = 15

type regexpGenerator func(*strings.Builder, io.Reader) error

type RegexpParser struct {
	unicodeAny      bool
	specialCaptures map[string]func(*syntax.Regexp) (Template, error)
}

func ParseRegexp(pattern string, flags syntax.Flags) (Template, error) {
	return new(RegexpParser).Parse(pattern, flags)
}

func (p *RegexpParser) SetUnicodeAny() { p.unicodeAny = true }

func (p *RegexpParser) SetSpecialCapture(name string, factory func(*syntax.Regexp) (Template, error)) {
	if p.specialCaptures == nil {
		p.specialCaptures = make(map[string]func(*syntax.Regexp) (Template, error))
	}

	p.specialCaptures[name] = factory
}

type regexpTemplate struct{ gen regexpGenerator }

func (p *RegexpParser) Parse(pattern string, flags syntax.Flags) (Template, error) {
	// We intentionally never generate newlines, but passing syntax.MatchNL to
	// syntax.Parse simplifies the parsed character classes.
	flags |= syntax.MatchNL

	// The generator acts badly when used with syntax.FoldCase. Zero it out.
	flags &= ^syntax.FoldCase

	r, err := syntax.Parse(pattern, flags)
	if err != nil {
		return nil, err
	}

	// Simplify is great if one is trying to match the regexp, but less so when
	// generating a string from a regexp. Simplify transforms the regexp in ways
	// that might not be obvious. Because the output is dependent on the
	// specific stream read from the io.Reader, some of these transformations
	// are less than ideal. It's also possible that the simplifications
	// performed could differ between go releases.
	//
	// For example x{3,6} is transformed to xxx(?:x(?:xx?)?)?. This requires
	// more reads from the underlying io.Reader and is more work for the
	// generator.
	//
	// TODO(tmthrgd): It may be worth adding some simplifications of our own.
	//
	// r = r.Simplify()

	gen, err := p.parse(r)
	if err != nil {
		return nil, err
	}

	return &regexpTemplate{gen}, nil
}

func (rt *regexpTemplate) Password(r io.Reader) (string, error) {
	var b strings.Builder
	if err := rt.gen(&b, r); err != nil {
		return "", err
	}

	return b.String(), nil
}

func (p *RegexpParser) parse(r *syntax.Regexp) (regexpGenerator, error) {
	switch r.Op {
	case syntax.OpEmptyMatch:
		return p.noop(r)
	case syntax.OpLiteral:
		return p.literal(r)
	case syntax.OpCharClass:
		return p.charClass(r)
	case syntax.OpAnyCharNotNL, syntax.OpAnyChar:
		return p.anyCharNotNL(r)
	case syntax.OpBeginLine, syntax.OpEndLine,
		syntax.OpBeginText, syntax.OpEndText,
		syntax.OpWordBoundary, syntax.OpNoWordBoundary:
		return p.noop(r)
	case syntax.OpCapture:
		return p.capture(r)
	case syntax.OpStar:
		return p.star(r)
	case syntax.OpPlus:
		return p.plus(r)
	case syntax.OpQuest:
		return p.quest(r)
	case syntax.OpRepeat:
		return p.repeat(r)
	case syntax.OpConcat:
		return p.concat(r)
	case syntax.OpAlternate:
		return p.alternate(r)
	default:
		return nil, fmt.Errorf("strongroom/password: invalid regexp %q, unhandled op %s", r, r.Op)
	}
}

func (p *RegexpParser) parseMany(rs []*syntax.Regexp) ([]regexpGenerator, error) {
	gens := make([]regexpGenerator, len(rs))
	for i, r := range rs {
		gen, err := p.parse(r)
		if err != nil {
			return nil, err
		}

		gens[i] = gen
	}

	return gens, nil
}

func (*RegexpParser) noop(*syntax.Regexp) (regexpGenerator, error) {
	return func(*strings.Builder, io.Reader) error {
		return nil
	}, nil
}

func (*RegexpParser) literal(sr *syntax.Regexp) (regexpGenerator, error) {
	s := string(sr.Rune)
	if idx := strings.IndexFunc(s, notAllowed); idx >= 0 {
		r, _ := utf8.DecodeRuneInString(s[idx:])
		return nil, fmt.Errorf("strongroom/password: regexp literal contains prohibitted rune %U", r)
	}

	return func(b *strings.Builder, r io.Reader) error {
		b.WriteString(s)
		return nil
	}, nil
}

func (p *RegexpParser) charClass(sr *syntax.Regexp) (regexpGenerator, error) {
	const maxR16 = 1<<16 - 1

	tab := new(unicode.RangeTable)
	for i := 0; i < len(sr.Rune); i += 2 {
		start, end := sr.Rune[i], sr.Rune[i+1]

		if start > maxR16 {
			tab.R32 = append(tab.R32, unicode.Range32{
				Lo:     uint32(start),
				Hi:     uint32(end),
				Stride: 1,
			})
			continue
		}

		if end > maxR16 {
			tab.R32 = append(tab.R32, unicode.Range32{
				Lo:     maxR16 + 1,
				Hi:     uint32(end),
				Stride: 1,
			})
			end = maxR16
		}

		tab.R16 = append(tab.R16, unicode.Range16{
			Lo:     uint16(start),
			Hi:     uint16(end),
			Stride: 1,
		})

		if end <= unicode.MaxLatin1 {
			tab.LatinOffset++
		}
	}

	tab = intersectRangeTables(tab, p.anyRangeTable())
	return p.charClassInternal(sr, tab)
}

func (p *RegexpParser) anyCharNotNL(sr *syntax.Regexp) (regexpGenerator, error) {
	return p.charClassInternal(sr, p.anyRangeTable())
}

func (*RegexpParser) charClassInternal(sr *syntax.Regexp, tab *unicode.RangeTable) (regexpGenerator, error) {
	count := countTableRunes(tab)
	if count == 0 {
		return nil, fmt.Errorf("strongroom/password: character class %s contains zero allowed runes", sr)
	}

	return func(b *strings.Builder, r io.Reader) error {
		v, err := readRune(r, tab, count)
		b.WriteRune(v)
		return err
	}, nil
}

func (p *RegexpParser) capture(sr *syntax.Regexp) (regexpGenerator, error) {
	if sr.Name != "" {
		return p.namedCapture(sr)
	}

	return p.parse(sr.Sub[0])
}

func (p *RegexpParser) namedCapture(sr *syntax.Regexp) (regexpGenerator, error) {
	factory, ok := p.specialCaptures[sr.Name]
	if !ok {
		return p.parse(sr.Sub[0])
	}

	tmpl, err := factory(sr)
	if err != nil {
		return nil, err
	}

	return func(b *strings.Builder, r io.Reader) error {
		pass, err := tmpl.Password(r)
		b.WriteString(pass)
		return err
	}, nil
}

func (p *RegexpParser) star(sr *syntax.Regexp) (regexpGenerator, error) {
	return p.repeatInternal(sr, 0, maxUnboundedRepeatCount)
}

func (p *RegexpParser) plus(sr *syntax.Regexp) (regexpGenerator, error) {
	// We use maxUnboundedRepeatCount+1 here so that x{1,} and x+ are identical,
	// x{0,} and x* are already identical.
	return p.repeatInternal(sr, 1, maxUnboundedRepeatCount+1)
}

func (p *RegexpParser) quest(sr *syntax.Regexp) (regexpGenerator, error) {
	return p.repeatInternal(sr, 0, 1)
}

func (p *RegexpParser) repeat(sr *syntax.Regexp) (regexpGenerator, error) {
	max := sr.Max
	if max == -1 {
		max = sr.Min + maxUnboundedRepeatCount
	}

	return p.repeatInternal(sr, sr.Min, max)
}

func (p *RegexpParser) repeatInternal(sr *syntax.Regexp, min, max int) (regexpGenerator, error) {
	gen, err := p.parse(sr.Sub[0])
	if err != nil {
		return nil, err
	}

	return func(b *strings.Builder, r io.Reader) error {
		n, err := readUint32n(r, uint32(max-min+1))
		if err != nil {
			return err
		}

		for n += uint32(min); n > 0; n-- {
			if err := gen(b, r); err != nil {
				return err
			}
		}

		return nil
	}, nil
}

func (p *RegexpParser) concat(sr *syntax.Regexp) (regexpGenerator, error) {
	gens, err := p.parseMany(sr.Sub)
	if err != nil {
		return nil, err
	}

	return func(b *strings.Builder, r io.Reader) error {
		for _, gen := range gens {
			if err := gen(b, r); err != nil {
				return err
			}
		}

		return nil
	}, nil
}

func (p *RegexpParser) alternate(sr *syntax.Regexp) (regexpGenerator, error) {
	gens, err := p.parseMany(sr.Sub)
	if err != nil {
		return nil, err
	}

	return func(b *strings.Builder, r io.Reader) error {
		idx, err := readUint32n(r, uint32(len(gens)))
		if err != nil {
			return err
		}

		return gens[idx](b, r)
	}, nil
}

var regexpAnyRangeTableUni struct {
	tab *unicode.RangeTable
	sync.Once
}

func (p *RegexpParser) anyRangeTable() *unicode.RangeTable {
	if !p.unicodeAny {
		return rangeTableASCII
	}

	regexpAnyRangeTableUni.Do(func() {
		regexpAnyRangeTableUni.tab = rangetable.Merge(allowedRanges...)
	})
	return regexpAnyRangeTableUni.tab
}

func intersectRangeTables(a, b *unicode.RangeTable) *unicode.RangeTable {
	var rt unicode.RangeTable

	/*if r0.Stride != 1 {
		panic("strongroom/password: unicode.RangeTable has entry with Stride > 1")
	}*/

	for _, r0 := range a.R16 {
		for _, r1 := range b.R16 {
			if r1.Lo > r0.Hi {
				break
			} else if r0.Lo > r1.Hi {
				continue
			}

			lo, hi := intersection(rune(r0.Lo), rune(r0.Hi), rune(r1.Lo), rune(r1.Hi), rune(r1.Stride))
			if lo > hi {
				continue
			}

			if hi <= unicode.MaxLatin1 {
				rt.LatinOffset++
			}

			rt.R16 = append(rt.R16, unicode.Range16{Lo: uint16(lo), Hi: uint16(hi), Stride: r1.Stride})
		}
	}

	for _, r0 := range a.R32 {
		for _, r1 := range b.R32 {
			if r1.Lo > r0.Hi {
				break
			} else if r0.Lo > r1.Hi {
				continue
			}

			lo, hi := intersection(rune(r0.Lo), rune(r0.Hi), rune(r1.Lo), rune(r1.Hi), rune(r1.Stride))
			if lo > hi {
				continue
			}

			rt.R32 = append(rt.R32, unicode.Range32{Lo: uint32(lo), Hi: uint32(hi), Stride: r1.Stride})
		}
	}

	return &rt
}

func intersection(lo0, hi0, lo1, hi1, stride1 rune) (lo, hi rune) {
	lo, hi = lo1, hi1

	if stride1 == 1 {
		if lo < lo0 {
			lo = lo0
		}
		if hi > hi0 {
			hi = hi0
		}
	} else {
		if lo < lo0 {
			c := lo0 - lo1
			c += stride1 - 1
			c -= c % stride1
			lo += c
		}
		if hi > hi0 {
			c := hi1 - hi0
			c += stride1 - 1
			c -= c % stride1
			hi -= c
		}
	}

	return lo, hi
}
