package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const JWTKEY = "Test123"

type UserClaims struct {
	Username string
	jwt.StandardClaims
}

// createJwt creates a JWT with the given expiration time and UserClaims
func createJwt(expTime time.Time, username string) (token *jwt.Token) {
	claims := &UserClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return
}

// validateJwtString ensures the given token is valid and returns the claims associated with it
func validateJwtString(tokenStr string) (*UserClaims, error) {
	var claims *UserClaims
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return JWTKEY, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims")
	}

	return claims, nil
}
