package pwordgen

import (
	"crypto/rand"
	"math/big"
)

var StdChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-_=+,.?/:;{}[]`~")

func NewPassword(length int) string {
	return rand_char(length, StdChars)
}

func rand_char(length int, chars []byte) string {
	new_pword := make([]byte, length)
	for i := 0; i < length; i++ {
		j, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			panic(err)
		}
		if int(j.Int64()) == 0 {
			new_pword[i] = chars[0]
			continue
		}
		new_pword[i] = chars[len(chars)%int(j.Int64())]
	}
	return string(new_pword)
}
