package helper

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// to encrypt password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

// to compare hashed password with password
func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

// jwt key
var secretKey = []byte(os.Getenv("JWT_SECRET"))

// to create jwt token
func CreateToken(email string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"iss":   "go-auth",
		"exp":   time.Now().Add(time.Hour).Unix(),
		"iat":   time.Now().Unix(),
	})

	tokenStr, err := claims.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

// to verify jwt token
func VerifyToken(token string) error {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !t.Valid {
		return fmt.Errorf("invalid token")
	}

	log.Println(t)
	return err
}
