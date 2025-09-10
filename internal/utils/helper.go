package utils

import (
	
	"github.com/gin-gonic/gin"
	"errors"
)

// ErrorResponse mengirimkan respons JSON yang terstandarisasi untuk error.
func ErrorResponse(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"success": false,
		"error":   message,
	})
}

// SuccessResponse mengirimkan respons JSON yang terstandarisasi untuk success.
func SuccessResponse(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, gin.H{
		"success": true,
		"message": message,
		"data":    data,
	})
}

// GetUserID mengambil ID pengguna dari konteks Gin setelah autentikasi.
// Ini sangat penting untuk mencatat siapa yang membuat transaksi.
func GetUserID(c *gin.Context) (uint, error) {
	userID, exists := c.Get("userID")
	if !exists {
		return 0, errors.New("user ID not found in context")
	}

	uid, ok := userID.(uint)
	if !ok {
		return 0, errors.New("invalid user ID type in context")
	}

	return uid, nil
}