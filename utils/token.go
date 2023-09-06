package utils

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

func GenerateToken(ttl time.Duration, payload interface{}, secretJwtKey string) (*string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	now := time.Now().UTC()
	claims := token.Claims.(jwt.MapClaims)

	claims["sub"] = payload
	claims["exp"] = now.Add(ttl).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	tokenString, err := token.SignedString([]byte(secretJwtKey))
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}

func ValidateToken(token string, signedJwtKey string) (interface{}, error) {
	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected method")
		}

		return []byte(signedJwtKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return nil, err
	}

	return claims["sub"], nil

}
