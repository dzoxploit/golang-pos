package utils

import (
	"errors"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// JWTSecretKey is the secret key used to sign JWT tokens. This should be kept secure and should not be hard-coded in production.
var JWTSecretKey = []byte("your-secret-key")

var ErrInvalidCredentials = errors.New("invalid credentials")

// GenerateToken generates a new JWT token with the given user ID as the subject.
func GenerateToken(userID uint) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &jwt.StandardClaims{
		Subject:   strconv.Itoa(int(userID)),
		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(JWTSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken validates the provided JWT token and returns the token object if it's valid.
func ValidateToken(tokenString string) (*jwt.Token, error) {
	// Parse the token with the provided secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return JWTSecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}

// HashPassword hashes the provided password using a secure hashing algorithm (bcrypt).
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPasswordHash compares the provided password with the hashed password to check if they match.
func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
