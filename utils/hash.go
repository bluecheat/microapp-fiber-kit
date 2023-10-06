package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash(text string) (string, error) {
	hexPassword, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	return string(hexPassword), err
}

func VerifyHash(hash, plainText string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plainText))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			// MEMO: err를 wrap 하여 상세를 전달하면 좋다
			return err
		}
		return err
	}
	return nil
}
