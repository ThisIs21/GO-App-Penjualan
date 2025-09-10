package models

import (
    "time"
)

type Sale struct {
    ID           uint `gorm:"primaryKey"`
    CustomerID   *uint
    Customer     *Customer
    UserID       uint // kasir
    User         User
    VoucherID    *uint
    Voucher      *Voucher
    
    Subtotal     float64 `gorm:"not null"`
    Discount     float64 `gorm:"not null"`
    Total        float64 `gorm:"not null;default:0"`
    Paid         float64 `gorm:"not null"`
    Change       float64 `gorm:"not null"`
    
    Date         time.Time
    Items        []SaleItem
    Audit
}

type SaleItem struct {
    ID        uint `gorm:"primaryKey"`
    SaleID    uint
    ProductID uint
    Product   Product
    Qty       int
    Price     float64
    Subtotal  float64
}
