package main

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func TestCreateJwt(t *testing.T) {
	expTime := time.Now().Add(5 * time.Minute)
	username := "test"

	// Execution
	token := createJwt(expTime, username)

	// Assertions
	// Check signing method
	if token.Method != jwt.SigningMethodHS256 {
		t.Errorf("Invalid token signing method: %s\n", token.Method)
	}

	// Check username in claims
	if claims, ok := token.Claims.(*UserClaims); !ok || ok && claims.Username != "test" {
		t.Errorf("Invalid claims: %v\n", token.Claims)
	}

	// Check expiration date in claims
	if claims, ok := token.Claims.(*UserClaims); !ok || ok && claims.ExpiresAt != expTime.Unix() {
		t.Errorf("Invalid expiration time: %d\n", claims.ExpiresAt)
	}
}
