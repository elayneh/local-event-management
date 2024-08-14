package utils

import "github.com/dgrijalva/jwt-go"

var (
	jwtKey          = []byte("my_secret_key")
	graphqlEndpoint = "http://localhost:8080/v1/graphql"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
