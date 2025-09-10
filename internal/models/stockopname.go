package models

import "time"

type StockOpname struct {
    ID     uint `gorm:"primaryKey"`
    UserID uint
    User   User
    Date   time.Time
    Note   string `gorm:"size:255"`
    Items  []StockOpnameItem
    Audit
}

type StockOpnameItem struct {
    ID            uint `gorm:"primaryKey"`
    // Perbaikan: Ganti OpnameID menjadi StockOpnameID
    StockOpnameID uint
    ProductID     uint
    Product       Product
    QtySystem     int
    QtyPhysical   int
}