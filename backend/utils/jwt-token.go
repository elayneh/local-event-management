package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func createJWTToken(payload map[string]interface{}, secretKey string, tokenExpiration int64) (string, error) {
	claims := jwt.MapClaims(payload)
	claims["exp"] = tokenExpiration
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func GetToken(userId string, role string, isEmailVerified bool) (string, error) {
	isEmailVerifiedStr := "false"
	if isEmailVerified {
		isEmailVerifiedStr = "true"
	}
	secretKey := os.Getenv("JWT_SECRET")
	payload := map[string]interface{}{
		"sub": userId,
		"iat": time.Now().Unix(),
		"https://hasura.io/jwt/claims": map[string]interface{}{
			"x-hasura-allowed-roles":     []string{"user", "admin"},
			"x-hasura-default-role":      "user",
			"x-hasura-user-id":           userId,
			"x-hasura-role":              role,
			"x-hasura-is-email-verified": isEmailVerifiedStr,
		},
		"metadata": map[string]interface{}{
			"userId": userId,
		},
	}
	tokenExpiration := time.Now().Add(time.Hour * 24).Unix()
	token, err := createJWTToken(payload, secretKey, tokenExpiration)
	if err != nil {
		return "", err
	}
	return token, nil
}
