package authn

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Auth interface {
	IsTokenValid() (bool, error)
	CreateToken(email string, userid int) (string, error)
}

type SimpleAuth struct {
	salt      []byte
	tokenLife time.Duration
}

func NewSimpleAuth(salt []byte, tokenLife time.Duration) SimpleAuth {
	return SimpleAuth{
		salt:      salt,
		tokenLife: tokenLife,
	}
}

func (s SimpleAuth) IsTokenValid() (bool, error) {
	return true, nil
}

func (s SimpleAuth) CreateToken(email string, userid int) (string, error) {
	claims := Claims{
		Email: email,
		ID:    userid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(s.tokenLife).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.salt)
}
