package services

import (
	"app-penjualan/internal/dto"
	"app-penjualan/internal/models"
	"errors"
	"gorm.io/gorm"
)

type UnitService struct {
	DB *gorm.DB
}

func NewUnitService(db *gorm.DB) *UnitService {
	return &UnitService{DB: db}
}

func (s *UnitService) Create(req dto.CreateUnitReq) (*models.Unit, error) {
	unit := models.Unit{Name: req.Name}
	if err := s.DB.Create(&unit).Error; err != nil {
		return nil, errors.New("gagal membuat unit")
	}
	return &unit, nil
}

func (s *UnitService) List() ([]models.Unit, error) {
	var units []models.Unit
	if err := s.DB.Find(&units).Error; err != nil {
		return nil, errors.New("gagal mengambil daftar unit")
	}
	return units, nil
}

func (s *UnitService) Get(id uint) (*models.Unit, error) {
	var unit models.Unit
	if err := s.DB.First(&unit, id).Error; err != nil {
		return nil, errors.New("unit tidak ditemukan")
	}
	return &unit, nil
}

func (s *UnitService) Update(id uint, req dto.UpdateUnitReq) (*models.Unit, error) {
	var unit models.Unit
	if err := s.DB.First(&unit, id).Error; err != nil {
		return nil, errors.New("unit tidak ditemukan")
	}
	if req.Name != nil {
		unit.Name = *req.Name
	}
	if err := s.DB.Save(&unit).Error; err != nil {
		return nil, errors.New("gagal memperbarui unit")
	}
	return &unit, nil
}

func (s *UnitService) Delete(id uint) error {
	if err := s.DB.Delete(&models.Unit{}, id).Error; err != nil {
		return errors.New("gagal menghapus unit")
	}
	return nil
}