package services

import (
	"app-penjualan/internal/models"
	"app-penjualan/internal/utils"
	"gorm.io/gorm"
)

type HistoryService struct {
	DB *gorm.DB
}

func NewHistoryService(db *gorm.DB) *HistoryService {
	return &HistoryService{DB: db}
}

// Generic date filter helper
func applyRange(q *gorm.DB, dr utils.DateRange, field string) *gorm.DB {
	if dr.From != nil {
		q = q.Where(field+" >= ?", *dr.From)
	}
	if dr.To != nil {
		q = q.Where(field+" <= ?", *dr.To)
	}
	return q
}

// Note: You need to have a `Paginate` function defined somewhere for this to work.
// It is likely in a separate file (e.g., pagination.go) in the same services package.

func (s *HistoryService) Purchases(dr utils.DateRange, page, size int, search string) (any, int64, error) {
	q := s.DB.Table("purchases p").
		Select(`p.id, p.supplier_id, p.user_id, p.total, p.tanggal`).
		Scopes(Paginate(page, size))
	if search != "" {
		q = q.Where("CAST(p.id as CHAR) LIKE ? OR CAST(p.total as CHAR) LIKE ?", "%"+search+"%", "%"+search+"%")
	}
	q = applyRange(q, dr, "p.tanggal")
	var rows []map[string]any
	var total int64
	if err := q.Count(&total).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *HistoryService) Sales(dr utils.DateRange, page, size int, search string) ([]models.Sale, int64, error) {
	var sales []models.Sale
	var total int64

	// Basis query
	q := s.DB.Model(&models.Sale{})

	// Filter pencarian
	if search != "" {
		q = q.
			Joins("left join customers on customers.id = sales.customer_id").
			Joins("left join sale_items on sale_items.sale_id = sales.id").
			Joins("left join products on products.id = sale_items.product_id").
			Where("LOWER(customers.name) LIKE ? OR LOWER(products.name) LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	// Filter berdasarkan tanggal
	if dr.From != nil {
		q = q.Where("sales.date >= ?", *dr.From)
	}
	if dr.To != nil {
		q = q.Where("sales.date <= ?", *dr.To)
	}

	// Count total sebelum melakukan paginasi
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Terapkan paginasi
	q = q.Scopes(Paginate(page, size))

	// Muat relasi setelah paginasi
	q = q.Preload("Customer").
		Preload("Voucher").
		Preload("Items.Product")

	// Eksekusi query
	if err := q.Find(&sales).Error; err != nil {
		return nil, 0, err
	}

	return sales, total, nil
}

func (s *HistoryService) PurchaseReturns(dr utils.DateRange, page, size int, search string) (any, int64, error) {
	q := s.DB.Table("purchase_returns pr").
		Select(`pr.id, pr.purchase_id, pr.reason, pr.tanggal`).
		Scopes(Paginate(page, size))
	if search != "" {
		q = q.Where("CAST(pr.id as CHAR) LIKE ? OR pr.reason LIKE ?", "%"+search+"%", "%"+search+"%")
	}
	q = applyRange(q, dr, "pr.tanggal")
	var rows []map[string]any
	var total int64
	if err := q.Count(&total).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *HistoryService) SaleReturns(dr utils.DateRange, page, size int, search string) (any, int64, error) {
	q := s.DB.Table("sale_returns sr").
		Select(`sr.id, sr.sale_id, sr.reason, sr.tanggal`).
		Scopes(Paginate(page, size))
	if search != "" {
		q = q.Where("CAST(sr.id as CHAR) LIKE ? OR sr.reason LIKE ?", "%"+search+"%", "%"+search+"%")
	}
	q = applyRange(q, dr, "sr.tanggal")
	var rows []map[string]any
	var total int64
	if err := q.Count(&total).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *HistoryService) StockOpnames(dr utils.DateRange, page, size int, search string) (any, int64, error) {
	q := s.DB.Table("stock_opnames so").
		Select(`so.id, so.user_id, so.tanggal, so.catatan`).
		Scopes(Paginate(page, size))
	if search != "" {
		q = q.Where("CAST(so.id as CHAR) LIKE ? OR so.catatan LIKE ?", "%"+search+"%", "%"+search+"%")
	}
	q = applyRange(q, dr, "so.tanggal")
	var rows []map[string]any
	var total int64
	if err := q.Count(&total).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}