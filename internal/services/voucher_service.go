package services

import (
	"app-penjualan/internal/dto"
	"app-penjualan/internal/models"
	"errors"
	"gorm.io/gorm"
)

type VoucherService struct {
	DB *gorm.DB
}

func NewVoucherService(db *gorm.DB) *VoucherService {
	return &VoucherService{DB: db}
}

// internal/services/voucher_service.go
// ...
func (s *VoucherService) Create(req dto.CreateVoucherReq) (*models.Voucher, error) {
    // Periksa apakah pointer `Active` tidak nil sebelum dereference
    active := true // Nilai default
    if req.Active != nil {
        active = *req.Active
    }

    voucher := models.Voucher{
        Code:   req.Code,
        Type:   req.Type,
        Value:  req.Value,
        Active: active, // <-- Gunakan variabel `active` yang sudah diperbaiki
    }
    if err := s.DB.Create(&voucher).Error; err != nil {
		return nil, errors.New("gagal membuat voucher")
	}
	return &voucher, nil
}
	


func (s *VoucherService) List() ([]models.Voucher, error) {
	var vouchers []models.Voucher
	if err := s.DB.Find(&vouchers).Error; err != nil {
		return nil, errors.New("gagal mengambil daftar voucher")
	}
	return vouchers, nil
}

func (s *VoucherService) Get(id uint) (*models.Voucher, error) {
	var voucher models.Voucher
	if err := s.DB.First(&voucher, id).Error; err != nil {
		return nil, errors.New("voucher tidak ditemukan")
	}
	return &voucher, nil
}

func (s *VoucherService) Update(id uint, req dto.UpdateVoucherReq) (*models.Voucher, error) {
	var voucher models.Voucher
	if err := s.DB.First(&voucher, id).Error; err != nil {
		return nil, errors.New("voucher tidak ditemukan")
	}
	if req.Code != nil {
		voucher.Code = *req.Code
	}
	if req.Type != nil {
		voucher.Type = *req.Type
	}
	if req.Value != nil {
		voucher.Value = *req.Value
	}
	if req.Active != nil {
		voucher.Active = *req.Active
	}
	if err := s.DB.Save(&voucher).Error; err != nil {
		return nil, errors.New("gagal memperbarui voucher")
	}
	return &voucher, nil
}

func (s *VoucherService) Delete(id uint) error {
	if err := s.DB.Delete(&models.Voucher{}, id).Error; err != nil {
		return errors.New("gagal menghapus voucher")
	}
	return nil
}