package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (*string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}

	hashedPasswordStr := string(hashedPassword)

	return &hashedPasswordStr, nil
}

func VerifyPassword(hashedPassword string, candidatePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
}
