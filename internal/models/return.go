package models

import "time"

type PurchaseReturn struct {
	ID           uint `gorm:"primaryKey"`
	PurchaseID   uint
	UserID       uint // dibuat oleh
	Total        float64
	Date         time.Time
	ApprovedBy   *uint
	Approver     *User  `gorm:"foreignKey:ApprovedBy"`
	ApprovedAt   *time.Time
	Status       string `gorm:"size:16;default:'SUBMITTED'"` // APPROVED/REJECTED
	Items        []PurchaseReturnItem
	Audit
}

type PurchaseReturnItem struct {
	ID          uint `gorm:"primaryKey"`
	PurchaseReturnID uint
	ProductID   uint
	Product     Product
	Qty         int
	Price       float64
	Subtotal    float64
}

type SaleReturn struct {
	ID         uint `gorm:"primaryKey"`
	SaleID     uint
	UserID     uint
	Total      float64
	Date       time.Time
	ApprovedBy *uint
	Approver   *User `gorm:"foreignKey:ApprovedBy"`
	ApprovedAt *time.Time
	Status     string `gorm:"size:16;default:'SUBMITTED'"`
	Items      []SaleReturnItem
	Audit
}

type SaleReturnItem struct {
	ID          uint `gorm:"primaryKey"`
	SaleReturnID uint
	ProductID   uint
	Product     Product
	Qty         int
	Price       float64
	Subtotal    float64
}
