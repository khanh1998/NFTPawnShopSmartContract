package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const minLengthSecretKey = 32

// JwtMaker is like a container which contains the secret key and methods to generate JWT based on the secret key
type JwtMaker struct {
	secretKet string
}

func NewJWTMaker(secretKey string) (*JwtMaker, error) {
	if len(secretKey) < minLengthSecretKey {
		return nil, errors.New(fmt.Sprintf("Minimum length of secret key is %v", minLengthSecretKey))
	}
	return &JwtMaker{
		secretKet: secretKey,
	}, nil
}

func (j *JwtMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(j.secretKet))
}

func (maker *JwtMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrorInvalidToken
		}
		return []byte(maker.secretKet), nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrorExpiredToken) {
			return nil, ErrorExpiredToken
		}
		return nil, ErrorInvalidToken
	}
	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, err
	}
	return payload, nil
}
