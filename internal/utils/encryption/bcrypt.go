package encryption

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrPasswordTooLong           = errors.New("password too long")
	ErrMismatchedHashAndPassword = errors.New("mismatched hash and password")
)

func (e *EncryptionUtils) GenerateHashFromPassword(password string) (string, error) {
	hashString, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		if err == bcrypt.ErrPasswordTooLong {
			return "", ErrPasswordTooLong
		}
		return "", err
	}
	return string(hashString), nil
}

func (e *EncryptionUtils) CompareHashAndPassword(hash, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		if err == bcrypt.ErrPasswordTooLong {
			return ErrPasswordTooLong
		}
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return ErrMismatchedHashAndPassword
		}
		return err
	}
	return nil
}
