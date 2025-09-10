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
	// 1. Tentukan tanggal transaksi: gunakan `req.Date` jika ada, jika tidak, gunakan `time.Now()`
	var saleDate time.Time
	if req.Date != nil {
		saleDate = *req.Date
	} else {
		saleDate = time.Now()
	}

	// 2. Inisialisasi objek `Sale`
	sale := &models.Sale{UserID: uid, Date: saleDate}

	// Tangani pelanggan
	if req.CustomerName != nil && *req.CustomerName != "" {
		newCustomer := models.Customer{Name: *req.CustomerName}
		if err := s.db.Create(&newCustomer).Error; err != nil {
			return nil, errors.New("gagal membuat pelanggan baru")
		}
		sale.CustomerID = &newCustomer.ID
	} else if req.CustomerID != nil {
		sale.CustomerID = req.CustomerID
	}

	// Hitung total harga terlebih dahulu
	var subtotal float64
	var saleItems []models.SaleItem

	for _, it := range req.Items {
		var product models.Product
		// Ambil harga produk dari database
		if err := s.db.First(&product, it.ProductID).Error; err != nil {
			return nil, errors.New("produk tidak ditemukan")
		}

		// Hitung subtotal secara otomatis
		itemSubtotal := float64(it.Qty) * product.SellPrice

		item := models.SaleItem{
			ProductID: it.ProductID,
			Qty:       it.Qty,
			Price:     product.SellPrice,
			Subtotal:  itemSubtotal,
		}

		saleItems = append(saleItems, item)
		subtotal += itemSubtotal

		// Kurangi stok
		if err := s.prod.AdjustStock(it.ProductID, -it.Qty); err != nil {
			return nil, errors.New("stok tidak cukup")
		}
	}
	
	sale.Subtotal = subtotal

	// Tangani voucher dan hitung diskon
	var discountAmount float64
	if req.VoucherCode != nil {
		var v models.Voucher
		if err := s.db.Where("code = ? AND active = ?", *req.VoucherCode, 1).First(&v).Error; err == nil {
			sale.VoucherID = &v.ID
			// --- PERBAIKAN: Gunakan switch untuk menangani berbagai tipe voucher ---
			switch v.Type {
			case models.VoucherPercent:
				// Hitung diskon dalam Rupiah
				discountAmount = (v.Value / 100.0) * subtotal
			case models.VoucherAmount:
				// Diskon langsung berupa nilai
				discountAmount = v.Value
			}
			// --- AKHIR PERBAIKAN ---
		}
	}
	
	// Perbaikan utama: Simpan nilai diskon ke struct `sale`
	sale.Discount = discountAmount
	
	// Hitung total akhir setelah diskon
	total := subtotal - discountAmount
	sale.Total = total
	
	// Pastikan paid_amount tidak kosong sebelum menghitung kembalian
	if req.PaidAmount != nil {
		sale.Paid = *req.PaidAmount
		sale.Change = *req.PaidAmount - total
	} else {
		sale.Paid = 0.0 // Atur nilai default jika tidak ada PaidAmount
		sale.Change = 0.0
	}
	
	sale.Items = saleItems

	if err := s.srepo.Create(sale); err != nil {
		return nil, err
	}

	return sale, nil
}
