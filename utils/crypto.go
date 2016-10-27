package utils

import "golang.org/x/crypto/bcrypt"

const cost = 15

// GeneratePassword generates a hashed password from a string
func GeneratePassword(password string) (string, error) {
	if hash, err := bcrypt.GenerateFromPassword([]byte(password), cost); err == nil {
		return string(hash), nil
	} else {
		return "", err
	}
}

// CheckPassword compares a plain password with the equivelant hashed,
// returns true if they match, false otherwise
func CheckPassword(plain, hashed []byte) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain)); err == nil {
		return true
	} else {
		return false
	}
}
