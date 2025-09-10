package services

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"app-penjualan/config"
	"app-penjualan/internal/dto"
	"app-penjualan/internal/models"
	"app-penjualan/internal/utils"
)

type AuthService struct {
	db  *gorm.DB
	cfg *config.AppConfig
}

func NewAuthService(db *gorm.DB, cfg *config.AppConfig) *AuthService {
	return &AuthService{db, cfg}
}

// Login finds a user by email and checks their password.
func (s *AuthService) Login(email, password string) (string, string, string, error) {
	var u models.User
	if err := s.db.Preload("Role").Where("email = ?", email).First(&u).Error; err != nil {
		return "", "", "", errors.New("user not found")
	}
	if !utils.CheckPassword(u.Password, password) {
		return "", "", "", errors.New("wrong password")
	}
	tok, err := utils.GenerateJWT(s.cfg.JWTSecret, s.cfg.JWTTTLHrs, u.ID, u.Name, string(u.Role.Name))
	return tok, u.Name, string(u.Role.Name), err
}

// CreateUser creates a new user, requiring an existing RoleID.
func (s *AuthService) CreateUser(name, email, pass string, roleID uint, status string) error {
	hash, _ := utils.HashPassword(pass)
	u := models.User{
		Name:     name,
		Email:    email,
		Password: hash,
		RoleID:   roleID,
		Status:   status,
	}
	return s.db.Create(&u).Error
}

// ListUsers retrieves all users from the database.
func (s *AuthService) ListUsers() ([]models.User, error) {
	var users []models.User
	if err := s.db.Preload("Role").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// UpdateUser updates an existing user's details.
func (s *AuthService) UpdateUser(id uint, req dto.UpdateUserRequest) error {
	var user models.User
	if err := s.db.First(&user, id).Error; err != nil {
		return errors.New("user not found")
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		var existingUser models.User
		if err := s.db.Where("email = ? AND id <> ?", req.Email, id).First(&existingUser).Error; err == nil {
			return errors.New("email is already in use by another user")
		}
		user.Email = req.Email
	}
	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	}
	if req.Role != "" {
		var role models.Role
		if err := s.db.Where("name = ?", req.Role).First(&role).Error; err != nil {
			return errors.New("invalid role")
		}
		user.RoleID = role.ID
	}
	if req.Status != "" {
		user.Status = req.Status
	}

	if err := s.db.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

// DeleteUser deletes a user by their ID.
func (s *AuthService) DeleteUser(id uint) error {
	var user models.User
	if err := s.db.First(&user, id).Error; err != nil {
		return errors.New("user not found")
	}

	if err := s.db.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

// ListRoles retrieves all roles from the database.
func (s *AuthService) ListRoles() ([]models.Role, error) {
	var roles []models.Role
	if err := s.db.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}