package main
import (
	"time"
	"crypto/rand"
	"encoding/hex"
	mathrand "math/rand"
)

func RandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func RandomHex(s int) (string, error) {
	b, err := RandomBytes(s)
	if err != nil {
		return "", err
	}
	hexstring :=  hex.EncodeToString(b)
	return hexstring, err
}

func RandomString(s int) string { // s number of character
	randomFactor, _ := RandomBytes(1)
	mathrand.Seed(time.Now().UnixNano() * int64(randomFactor[0]))
	var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, s)
	for i := range b {
		b[i] = letterRunes[mathrand.Intn(len(letterRunes))]
	}
	return string(b)
}
