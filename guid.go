package randstr

import (
	"time"
)

import(
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

func RandomString(s int) string {
	// generate random string
	mathrand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 32)
	for i := range b {
		b[i] = letterRunes[mathrand.Intn(len(letterRunes))]
	}
	return string(b)
}
