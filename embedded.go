package passit

import (
	_ "embed" // for go:embed
	"io"
	"strings"
	"sync"
)

//go:generate go run emoji_generate.go emoji_generate_gen.go emoji_generate_ucd.go -unicode 13.0.0

type embeddedGenerator struct {
	once sync.Once
	raw  string
	list []string
}

func (eg *embeddedGenerator) Password(r io.Reader) (string, error) {
	eg.once.Do(func() {
		eg.list = strings.Split(eg.raw, "\n")
	})
	return readSliceN(r, eg.list)
}

var (
	// This wordlist was taken from:
	// https://www.eff.org/files/2016/07/18/eff_large_wordlist.txt.
	//
	// eff_large_wordlist.txt is licensed by the Electronic Frontier Foundation
	// under a CC BY 3.0 US license
	// (https://creativecommons.org/licenses/by/3.0/us/).
	//
	//go:embed eff_large_wordlist.txt
	effLargeWordlist string

	// This wordlist was taken from:
	// https://www.eff.org/files/2016/09/08/eff_short_wordlist_1.txt.
	//
	// eff_short_wordlist_1.txt is licensed by the Electronic Frontier Foundation
	// under a CC BY 3.0 US license
	// (https://creativecommons.org/licenses/by/3.0/us/).
	//
	//go:embed eff_short_wordlist_1.txt
	effShortWordlist1 string

	// This wordlist was taken from:
	// https://www.eff.org/files/2016/09/08/eff_short_wordlist_2_0.txt.
	//
	// eff_short_wordlist_2_0.txt is licensed by the Electronic Frontier Foundation
	// under a CC BY 3.0 US license
	// (https://creativecommons.org/licenses/by/3.0/us/).
	//
	//go:embed eff_short_wordlist_2_0.txt
	effShortWordlist2 string

	//go:embed emoji_13.0.txt
	emoji13List string
)

// EFFLargeWordlist is a Generator that returns a random word from the
// EFF Large Wordlist for Passphrases (eff_large_wordlist.txt).
var EFFLargeWordlist Generator = &embeddedGenerator{raw: effLargeWordlist}

// EFFShortWordlist1 is a Generator that returns a random word from the
// EFF Short Wordlist for Passphrases #1 (eff_short_wordlist_1.txt).
var EFFShortWordlist1 Generator = &embeddedGenerator{raw: effShortWordlist1}

// EFFShortWordlist2 is a Generator that returns a random word from the
// EFF Short Wordlist for Passphrases #2 (eff_short_wordlist_2_0.txt).
var EFFShortWordlist2 Generator = &embeddedGenerator{raw: effShortWordlist2}

// Emoji13 is a Generator that returns a random emoji from the Unicode 13.0 emoji
// list.
var Emoji13 Generator = &embeddedGenerator{raw: emoji13List}
