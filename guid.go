package randstr

import(
	"crypto/rand"
	"encoding/hex"
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

