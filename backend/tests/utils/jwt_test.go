package utils_test

import (
	"crypto/rand"
	"crypto/rsa"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sasanzare/go-cms/utils"
	"github.com/stretchr/testify/assert"
)

// TestGenerateAndValidateToken tests both token generation and validation functionality.
// It covers various scenarios including valid tokens, invalid tokens, expired tokens,
// and tokens with invalid signing methods.
//
// Test Cases:
//   1. Successful token generation
//   2. Valid token validation
//   3. Invalid token format
//   4. Expired token
//   5. Token with invalid signing method
//
// Dependencies:
//   - utils.GenerateToken
//   - utils.ValidateToken
//   - utils.SetJWTSecret
//   - utils.GetJWTSecret
func TestGenerateAndValidateToken(t *testing.T) {
	// Setup test environment
	originalSecret := utils.GetJWTSecret()
	defer utils.SetJWTSecret(originalSecret) // Restore original secret after test
	utils.SetJWTSecret("test-secret")

	userID := uint(123)
	role := "admin"

	// Test token generation
	token, err := utils.GenerateToken(userID, role)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Test valid token validation
	claims, err := utils.ValidateToken(token)
	assert.NoError(t, err)
	assert.Equal(t, userID, claims.UserID)
	assert.Equal(t, role, claims.Role)

	// Test invalid token format
	_, err = utils.ValidateToken("invalid-token")
	assert.Error(t, err)

	// Test expired token
	expiredToken := generateExpiredToken(userID, role)
	_, err = utils.ValidateToken(expiredToken)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "token is expired")

	// Test token with invalid signing method
	invalidMethodToken := generateTokenWithInvalidSigningMethod(userID, role)
	_, err = utils.ValidateToken(invalidMethodToken)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "unexpected signing method")
}

// generateExpiredToken generates an intentionally expired JWT token for testing purposes.
//
// Parameters:
//   - userID: The user ID to include in token claims (uint)
//   - role: The user role to include in token claims (string)
//
// Returns:
//   - string: The signed expired JWT token
//
// Note:
//   - Uses HS256 signing method
//   - Sets expiration to 1 hour in the past
func generateExpiredToken(userID uint, role string) string {
	claims := &utils.Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(-1 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte("test-secret"))
	return signedToken
}

// generateTokenWithInvalidSigningMethod generates a JWT token with RS256 signing method
// while the server expects HS256, used to test invalid signing method detection.
//
// Parameters:
//   - userID: The user ID to include in token claims (uint)
//   - role: The user role to include in token claims (string)
//
// Returns:
//   - string: The signed JWT token with invalid signing method
//
// Security Note:
//   - Generates a temporary RSA key for testing purposes only
func generateTokenWithInvalidSigningMethod(userID uint, role string) string {
	// Generate temporary RSA key for testing
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)

	claims := &utils.Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedToken, _ := token.SignedString(privateKey)
	return signedToken
}