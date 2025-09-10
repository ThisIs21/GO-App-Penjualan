package models

import "time"

type Audit struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

type CorrectionType string
const (
	CorrectionPurchase CorrectionType = "PURCHASE"
	CorrectionSale     CorrectionType = "SALE"
	CorrectionPRet     CorrectionType = "PURCHASE_RETURN"
	CorrectionSRet     CorrectionType = "SALE_RETURN"
)

type Correction struct {
	ID         uint           `gorm:"primaryKey"`
	RefTable   string         // nama tabel target (purchases/sales/dll)
	RefID      uint           // id record yang dikoreksi
	BeforeJSON string         `gorm:"type:json"`
	AfterJSON  string         `gorm:"type:json"`
	DeltaStock int            // perubahan stok (positif/negatif)
	Reason     string
	UserID     uint
	CreatedAt  time.Time
}