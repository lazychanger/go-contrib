package stringer

import (
	"math"
	"math/rand"
	"strings"
	"time"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

var defaultAlphabet = []rune(Base62Chars)

func Random(length int, alphabets ...string) string {
	var (
		alphabet = getAlphabet(alphabets...)
		rs       = make([]rune, length)
	)

	for i := 0; i < length; i++ {
		rs[i] = alphabet[r.Intn(len(alphabet))]
	}

	return string(rs)
}

func getAlphabet(alphabets ...string) (alphabet []rune) {
	if len(alphabets) == 0 {
		alphabet = defaultAlphabet
	} else {
		alphabet = []rune(strings.Join(alphabets, ""))
	}
	return alphabet
}

func Shuffle(s string) string {
	rs := []rune(s)
	for i := range rs {
		j := r.Intn(i + 1)
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}

// ToShort is make string transfer to short string.
// example as aaaaa-bbbb to ab
func ToShort(s string) string {
	var (
		slen   = len(s)
		result = [math.MaxUint8]byte{}
		ri     = uint8(0)
	)
	for i := 0; i < slen; i++ {
		if i == 0 || s[i-1] == '-' || s[i-1] == '_' || s[i-1] == ' ' {
			ri += 1
			result[ri-1] = s[i]
		}
	}
	return string(result[:ri])
}

// UpperToShort is make string transfer to short string.
// example as AbbbsBbbCb to ABC
func UpperToShort(s string) string {
	var (
		slen             = len(s)
		result           = [math.MaxUint8]byte{}
		ri               = uint8(0)
		minRune, maxRune = 'A', 'Z'
	)

	for i := 0; i < slen; i++ {
		if i == 0 || s[i-1] == '-' || s[i-1] == '_' || s[i-1] == ' ' || (int32(s[i]) <= maxRune && int32(s[i]) >= minRune) {
			ri += 1

			result[ri-1] = s[i]
		}
	}
	return strings.ToUpper(string(result[:ri]))
}
