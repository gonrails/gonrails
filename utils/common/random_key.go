package common

import "math/rand"

const (
	randLetters       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ012456789!@#$%^&*()_+"
	randLetterIdxBits = 7
	randletterIdxMask = 1<<randLetterIdxBits - 1
	randletterIdxMax  = 63 / randLetterIdxBits
)

func RandStr(n int) string {
	str := make([]byte, n)
	for i, cache, remain := n-1, rand.Int63(), randletterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), randletterIdxMax
		}

		if idx := int(cache & randletterIdxMask); idx < len(randLetters) {
			str[i] = randLetters[idx]
			i--
		}
		cache >>= randLetterIdxBits
		remain--
	}
	return string(str)
}
