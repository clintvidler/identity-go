package rpc

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func (s *IdentityService) Generate(uid, aid string) (string, string, error) {
	// Read the private RSA key
	key, err := jwt.ParseRSAPrivateKeyFromPEM(s.prvKey)
	if err != nil {
		return "", "", err
	}

	// Generate the access token
	attl, err := time.ParseDuration(os.Getenv("ACCESS_TOKEN_DURATION"))
	if err != nil {
		panic(err)
	}

	ac := jwt.StandardClaims{}
	ac.Audience = aid
	ac.Subject = uid
	ac.ExpiresAt = time.Now().Add(attl).Unix()

	at, err := jwt.NewWithClaims(jwt.SigningMethodRS256, ac).SignedString(key)
	if err != nil {
		return "", "", err
	}

	// Generate the refresh token
	rttl, err := time.ParseDuration(os.Getenv("REFRESH_TOKEN_DURATION"))
	if err != nil {
		panic(err)
	}

	rc := jwt.StandardClaims{}
	rc.Subject = uid
	rc.Audience = aid
	rc.ExpiresAt = time.Now().Add(rttl).Unix()

	rt, err := jwt.NewWithClaims(jwt.SigningMethodRS256, rc).SignedString(key)
	if err != nil {
		return "", "", err
	}

	// Save the refresh token to the database

	err = s.data.Token.CreateOne(uid, aid, rt, rttl)
	if err != nil {
		return "", "", err
	}

	return at, rt, err
}
