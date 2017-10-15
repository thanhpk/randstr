package randstr

import (
	"time"
	"crypto/rand"
	"encoding/hex"
	mathrand "math/rand"
)

func RandomBytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}

func RandomHex(s int) string {
	b := RandomBytes(s)
	hexstring :=  hex.EncodeToString(b)
	return hexstring
}

func RandomString(s int, letters ...string) string { // s number of character
	randomFactor := RandomBytes(1)
	mathrand.Seed(time.Now().UnixNano() * int64(randomFactor[0]))

	var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	if len(letters) > 0 {
		letterRunes = []rune(letters[0])
	}
	b := make([]rune, s)
	for i := range b {
		b[i] = letterRunes[mathrand.Intn(len(letterRunes))]
	}
	return string(b)
}

func init() {}
