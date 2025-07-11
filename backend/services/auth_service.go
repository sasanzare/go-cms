package services

import (
	"errors"
	"fmt"
	"github.com/sasanzare/go-cms/models"
	"github.com/sasanzare/go-cms/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// AuthService provides authentication and authorization operations
// including user registration, login, role checking and user management.
type AuthService struct {
	db *gorm.DB
}

// NewAuthService creates a new instance of AuthService with database dependency
//
// Parameters:
//   - db: GORM database instance
//
// Returns:
//   - *AuthService: initialized AuthService
func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{db: db}
}

// RegisterUser registers a new user with validation and role assignment
//
// Parameters:
//   - user: User model with registration data
//   - role: Desired role for the user (defaults to UserRoleUser)
//
// Returns:
//   - *models.User: Registered user data
//   - error: Registration error if any
//
// Validation Rules:
//   - Email must be valid format
//   - Password must meet complexity requirements
//   - Email must be unique
//
// Security:
//   - Password is hashed before storage
func (s *AuthService) RegisterUser(user *models.User, role string) (*models.User, error) {
	// Validate email format
	if !utils.ValidateEmail(user.Email) {
		return nil, errors.New("invalid email format")
	}

	// Validate password complexity
	if !utils.ValidatePassword(user.Password) {
		return nil, errors.New("password must contain at least 8 characters, uppercase, lowercase, number and special character")
	}

	// Check for existing email
	var existingUser models.User
	if err := s.db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return nil, errors.New("email already registered")
	}

	// Hash password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}
	user.Password = string(hashedPassword)

	// Assign role with validation
	switch role {
	case models.UserRoleAdmin:
		user.Role = models.UserRoleAdmin
	case models.UserRoleAuthor:
		user.Role = models.UserRoleAuthor
	case models.UserRoleEditor:
		user.Role = models.UserRoleEditor
	default:
		user.Role = models.UserRoleUser
	}

	// Create user record
	if err := s.db.Create(&user).Error; err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}

// Login authenticates user and generates JWT token
//
// Parameters:
//   - email: User's email address
//   - password: User's password
//
// Returns:
//   - string: JWT token for authenticated session
//   - error: Authentication error if any
//
// Flow:
//   1. Validates email format
//   2. Checks user existence
//   3. Verifies password
//   4. Generates JWT token
func (s *AuthService) Login(email, password string) (string, error) {
	// Validate email format
	if !utils.ValidateEmail(email) {
		return "", errors.New("invalid email format")
	}

	// Find user by email
	var user models.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		return "", errors.New("user not found")
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid password")
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return token, nil
}

// CheckUserRole verifies if user has the required role
//
// Parameters:
//   - userID: ID of the user to check
//   - requiredRole: Role to verify against
//
// Returns:
//   - bool: True if user has the required role
//
// Note:
//   - Typically used for middleware authorization
func (s *AuthService) CheckUserRole(userID uint, requiredRole string) bool {
	var user models.User
	if err := s.db.First(&user, userID).Error; err != nil {
		return false
	}
	return user.Role == requiredRole
}

// GetUserByID retrieves user by ID for access control purposes
//
// Parameters:
//   - userID: ID of the user to retrieve
//
// Returns:
//   - *models.User: User data if found
//   - error: Error if user not found
func (s *AuthService) GetUserByID(userID uint) (*models.User, error) {
	var user models.User
	if err := s.db.First(&user, userID).Error; err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}