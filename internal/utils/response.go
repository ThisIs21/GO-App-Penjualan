package utils

import "github.com/gin-gonic/gin"

func OK(c *gin.Context, data any) {
	c.JSON(200, gin.H{"success": true, "data": data})
}
func Created(c *gin.Context, data any) {
	c.JSON(201, gin.H{"success": true, "data": data})
}
func BadRequest(c *gin.Context, msg string) {
	c.JSON(400, gin.H{"success": false, "error": msg})
}
func Unauthorized(c *gin.Context, msg string) {
	c.JSON(401, gin.H{"success": false, "error": msg})
}
func Forbidden(c *gin.Context, msg string) {
	c.JSON(403, gin.H{"success": false, "error": msg})
}
func NotFound(c *gin.Context, msg string) {
	c.JSON(404, gin.H{"success": false, "error": msg})
}
func ServerError(c *gin.Context, err error) {
	c.JSON(500, gin.H{"success": false, "error": err.Error()})
}
