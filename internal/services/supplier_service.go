package services

import (
	"app-penjualan/internal/dto"
	"app-penjualan/internal/models"
	"errors"
	"gorm.io/gorm"
)

type SupplierService struct {
	DB *gorm.DB
}

func NewSupplierService(db *gorm.DB) *SupplierService {
	return &SupplierService{DB: db}
}

func (s *SupplierService) Create(req dto.CreateSupplierReq) (*models.Supplier, error) {
	supplier := models.Supplier{
		Name:    req.Name,
		Address: req.Address,
		Contact:   req.Phone,
	}
	if err := s.DB.Create(&supplier).Error; err != nil {
		return nil, errors.New("gagal membuat supplier")
	}
	return &supplier, nil
}

func (s *SupplierService) List() ([]models.Supplier, error) {
	var suppliers []models.Supplier
	if err := s.DB.Find(&suppliers).Error; err != nil {
		return nil, errors.New("gagal mengambil daftar supplier")
	}
	return suppliers, nil
}

func (s *SupplierService) Get(id uint) (*models.Supplier, error) {
	var supplier models.Supplier
	if err := s.DB.First(&supplier, id).Error; err != nil {
		return nil, errors.New("supplier tidak ditemukan")
	}
	return &supplier, nil
}

func (s *SupplierService) Update(id uint, req dto.UpdateSupplierReq) (*models.Supplier, error) {
	var supplier models.Supplier
	if err := s.DB.First(&supplier, id).Error; err != nil {
		return nil, errors.New("supplier tidak ditemukan")
	}
	if req.Name != nil {
		supplier.Name = *req.Name
	}
	if req.Address != nil {
		supplier.Address = *req.Address
	}
	if req.Phone != nil {
		supplier.Contact = *req.Phone
	}
	if err := s.DB.Save(&supplier).Error; err != nil {
		return nil, errors.New("gagal memperbarui supplier")
	}
	return &supplier, nil
}

func (s *SupplierService) Delete(id uint) error {
	if err := s.DB.Delete(&models.Supplier{}, id).Error; err != nil {
		return errors.New("gagal menghapus supplier")
	}
	return nil
}