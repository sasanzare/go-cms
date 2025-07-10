package utils

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

const (
	ValidationFailedMsg = "Validation failed"
)

// JSONResponse represents the standard response structure for API endpoints.
//
// Fields:
//   - Success: indicates if the request was successful
//   - Message: optional success message
//   - Data: optional payload data
//   - Error: optional error message
type JSONResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// SendSuccess sends a successful JSON response with data and optional message.
//
// Parameters:
//   - c: Gin context
//   - message: optional success message
//   - data: payload data to include in response
//
// Example:
//   SendSuccess(c, "Operation successful", map[string]interface{}{"id": 123})
func SendSuccess(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, JSONResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// SendSuccessMessage sends a successful JSON response with only a message.
//
// Parameters:
//   - c: Gin context
//   - message: success message to include
//
// Example:
//   SendSuccessMessage(c, "Operation completed successfully")
func SendSuccessMessage(c *gin.Context, message string) {
	c.JSON(http.StatusOK, JSONResponse{
		Success: true,
		Message: message,
	})
}

// SendError sends an error JSON response with specified status code.
//
// Parameters:
//   - c: Gin context
//   - statusCode: HTTP status code for error
//   - errorMessage: description of the error
//
// Example:
//   SendError(c, http.StatusNotFound, "Resource not found")
func SendError(c *gin.Context, statusCode int, errorMessage string) {
	c.JSON(statusCode, JSONResponse{
		Success: false,
		Error:   errorMessage,
	})
}

// SendValidationError sends a validation error response with details.
//
// Parameters:
//   - c: Gin context
//   - errors: validation error details
//
// Example:
//   SendValidationError(c, map[string]string{"email": "Invalid format"})
func SendValidationError(c *gin.Context, errors interface{}) {
	c.JSON(http.StatusBadRequest, JSONResponse{
		Success: false,
		Error:   ValidationFailedMsg,
		Data:    errors,
	})
}