package models

import "time"

type PurchaseStatus string

const (
	PurchaseDraft    PurchaseStatus = "DRAFT"
	PurchaseSubmitted PurchaseStatus = "SUBMITTED"
	PurchaseApproved  PurchaseStatus = "APPROVED"
	PurchaseRejected  PurchaseStatus = "REJECTED"
)

type Purchase struct {
	ID         uint           `gorm:"primaryKey"`
	SupplierID uint           `gorm:"not null"`
	Supplier   Supplier
	UserID     uint // dibuat oleh
	User       User
	ApprovedBy *uint
	Approver   *User `gorm:"foreignKey:ApprovedBy"`
	ApprovedAt *time.Time 
	Total      float64    `gorm:"not null;default:0"`
	Status     PurchaseStatus `gorm:"size:16;not null;default:'DRAFT'"`
	Date       time.Time  `gorm:"not null"`
	Items      []PurchaseItem
	Audit
}

type PurchaseItem struct {
	ID         uint `gorm:"primaryKey"`
	PurchaseID uint
	ProductID  uint
	Product    Product
	Qty        int
	Price      float64
	Subtotal   float64
}
