package jwt

import (
	"crypto/rsa"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	privateKey  *rsa.PrivateKey
	tokenString string
)

func LoadPrivateKey(path string) error {
	fileByte, err := os.ReadFile(path)

	if err != nil {
		return err
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(fileByte)

	if err != nil {
		return err
	}
	return nil
}

func GenerateToken(iss string, userId uint, exp time.Time, subjet string) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iss":    iss,
		"sub":    subjet,
		"exp":    exp,
		"userId": userId,
	}).SignedString(privateKey)
}
