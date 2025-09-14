
package repositories

import (
	"gorm.io/gorm"
	"app-penjualan/internal/models"
)

type SaleRepo struct{ db *gorm.DB }
func NewSaleRepo(db *gorm.DB) *SaleRepo { return &SaleRepo{db} }

func (r *SaleRepo) Create(s *models.Sale) error { 
	return r.db.Create(s).Error 
}

func (r *SaleRepo) CreateWithTx(tx *gorm.DB, s *models.Sale) error {
	return tx.Create(s).Error
}

func (r *SaleRepo) WithItems(id uint) (*models.Sale, error) {
	var m models.Sale
	err := r.db.Preload("Customer").Preload("User").Preload("Voucher").
		Preload("Items.Product").First(&m, id).Error
	return &m, err
}
