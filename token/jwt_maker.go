package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTMaker struct {
	secretKey string
}

const minSecretKeySize = 32

// returning interface so that jwt need to implement function of token Maker
func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d character", minSecretKeySize)
	}
	return &JWTMaker{secretKey}, nil
}

func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	// create a new jwtToken with jwt.NewWithClaims() function with HS256 algorithm of the jwt-go package
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	// return string of jwttoken by convert it into []byte
	return jwtToken.SignedString([]byte(maker.secretKey))
}

func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	// verify the signing algorithm from token header
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		// get signing method with token.Method, then convert it (interface) into signingMethodHMAC struct which is the instance source of HS256 Algorithm we used
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	}

	// before do anything with the token, we need to make sure the signing algorithm is same with the one we used
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)

	if err != nil {
		ver, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(ver.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}

		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
