package utils

import (
	"math/rand"
	"regexp"
	"time"
)

//RandString :: generates random string with lenth of n.
func RandString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// IsURL - checks if string is URL.
func IsURL(url string) bool {
	r, _ := regexp.Compile("https?://.+\\..+")
	return r.MatchString(url)
}
