package repositories

import (
	"gorm.io/gorm"
	"app-penjualan/internal/models"
)

type StockRepo struct{ db *gorm.DB }
func NewStockRepo(db *gorm.DB) *StockRepo { return &StockRepo{db} }

func (r *StockRepo) Create(op *models.StockOpname) error { return r.db.Create(op).Error }
