package mmp

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword encrypts a plaintext string and returns the hashed version in base64
func HashPassword(password string, format string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	switch format {
	case "base64":
		return base64.StdEncoding.EncodeToString(hashed), nil
	case "hex":
		return hex.EncodeToString(hashed), nil
	}

	return "", fmt.Errorf("unknown format: " + format)
}
