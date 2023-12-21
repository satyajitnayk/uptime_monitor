package auth

import (
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

// TestGenerateJWT tests the GenerateJWT function.
func TestGenerateJWT(t *testing.T) {
	// Save the original secret key and restore it at the end of the test
	originalSecretKey := secretKey
	defer func() {
		secretKey = originalSecretKey
	}()

	// Set a temporary secret key for testing
	testSecretKey := []byte("test_secret_key")
	secretKey = testSecretKey

	// Test data
	email := "test@example.com"

	// Generate JWT
	tokenString, err := GenerateJWT(email)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	// Parse the token to verify its contents
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return testSecretKey, nil
	})
	assert.NoError(t, err)
	assert.True(t, token.Valid)

	// Verify the claims
	claims, ok := token.Claims.(*CustomClaims)
	assert.True(t, ok)
	assert.Equal(t, email, claims.Email)
}

// TestValidateJWT tests the ValidateJWT function.
func TestValidateJWT(t *testing.T) {
	// Save the original secret key and restore it at the end of the test
	originalSecretKey := secretKey
	defer func() {
		secretKey = originalSecretKey
	}()

	// Set a temporary secret key for testing
	testSecretKey := []byte("test_secret_key")
	secretKey = testSecretKey

	// Test data
	email := "test@example.com"

	// Generate JWT
	tokenString, err := GenerateJWT(email)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	// Validate JWT
	claims, err := ValidateJWT(tokenString)
	assert.NoError(t, err)
	assert.NotNil(t, claims)
	assert.Equal(t, email, claims.Email)
}

// TestInvalidToken tests the ValidateJWT function with an invalid token.
func TestInvalidToken(t *testing.T) {
	// Test data: an invalid token string
	invalidTokenString := "invalid_token_string"

	// Validate JWT with an invalid token
	claims, err := ValidateJWT(invalidTokenString)
	assert.Error(t, err)
	assert.Nil(t, claims)
	assert.Equal(t, "token contains an invalid number of segments", err.Error())
}
