package utils

import (
	"fmt"
	"log"

	"app-penjualan/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// SeedRoles memasukkan data role default ke dalam database.
func SeedRoles(db *gorm.DB) {
	roles := []models.Role{
		{Name: models.RoleOwner},
		{Name: models.RoleKasir},
		{Name: models.RolePembelian},
		{Name: models.RoleGudang},
		{Name: models.RoleKepalaGudang},
	}

	for _, role := range roles {
		var existingRole models.Role
		if err := db.Where("name = ?", role.Name).First(&existingRole).Error; err != nil {
			if err != gorm.ErrRecordNotFound {
				log.Printf("Gagal memeriksa role %s: %v\n", role.Name, err)
			} else {
				if createErr := db.Create(&role).Error; createErr != nil {
					log.Fatalf("Gagal menanamkan role %s: %v\n", role.Name, createErr)
				}
				fmt.Printf("✅ Role '%s' berhasil ditanamkan.\n", role.Name)
			}
		} else {
			fmt.Printf("🟡 Role '%s' sudah ada, dilewati.\n", role.Name)
		}
	}
}

// SeedUsers memasukkan data user default untuk setiap role.
func SeedUsers(db *gorm.DB) {
	// Kata sandi default untuk semua pengguna
	password := "password123" 
	
	// Hashing kata sandi menggunakan bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Gagal menghash kata sandi: %v", err)
	}

	// Data pengguna yang akan dibuat
	users := []struct {
		Name     string
		Email    string
		RoleName models.RoleName
	}{
		{Name: "Owner", Email: "owner@mail.com", RoleName: models.RoleOwner},
		{Name: "Kasir1", Email: "kasir@mail.com", RoleName: models.RoleKasir},
		{Name: "Pembelian1", Email: "pembelian@mail.com", RoleName: models.RolePembelian},
		{Name: "Gudang1", Email: "gudang@mail.com", RoleName: models.RoleGudang},
		{Name: "KGudang1", Email: "kepalagudang@mail.com", RoleName: models.RoleKepalaGudang},
	}

	// Loop melalui data pengguna untuk membuat setiap user
	for _, userData := range users {
		var existingUser models.User
		if err := db.Where("email = ?", userData.Email).First(&existingUser).Error; err != nil {
			if err != gorm.ErrRecordNotFound {
				log.Printf("Gagal memeriksa user %s: %v\n", userData.Name, err)
				continue
			}
			
			// User belum ada, cari role-nya
			var role models.Role
			if findRoleErr := db.Where("name = ?", userData.RoleName).First(&role).Error; findRoleErr != nil {
				log.Fatalf("Role '%s' tidak ditemukan: %v\n", userData.RoleName, findRoleErr)
			}

			// Buat user baru
			newUser := models.User{
				Name:     userData.Name,
				Email:    userData.Email,
				Password: string(hashedPassword),
				RoleID:   role.ID,
			}

			if createErr := db.Create(&newUser).Error; createErr != nil {
				log.Fatalf("Gagal menanamkan user %s: %v\n", userData.Name, createErr)
			}
			fmt.Printf("✅ User '%s' dengan role '%s' berhasil ditanamkan.\n", newUser.Name, userData.RoleName)
		} else {
			fmt.Printf("🟡 User '%s' sudah ada, dilewati.\n", userData.Name)
		}
	}
}
