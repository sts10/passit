package passit

import (
	"math/rand"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/assert"
)

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
		assert.Equal(t, size, countEmojiInString(pass),
			"countEmojiInString(%q)", pass)
		assert.Truef(t, utf8.ValidString(pass),
			"utf8.ValidString(%q)", pass)
	}
}

func TestEmojiValid(t *testing.T) {
	for _, emoji := range unicodeEmoji {
		assert.Truef(t, utf8.ValidString(emoji),
			"utf8.ValidString(%q)", emoji)
	}
}

func countEmojiInString(s string) int {
	var count int
outer:
	for len(s) > 0 {
		for i := len(unicodeEmoji) - 1; i >= 0; i-- {
			emoji := unicodeEmoji[i]
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
