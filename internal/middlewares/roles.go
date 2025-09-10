package middlewares

import (
	"github.com/gin-gonic/gin"
	"app-penjualan/internal/utils"
)

// MustUserID mengambil user ID dari konteks Gin setelah token diuraikan.
// Fungsi ini menghentikan permintaan jika user ID tidak ditemukan.
func MustUserID(c *gin.Context) uint {
	uid, exists := c.Get("uid")
	if !exists {
		// Middleware sebelumnya (JWTAuth) seharusnya sudah menangani ini
		// Tetapi ini adalah langkah keamanan tambahan
		utils.Unauthorized(c, "invalid token")
		c.Abort()
		return 0
	}
	// Pastikan tipe data yang dikembalikan adalah uint
	return uid.(uint)
}

// RequireRoles adalah middleware untuk memeriksa apakah role pengguna diizinkan.
func RequireRoles(roles ...string) gin.HandlerFunc {
	allowed := map[string]bool{}
	for _, r := range roles { allowed[r] = true }

	return func(c *gin.Context) {
		role, _ := c.Get("role")
		if roleStr, ok := role.(string); ok && allowed[roleStr] {
			c.Next(); return
		}
		utils.Forbidden(c, "role not allowed")
		c.Abort()
	}
}