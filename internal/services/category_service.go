package services

import (
	"app-penjualan/internal/dto"
	"app-penjualan/internal/models"
	"errors"
	"gorm.io/gorm"
)

type CategoryService struct {
	DB *gorm.DB
}

func NewCategoryService(db *gorm.DB) *CategoryService {
	return &CategoryService{DB: db}
}

func (s *CategoryService) Create(req dto.CreateCategoryReq) (*models.Category, error) {
	category := models.Category{Name: req.Name}
	if err := s.DB.Create(&category).Error; err != nil {
		return nil, errors.New("gagal membuat kategori")
	}
	return &category, nil
}

func (s *CategoryService) List() ([]models.Category, error) {
	var categories []models.Category
	if err := s.DB.Find(&categories).Error; err != nil {
		return nil, errors.New("gagal mengambil daftar kategori")
	}
	return categories, nil
}

func (s *CategoryService) Get(id uint) (*models.Category, error) {
	var category models.Category
	if err := s.DB.First(&category, id).Error; err != nil {
		return nil, errors.New("kategori tidak ditemukan")
	}
	return &category, nil
}

func (s *CategoryService) Update(id uint, req dto.UpdateCategoryReq) (*models.Category, error) {
	var category models.Category
	if err := s.DB.First(&category, id).Error; err != nil {
		return nil, errors.New("kategori tidak ditemukan")
	}
	if req.Name != nil {
		category.Name = *req.Name
	}
	if err := s.DB.Save(&category).Error; err != nil {
		return nil, errors.New("gagal memperbarui kategori")
	}
	return &category, nil
}

func (s *CategoryService) Delete(id uint) error {
	if err := s.DB.Delete(&models.Category{}, id).Error; err != nil {
		return errors.New("gagal menghapus kategori")
	}
	return nil
}