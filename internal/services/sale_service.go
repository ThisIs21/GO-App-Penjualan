package services

import (
	"errors"
	"gorm.io/gorm"
	"time"
	"app-penjualan/internal/dto"
	"app-penjualan/internal/models"
	"app-penjualan/internal/repositories"
)

type SaleService struct {
	db    *gorm.DB
	srepo *repositories.SaleRepo
	prod  *ProductService
}

func NewSaleService(db *gorm.DB) *SaleService {
	return &SaleService{db: db, srepo: repositories.NewSaleRepo(db), prod: NewProductService(db)}
}

func (s *SaleService) Create(uid uint, req dto.CreateSaleReq) (*models.Sale, error) {
	var saleDate time.Time
	if req.Date != nil {
		saleDate = *req.Date
	} else {
		saleDate = time.Now()
	}

	sale := &models.Sale{UserID: uid, Date: saleDate}

	if req.CustomerName != nil && *req.CustomerName != "" {
		newCustomer := models.Customer{Name: *req.CustomerName}
		if err := s.db.Create(&newCustomer).Error; err != nil {
			return nil, errors.New("gagal membuat pelanggan baru")
		}
		sale.CustomerID = &newCustomer.ID
	} else if req.CustomerID != nil {
		sale.CustomerID = req.CustomerID
	}

	var subtotal float64
	var saleItems []models.SaleItem

	err := s.db.Transaction(func(tx *gorm.DB) error {
		for _, it := range req.Items {
			var product models.Product
			if err := tx.First(&product, it.ProductID).Error; err != nil {
				return errors.New("produk tidak ditemukan")
			}

			itemSubtotal := float64(it.Qty) * product.SellPrice
			item := models.SaleItem{
				ProductID: it.ProductID,
				Qty:       it.Qty,
				Price:     product.SellPrice,
				Subtotal:  itemSubtotal,
			}

			saleItems = append(saleItems, item)
			subtotal += itemSubtotal

			if err := s.prod.AdjustStock(it.ProductID, -it.Qty); err != nil {
				return errors.New("stok tidak cukup")
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	sale.Subtotal = subtotal

	var discountAmount float64
	if req.VoucherCode != nil {
		var v models.Voucher
		if err := s.db.Where("code = ? AND active = ?", *req.VoucherCode, 1).First(&v).Error; err == nil {
			sale.VoucherID = &v.ID
			switch v.Type {
			case models.VoucherPercent:
				discountAmount = (v.Value / 100.0) * subtotal
			case models.VoucherAmount:
				discountAmount = v.Value
			}
		}
	}

	sale.Discount = discountAmount
	total := subtotal - discountAmount
	sale.Total = total

	if req.PaidAmount == nil || *req.PaidAmount < total {
		return nil, errors.New("paid_amount harus lebih besar atau sama dengan total")
	}
	sale.Paid = *req.PaidAmount
	sale.Change = *req.PaidAmount - total

	sale.Items = saleItems

	if err := s.srepo.Create(sale); err != nil {
		return nil, err
	}

	return sale, nil
}

func (s *SaleService) LoadDetails(id uint) (*models.Sale, error) {
	var sale models.Sale
	err := s.db.Preload("Customer").Preload("User").Preload("Voucher").
		Preload("Items.Product").First(&sale, id).Error
	if err != nil {
		return nil, err
	}
	return &sale, nil
}