package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/sasanzare/go-cms/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)


// TestResponses tests all response utility functions in the utils package.
// It verifies the correctness of HTTP status codes and JSON response formats
// for various response scenarios.
//
// Test Cases:
//   1. SendSuccess with data payload
//   2. SendSuccess with message only
//   3. SendError with custom status code
//   4. SendValidationError with error details
//
// Dependencies:
//   - utils.SendSuccess
//   - utils.SendSuccessMessage
//   - utils.SendError
//   - utils.SendValidationError
func TestResponses(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Test SendSuccess with data payload
	t.Run("SendSuccess with data", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		utils.SendSuccess(c, "test message", gin.H{"key": "value"})

		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, `{"success":true,"message":"test message","data":{"key":"value"}}`, w.Body.String())
	})

	// Test SendSuccess with message only
	t.Run("SendSuccess without data", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		utils.SendSuccessMessage(c, "test message")

		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, `{"success":true,"message":"test message"}`, w.Body.String())
	})

	// Test SendError with custom status code
	t.Run("SendError", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		utils.SendError(c, http.StatusBadRequest, "test error")

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, `{"success":false,"error":"test error"}`, w.Body.String())
	})

	// Test SendValidationError with error details
	t.Run("SendValidationError", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		utils.SendValidationError(c, gin.H{"field": "error"})

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, `{"success":false,"error":"Validation failed","data":{"field":"error"}}`, w.Body.String())
	})
}