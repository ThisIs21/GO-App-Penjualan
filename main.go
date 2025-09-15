
package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"app-penjualan/config"
	"app-penjualan/internal/routes" // Import package routes
	"app-penjualan/internal/utils"

	// Import library CORS untuk Gin
	"github.com/gin-contrib/cors"
	
)

// main adalah fungsi utama yang menjalankan aplikasi.
func main() {
	// Panggil InitApp untuk menginisialisasi semua yang dibutuhkan
	InitApp()
}

// InitApp menginisialisasi semua komponen aplikasi.
func InitApp() {
	// Memuat konfigurasi dari file .env
	cfg := config.LoadConfig()

	// Koneksi ke database
	config.ConnectDB(cfg)

	// Migrasi database
	config.MigrateDB(config.DB)

	// Menanamkan data default (roles)
	utils.SeedRoles(config.DB)

	// Menanamkan data user default setelah roles dibuat
	utils.SeedUsers(config.DB)

	// Buat router Gin
	r := gin.Default()

	// Konfigurasi dan gunakan middleware CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:12345", "http://192.168.100.134:12345"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	r.Use(cors.New(corsConfig))

	// Daftarkan semua rute aplikasi menggunakan package routes
	routes.Register(r, config.DB, cfg)

	// Jalankan server
	log.Println("Server berjalan di port :8081")
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
