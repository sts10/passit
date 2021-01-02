package password

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"sync"
	"unicode"
	"unicode/utf8"
)

type words struct {
	list  []string
	count int
}

// FromWords returns a Template factory that generates passwords of count words
// length by joining random words from list. It returns an error if the list of
// words is invalid.
func FromWords(list ...string) (func(count int) Template, error) {
	if len(list) < 2 {
		return nil, errors.New("password: list too short")
	} else if len(list) > maxUint32 {
		return nil, errors.New("password: list too long")
	}

	seen := make(map[string]struct{}, len(list))
	for _, word := range list {
		if len(word) < 1 {
			return nil, errors.New("password: empty word in list")
		} else if !utf8.ValidString(word) {
			return nil, errors.New("password: word contains invalid unicode rune")
		} else if idx := strings.IndexFunc(word, notAllowed); idx >= 0 {
			r, _ := utf8.DecodeRuneInString(word[idx:])
			return nil, fmt.Errorf("password: word contains prohibited rune %U", r)
		} else if strings.IndexFunc(word, unicode.IsSpace) >= 0 {
			return nil, errors.New("password: word contains space")
		}

		if _, dup := seen[word]; dup {
			return nil, errors.New("password: list contains duplicate word")
		}
		seen[word] = struct{}{}
	}

	list = append([]string(nil), list...)
	return func(count int) Template { return &words{list, count} }, nil
}

func (w *words) Password(r io.Reader) (string, error) {
	if w.count <= 0 {
		return "", errors.New("password: count must be greater than zero")
	}

	words := make([]string, w.count)
	for i := range words {
		idx, err := readUint32n(r, uint32(len(w.list)))
		if err != nil {
			return "", err
		}

		words[i] = w.list[idx]
	}

	return strings.Join(words, " "), nil
}

var effLargeWordlistVal struct {
	sync.Once
	list []string
}

// EFFLargeWordlist returns a Template that generates passwords of count words
// length by joining random words from the EFF Large Wordlist for Passphrases
// (eff_large_wordlist.txt).
func EFFLargeWordlist(count int) Template {
	effLargeWordlistVal.Do(func() {
		effLargeWordlistVal.list = strings.Split(effLargeWordlist, "\n")
	})
	return &words{effLargeWordlistVal.list, count}
}
