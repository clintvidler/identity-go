package rpc

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

func (s *IdentityService) ParseClaims(token string) (jwt.MapClaims, error) {
	pk, err := jwt.ParseRSAPublicKeyFromPEM(s.pubKey)

	if err != nil {
		panic(err)
	}

	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}

		return pk, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := tok.Claims.(jwt.MapClaims)

	if !ok || !tok.Valid {
		return nil, err
	}

	return claims, nil
}
