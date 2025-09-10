package config

import (
	"fmt"
	"log"
	"os"

	"app-penjualan/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DBConfig holds the database connection configuration.
type DBConfig struct {
	User string
	Pass string
	Host string
	Port string
	Name string
}

// AppConfig holds the main application configuration.
type AppConfig struct {
	DB        DBConfig
	JWTSecret string
	JWTTTLHrs int
}

// Global variable for the database connection
var DB *gorm.DB

// LoadConfig reads configuration from a .env file and returns AppConfig.
func LoadConfig() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Gagal membaca file .env: %v", err)
	}

	return &AppConfig{
		DB: DBConfig{
			User: os.Getenv("DB_USER"),
			Pass: os.Getenv("DB_PASS"),
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
			Name: os.Getenv("DB_NAME"),
		},
		JWTSecret: os.Getenv("JWT_SECRET"),
		JWTTTLHrs: 24, // You can change this to read from .env if needed
	}
}

// ConnectDB uses the AppConfig to initialize and return a GORM database connection.
func ConnectDB(cfg *AppConfig) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DB.User, cfg.DB.Pass, cfg.DB.Host, cfg.DB.Port, cfg.DB.Name)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		log.Fatalf("Gagal koneksi database: %v", err)
	}

	DB = database
	fmt.Println("✅ Koneksi database berhasil")
}

// MigrateDB will perform a database migration based on the defined models.
func MigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.Role{}, // Tambahkan Role, jika belum ada
		&models.User{},
		&models.Category{},
		&models.Unit{},
		&models.Supplier{},
		&models.Customer{},
		&models.Voucher{},
		&models.Product{},
		
		&models.Sale{},
		&models.SaleItem{},
		&models.Purchase{},
		&models.PurchaseItem{},
		&models.StockOpname{},
		&models.SaleReturn{},
		&models.SaleReturnItem{},
		&models.PurchaseReturn{},
		&models.PurchaseReturnItem{},
	)
	if err != nil {
		log.Fatalf("Gagal melakukan migrasi database: %v", err)
	}
	fmt.Println("✅ Migrasi database berhasil")
}