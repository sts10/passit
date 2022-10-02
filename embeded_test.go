package passit

import (
	"math/rand"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEFFLargeWordlist(t *testing.T) {
	const size = 8

	testRand := rand.New(rand.NewSource(0))

	pass, err := EFFLargeWordlist(size).Password(testRand)
	require.NoError(t, err)

	assert.Equal(t, "native remover dismay vocation sepia backtalk think conjure", pass)
	assert.Equal(t, size-1, strings.Count(pass, " "),
		`strings.Count(%q, " ")`, pass)
	assert.Truef(t, utf8.ValidString(pass),
		"utf8.ValidString(%q)", pass)

	_, err = FromWords(effLargeWordlistVal.list...)
	assert.NoError(t, err, "wordlist valid")
}

func TestEmoji11(t *testing.T) {
	const size = 25

	for i, expect := range []string{
		"🇳🇱🧚🏼\u200d♂️🚵🏼\u200d♂️🇱🇮🙅🏼👨\u200d⚕️🧚🏿\u200d♀️🙇🏻👦🏾🇧🇻🚴🏿🏊\u200d♀️🏌🏿\u200d♂️💂🏽\u200d♂️👨🏽\u200d🚀🎅🏽🇮🇸🙎🏻\u200d♀️🤴🏻🤸🏼\u200d♀️🤦🏿\u200d♂️🧛🏿👷🏾\u200d♀️🧜🏻\u200d♂️🛀🏿",
		"🤲🏼✍🏼🚴\u200d♂️🧛🏾\u200d♀️💂🏽\u200d♂️🙇🏿\u200d♂️🧜🏽👨🏿\u200d🔬🇳🇫👨🏿\u200d🔧👩🏾\u200d🎤🏌🏿\u200d♀️👨🏽\u200d🎨👩\u200d👦🧘🏻🧗🏽\u200d♀️🙎🏿👨\u200d👨\u200d👦👨🏼\u200d🎤💂\u200d♂️👌🏼🛀🏾👇🏾🧖🏼💆🏼\u200d♀️",
		"🤹🏿\u200d♂️🕵🏼\u200d♂️👨🏻\u200d🌾🙆🏾\u200d♀️🇲🇿🤾🏾\u200d♂️💆🏽\u200d♀️🇽🇰👩🏾\u200d🎨🏃🏻\u200d♀️🇵🇳🇬🇳🦹🏿\u200d♀️🕵🏿🇪🇭🏃🏾\u200d♂️👸🏾🧙🏼\u200d♀️🚴🏼\u200d♂️🧚🏾👁️\u200d🗨️🧛🏿🤾🏿\u200d♂️5️⃣👦🏼",
		"💂🏻🇲🇭🏇🏾👊🏿🚶🏻\u200d♀️💂🏾\u200d♀️🚵🏿\u200d♂️🙋🏼\u200d♂️👳🏻👩🏼\u200d🎤👱🏾\u200d♂️👨🏽\u200d🏭👩🏻\u200d🍳⛹🏼\u200d♀️🧑🏽👮🏻\u200d♀️🙍🏽🇸🇦🙆🏻\u200d♂️👩🏾\u200d🦳💇🏽\u200d♂️🇱🇧👩🏼\u200d🏭👱🏻🚴🏼\u200d♂️",
	} {
		testRand := rand.New(rand.NewSource(int64(i)))

		pass, err := Emoji11(size).Password(testRand)
		if !assert.NoError(t, err) {
			continue
		}

		assert.Equal(t, expect, pass)
		assert.Equal(t, size, countEmojiInString(emoji11ListVal.list, pass),
			"countEmojiInString(%q)", pass)
		assert.Truef(t, utf8.ValidString(pass),
			"utf8.ValidString(%q)", pass)
	}
}

func TestEmoji13(t *testing.T) {
	const size = 25

	for i, expect := range []string{
		"👷\u200d♀️🕸️🏃\u200d♀️🐎🚵🏽\u200d♂️👩🏿\u200d🏫🙋🏻💖👐🏽🧷🔇🌛🙍🏿\u200d♀️👨🏿\u200d🎓🕵🏼👴🧗🏽\u200d♀️💺🇹🇯🔎👳🏽\u200d♂️🤞🏼👩🏻\u200d🦰📦🎂",
		"🏮🤌🏽🧑🏽\u200d🤝\u200d🧑🏾👩🏼\u200d🍳🤽🏽🤳🏃🏾\u200d♂️❕🐣🆚🔧👏🏽🏄🏽💇🏼🥾🤟🏼👨\u200d🚀🦶🏻🧚🏻🛌🏻🚨💒😍🇵🇼🙎🏻",
		"😪🗨️📍☃️🏄🏼\u200d♂️🌑👩🏼\u200d🚒👷🏽\u200d♂️🧙🏾\u200d♀️👌🤹🏿\u200d♀️🈳🧑🏿\u200d🍳🏝️🇷🇸🧑🏼\u200d🎓🧑🏽\u200d⚕️🦻🏽👩\u200d🍼🧑🏿\u200d🏫🇸🇸👲⏺️☺️🦹\u200d♂️",
		"🏴\U000e0067\U000e0062\U000e0065\U000e006e\U000e0067\U000e007f🙆🏽🫑🧘🏽🚄🇸🇧🚶🏾\u200d♀️🤚🏻🦹\u200d♀️👩🏼\u200d🦼👎😵🤷🇻🇬👩🏿\u200d🚀🏊🏻\u200d♀️🙋🏻\u200d♂️👨🏽\u200d🦲🙂👩🏾\u200d🤝\u200d👩🏼🧍🏾\u200d♂️🧖🏾\u200d♀️👩🦸🏽🧜🏽\u200d♀️",
	} {
		testRand := rand.New(rand.NewSource(int64(i)))

		pass, err := Emoji13(size).Password(testRand)
		if !assert.NoError(t, err) {
			continue
		}

		assert.Equal(t, expect, pass)
		assert.Equal(t, size, countEmojiInString(emoji13ListVal.list, pass),
			"countEmojiInString(%q)", pass)
		assert.Truef(t, utf8.ValidString(pass),
			"utf8.ValidString(%q)", pass)
	}
}

func TestEmojiValid(t *testing.T) {
	for _, emoji := range strings.Split(emoji11List, "\n") {
		assert.Truef(t, utf8.ValidString(emoji),
			"utf8.ValidString(%q)", emoji)
	}

	for _, emoji := range strings.Split(emoji13List, "\n") {
		assert.Truef(t, utf8.ValidString(emoji),
			"utf8.ValidString(%q)", emoji)
	}
}

func countEmojiInString(list []string, s string) int {
	var count int
outer:
	for len(s) > 0 {
		for i := len(list) - 1; i >= 0; i-- {
			emoji := list[i]
			if strings.HasPrefix(s, emoji) {
				count++
				s = s[len(emoji):]
				continue outer
			}
		}

		return -1
	}

	return count
}
