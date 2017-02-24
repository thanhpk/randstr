package ranstr

import(
	"crypto/rand"
	"encoding/hex"
)

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	if err != nil {
		return "", err
	}
	
	hexstring :=  hex.EncodeToString(b)

	return hexstring, err
}

