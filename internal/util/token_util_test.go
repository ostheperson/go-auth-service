package util

import (
	"testing"

	"github.com/ostheperson/go-auth-service/internal/domain"
)

func TestVerifyAndExtract(t *testing.T) {
	secret := "secret"
	user := domain.Users{Username: "onion", ID: 1}
	tokenString, err := CreateAccessToken(&user, secret, 1)
	if err != nil {
		t.Errorf("error creating token: %v", err.Error())
		return
	}

	claims, err := VerifyAndExtract(tokenString, secret)
	if err != nil {
		t.Errorf("VerifyAndExtract returned an error: %v", err.Error())
		return
	}

	// Check if the claims were extracted successfully
	if claims == nil {
		t.Error("VerifyAndExtract returned nil claims")
		return
	}
}
