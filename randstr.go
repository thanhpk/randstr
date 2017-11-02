package randstr

import (
	"time"
	"crypto/rand"
	"encoding/hex"
	mathrand "math/rand"
)

func Byte(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}

func RandomBytes(n int) []byte {
	return Byte(n)
}

func Base64(s int) string {
	return String(s, "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+/")
}

func Base62(s int) string {
	return String(s)
}

func Hex(s int) string {
	b := RandomBytes(s)
	hexstring := hex.EncodeToString(b)
	return hexstring
}

func RandomHex(s int) string {
	return Hex(s)
}

func String(s int, letters ...string) string {
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

func RandomString(s int, letters ...string) string { // s number of character
	return String(s, letters...)
}
