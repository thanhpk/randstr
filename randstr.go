package randstr

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
)

func Bytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}

func Base64(s int) string {
	return String(s, "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+/")
}

func Base62(s int) string {
	return String(s)
}

func Hex(s int) string {
	b := Bytes(s)
	hexstring := hex.EncodeToString(b)
	return hexstring
}

var defRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func String(s int, letters ...string) string {
	var letterRunes []rune
	if len(letters) == 0 {
		letterRunes = defRunes
	} else {
		letterRunes = []rune(letters[0])
	}

	var bb bytes.Buffer
	bb.Grow(s)
	l := uint32(len(letterRunes))
	for i := 0; i < s; i++ {
		bb.WriteRune(letterRunes[binary.BigEndian.Uint32(Bytes(4))%l])
	}
	return bb.String()
}
