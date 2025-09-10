package services

import (
	"gorm.io/gorm"
	"app-penjualan/internal/dto"
	"app-penjualan/internal/models"
	"app-penjualan/internal/repositories"
)

type StockService struct{
	db *gorm.DB
	prod *ProductService
}

func NewStockService(db *gorm.DB) *StockService { return &StockService{db: db, prod: NewProductService(db)} }

func (s *StockService) DoStockOpname(uid uint, req dto.CreateOpnameReq) (*models.StockOpname, error) {
	op := &models.StockOpname{UserID: uid, Date: req.Date, Note: req.Note}
	for _, it := range req.Items {
		var p models.Product; s.db.First(&p, it.ProductID)
		op.Items = append(op.Items, models.StockOpnameItem{
			ProductID: it.ProductID, QtySystem: p.Stock, QtyPhysical: it.QtyPhysical,
		})
		if delta := it.QtyPhysical - p.Stock; delta != 0 {
			_ = s.prod.AdjustStock(it.ProductID, delta)
		}
	}
	repo := repositories.NewStockRepo(s.db)
	if err := repo.Create(op); err != nil { return nil, err }
	return op, nil
}
