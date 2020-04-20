package authn

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Email string `json:"email"`
	ID    int    `json:"id"`
	jwt.StandardClaims
}
