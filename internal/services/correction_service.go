package services

import (
	"encoding/json"
	"errors"

	"gorm.io/gorm"
	"app-penjualan/internal/models"
)

type CorrectionService struct{ DB *gorm.DB }
func NewCorrectionService(db *gorm.DB)*CorrectionService{ return &CorrectionService{DB:db} }

type CorrectionInput struct{
	RefTable string `json:"ref_table" binding:"required"`
	RefID    uint   `json:"ref_id" binding:"required"`
	After    map[string]any `json:"after" binding:"required"`
	Reason   string `json:"reason"`
	DeltaStock int  `json:"delta_stock"` // contoh: +2 menambah stok; -3 mengurangi
	UserID   uint
}

func (s *CorrectionService) Create(input CorrectionInput) (*models.Correction, error){
	// Ambil "before" sebagai JSON snapshot record existing
	var before map[string]any
	if err := s.DB.Table(input.RefTable).Where("id = ?", input.RefID).Take(&before).Error; err!=nil{
		return nil, err
	}
	beforeJSON, _ := json.Marshal(before)
	afterJSON, _  := json.Marshal(input.After)

	// Apply perubahan ke record target
	if err := s.DB.Table(input.RefTable).Where("id = ?", input.RefID).Updates(input.After).Error; err!=nil{
		return nil, err
	}

	// Jika perlu update stok (misal koreksi qty), terapkan delta ke tabel products
	if input.DeltaStock != 0 {
		if err := s.DB.Exec(`UPDATE products SET stok = stok + ? WHERE id = ?`, input.DeltaStock, input.After["product_id"]).Error; err != nil {
			return nil, errors.New("failed apply delta stock: "+err.Error())
		}
	}

	c := &models.Correction{
		RefTable:   input.RefTable,
		RefID:      input.RefID,
		BeforeJSON: string(beforeJSON),
		AfterJSON:  string(afterJSON),
		DeltaStock: input.DeltaStock,
		Reason:     input.Reason,
		UserID:     input.UserID,
	}
	if err := s.DB.Create(c).Error; err!=nil { return nil, err }
	return c, nil
}
