package services

import (
	"errors"
	"time"

	"app-penjualan/config"
	"app-penjualan/internal/models"
	"gorm.io/gorm"
)

type ReturnService struct {
	db *gorm.DB
}

func NewReturnService() *ReturnService {
	return &ReturnService{db: config.DB}
}

/* ====== DTO ====== */

type SaleReturnItemDTO struct {
	ProductID uint    `json:"product_id" binding:"required"`
	Qty       int     `json:"qty" binding:"required,gt=0"`
	Price     float64 `json:"price" binding:"required,gte=0"`
}

type CreateSaleReturnDTO struct {
	SaleID uint                `json:"sale_id" binding:"required"`
	Date   time.Time           `json:"date" binding:"required"`
	Items  []SaleReturnItemDTO `json:"items" binding:"required,dive"`
}

type PurchaseReturnItemDTO struct {
	ProductID uint    `json:"product_id" binding:"required"`
	Qty       int     `json:"qty" binding:"required,gt=0"`
	Price     float64 `json:"price" binding:"required,gte=0"`
}

type CreatePurchaseReturnDTO struct {
	PurchaseID uint                    `json:"purchase_id" binding:"required"`
	Date       time.Time               `json:"date" binding:"required"`
	Items      []PurchaseReturnItemDTO `json:"items" binding:"required,dive"`
}

/* ====== Retur Penjualan ====== */

func (s *ReturnService) CreateSaleReturn(userID uint, dto CreateSaleReturnDTO) (*models.SaleReturn, error) {
	if len(dto.Items) == 0 {
		return nil, errors.New("items wajib diisi")
	}

	var sale models.Sale
	if err := s.db.First(&sale, dto.SaleID).Error; err != nil {
		return nil, errors.New("sale tidak ditemukan")
	}

	ret := &models.SaleReturn{
		SaleID: dto.SaleID,
		UserID: userID,
		Date:   dto.Date,
		Status: "PENDING",
	}

	var total float64 = 0 // Ubah tipe data menjadi float64
	for _, it := range dto.Items {
		item := models.SaleReturnItem{
			ProductID: it.ProductID,
			Qty:       it.Qty,
			Price:     it.Price,
			Subtotal:  float64(it.Qty) * it.Price, // Perhitungan menggunakan float64
		}
		ret.Items = append(ret.Items, item)
		total += item.Subtotal
	}
	ret.Total = total

	if err := s.db.Create(ret).Error; err != nil {
		return nil, err
	}
	return ret, nil
}

func (s *ReturnService) ApproveSaleReturn(retID uint, approverID uint) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		var ret models.SaleReturn
		if err := tx.Preload("Items").First(&ret, retID).Error; err != nil {
			return errors.New("retur penjualan tidak ditemukan")
		}
		if ret.Status != "PENDING" {
			return errors.New("retur penjualan sudah diproses")
		}

		// Tambah stok
		for _, it := range ret.Items {
			if err := tx.Model(&models.Product{}).
				Where("id = ?", it.ProductID).
				Update("stock", gorm.Expr("stock + ?", it.Qty)).Error; err != nil {
				return err
			}
		}

		now := time.Now()
		ret.Status = "APPROVED"
		ret.ApprovedBy = &approverID
		ret.ApprovedAt = &now

		return tx.Save(&ret).Error
	})
}

/* ====== Retur Pembelian ====== */

func (s *ReturnService) CreatePurchaseReturn(userID uint, dto CreatePurchaseReturnDTO) (*models.PurchaseReturn, error) {
	if len(dto.Items) == 0 {
		return nil, errors.New("items wajib diisi")
	}

	var pur models.Purchase
	if err := s.db.First(&pur, dto.PurchaseID).Error; err != nil {
		return nil, errors.New("purchase tidak ditemukan")
	}

	ret := &models.PurchaseReturn{
		PurchaseID: dto.PurchaseID,
		UserID:     userID,
		Date:       dto.Date,
		Status:     "PENDING",
	}

	var total float64 = 0 // Ubah tipe data menjadi float64
	for _, it := range dto.Items {
		item := models.PurchaseReturnItem{
			ProductID: it.ProductID,
			Qty:       it.Qty,
			Price:     it.Price,
			Subtotal:  float64(it.Qty) * it.Price, // Perhitungan menggunakan float64
		}
		ret.Items = append(ret.Items, item)
		total += item.Subtotal
	}
	ret.Total = total

	if err := s.db.Create(ret).Error; err != nil {
		return nil, err
	}
	return ret, nil
}

func (s *ReturnService) ApprovePurchaseReturn(retID uint, approverID uint) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		var ret models.PurchaseReturn
		if err := tx.Preload("Items").First(&ret, retID).Error; err != nil {
			return errors.New("retur pembelian tidak ditemukan")
		}
		if ret.Status != "PENDING" {
			return errors.New("retur pembelian sudah diproses")
		}

		// Kurangi stok
		for _, it := range ret.Items {
			// Validasi stok cukup
			var prod models.Product
			if err := tx.First(&prod, it.ProductID).Error; err != nil {
				return err
			}
			if prod.Stock < it.Qty {
				return errors.New("stok tidak cukup untuk retur pembelian")
			}
			if err := tx.Model(&models.Product{}).
				Where("id = ?", it.ProductID).
				Update("stock", gorm.Expr("stock - ?", it.Qty)).Error; err != nil {
				return err
			}
		}

		now := time.Now()
		ret.Status = "APPROVED"
		ret.ApprovedBy = &approverID
		ret.ApprovedAt = &now

		return tx.Save(&ret).Error
	})
}