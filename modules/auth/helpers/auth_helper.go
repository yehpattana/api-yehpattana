package helpers

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func CheckIsValidEmailPattern(email string) bool {
	match, err := regexp.MatchString(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`, email)
	if err != nil {
		return false
	}
	return match
}

func BcryptHashingPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error hashing password: %v", err)
	}

	return string(hashedPassword), nil
}

func BcryptComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GeneratePassword(length int) string {
	// minimum length should be 6 characters - max 16 characters
	// TODO - should handle password must include at least one number, one special character, one uppercase letter, and one lowercase letter
	if length < 6 {
		length = 6
	}
	if length > 16 {
		length = 16
	}
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!#$%&*()_`"

	password := make([]byte, length)
	for i := 0; i < length; i++ {
		randIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		password[i] = charset[randIndex.Int64()]
	}

	return string(password)
}

func IsHashedPassword(password string) bool {
	// Regex to match the bcrypt hash pattern
	re := regexp.MustCompile(`^\$2[aby]\$\d{2}\$[./A-Za-z0-9]{53}$`)
	return re.MatchString(password)
}
