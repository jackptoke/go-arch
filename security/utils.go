package security

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %v", err)
	}
	return bytes, nil
}

func ComparePassword(password string, hashedPassword []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		return fmt.Errorf("invalid password: %v", err)
	}
	return nil
}

func SignMessage(msg []byte, key []byte) ([]byte, error) {
	h := hmac.New(sha512.New, key)

	if _, err := h.Write(msg); err != nil {
		return nil, fmt.Errorf("error hashing message: %v", err)
	}
	// signature, error
	return h.Sum(nil), nil
}

func CheckSignature(signature []byte, msg []byte, key []byte) (bool, error) {
	newSig, err := SignMessage(msg, key)
	if err != nil {
		return false, fmt.Errorf("error signing message: %v", err)
	}

	same := hmac.Equal(newSig, signature)
	return same, nil
}
