package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	UserID uint   `json:"user_id"` // Unique identifier for the user
	Role   string `json:"role"`    // User role (e.g., "admin", "user")
	jwt.RegisteredClaims          // Embedded standard JWT claims
}

// GenerateToken creates a new JWT token for a user with the given ID and role.
//
// Parameters:
//   - userID: The unique identifier of the user (uint).
//   - role:   The role assigned to the user (string).
//
// Returns:
//   - string: The signed JWT token.
//   - error:  An error if token generation fails (e.g., secret key not set).
//
// Example:
//   token, err := GenerateToken(123, "admin")
func GenerateToken(userID uint, role string) (string, error) {
	claims := &Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Token expires in 24 hours
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateToken checks if a JWT token is valid and returns its claims.
//
// Parameters:
//   - tokenString: The JWT token to validate (string).
//
// Returns:
//   - *Claims: The decoded claims if the token is valid.
//   - error:   An error if validation fails (e.g., expired, invalid signature).
//
// Example:
//   claims, err := ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...")
func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// GetJWTSecret returns the current JWT secret key as a string.
// This is primarily for debugging or logging purposes (avoid exposing secrets in production).
//
// Returns:
//   - string: The JWT secret key.
func GetJWTSecret() string {
	return string(jwtSecret)
}

// SetJWTSecret updates the JWT secret key used for signing and validation.
// Note: This should be used cautiously (e.g., during tests or key rotation).
//
// Parameters:
//   - secret: The new secret key (string).
func SetJWTSecret(secret string) {
	jwtSecret = []byte(secret)
}