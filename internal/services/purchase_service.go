
package services

import (
	"errors"
	"time"

	"gorm.io/gorm"
	"app-penjualan/internal/dto"
	"app-penjualan/internal/models"
	"app-penjualan/internal/repositories"
)

type PurchaseService struct {
	db    *gorm.DB
	prepo *repositories.PurchaseRepo
	prod  *ProductService
}

func NewPurchaseService(db *gorm.DB) *PurchaseService {
	return &PurchaseService{db: db, prepo: repositories.NewPurchaseRepo(db), prod: NewProductService(db)}
}

func (s *PurchaseService) Create(uid uint, req dto.CreatePurchaseReq) (*models.Purchase, error) {
	// Validasi SupplierID
	var supplier models.Supplier
	if err := s.db.First(&supplier, req.SupplierID).Error; err != nil {
		return nil, err // Mengembalikan error jika supplier tidak ditemukan
	}

	// Validasi ProductID di setiap item
	for _, it := range req.Items {
		var product models.Product
		if err := s.db.First(&product, it.ProductID).Error; err != nil {
			return nil, err // Mengembalikan error jika produk tidak ditemukan
		}
	}

	// Validasi tanggal
	if req.Date.IsZero() {
		return nil, errors.New("invalid date")
	}

	p := &models.Purchase{
		SupplierID: req.SupplierID,
		UserID:     uid,
		Date:       req.Date,
		Status:     models.PurchaseDraft,
	}

	var total float64
	for _, it := range req.Items {
		item := models.PurchaseItem{
			ProductID: it.ProductID,
			Qty:       it.Qty,
			Price:     it.Price,
			Subtotal:  float64(it.Qty) * it.Price,
		}
		p.Items = append(p.Items, item)
		total += item.Subtotal
	}

	p.Total = total
	if req.Submit {
		p.Status = models.PurchaseSubmitted
	}

	if err := s.prepo.Create(p); err != nil {
		return nil, err
	}

	return p, nil
}

func (s *PurchaseService) Approve(approverID, purchaseID uint, approve bool) (*models.Purchase, error) {
	po, err := s.prepo.WithItems(purchaseID)
	if err != nil {
		return nil, err
	}
	if approve {
		now := time.Now()
		po.Status = models.PurchaseApproved
		po.ApprovedBy = &approverID
		po.Approver = &models.User{ID: approverID}
		po.ApprovedAt = &now
		// Tambah stok
		for _, it := range po.Items {
			_ = s.prod.AdjustStock(it.ProductID, it.Qty)
		}
	} else {
		po.Status = models.PurchaseRejected
	}
	if err := s.prepo.Update(po); err != nil {
		return nil, err
	}
	return po, nil
}
