package grand

import "math/rand"

// Credits to icza of https://stackoverflow.com/a/31832326/1161743 for the original code.

// Character sets for random string generation
const (
	CharSetEnglishAlphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CharSetBase62          = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var defaultLetterBytes = CharSetEnglishAlphabet

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// Generator is the instance that generates random strings from the given charset
type Generator struct {
	charset string
}

// NewGenerator returns a new Generator instance
func NewGenerator(charset string) *Generator {
	return &Generator{
		charset: charset,
	}
}

// GenerateRandomString generates a length-n random string from the given character set
func (g *Generator) GenerateRandomString(n int) string {
	b := make([]byte, n)

	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}

		if idx := int(cache & letterIdxMask); idx < len(g.charset) {
			b[i] = g.charset[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

// GenerateRandomString generates a length-n random string from the default character set (English alphabets).
// Be sure to seed the random function using rand.Seed to initialize the generator first.
func GenerateRandomString(n int) string {
	return NewGenerator(CharSetEnglishAlphabet).GenerateRandomString(n)
}
