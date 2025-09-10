package models
type Category struct {
    ID   uint   `gorm:"primaryKey" json:"id"`
    Name string `gorm:"size:80;uniqueIndex;not null" json:"name"`
    Audit
}

type Unit struct {
    ID   uint   `gorm:"primaryKey" json:"id"`
    Name string `gorm:"size:30;uniqueIndex;not null" json:"name"`
    Audit
}

type Supplier struct {
    ID      uint   `gorm:"primaryKey" json:"id"`
    Name    string `gorm:"size:120;not null" json:"name"`
    Contact string `gorm:"size:120" json:"contact"`
    Address string `gorm:"size:255" json:"address"`
    Audit
}

type Customer struct {
    ID      uint   `gorm:"primaryKey" json:"id"`
    Name    string `gorm:"size:120;not null" json:"name"`
    Contact string `gorm:"size:120" json:"contact"`
    Audit
}

type VoucherType string

const (
	VoucherPercent VoucherType = "PERCENT"
	VoucherAmount  VoucherType = "AMOUNT"
)

type Voucher struct {
	ID     uint        `gorm:"primaryKey"`
	Code   string      `gorm:"size:40;uniqueIndex;not null"`
	Value  float64     `gorm:"not null"` // percent or amount
	Type   VoucherType `gorm:"size:16;not null"`
	Active bool        `gorm:"default:true"`
	Audit
}


type Product struct {
    ID         uint     `gorm:"primaryKey"                              json:"id"`
    Name       string   `gorm:"size:120;uniqueIndex;not null"           json:"name"            binding:"required"`
    CategoryID uint     `gorm:"not null"                                json:"category_id"     binding:"required"`
    Category   Category `json:"category"`
    UnitID     uint     `gorm:"not null"                                json:"unit_id"         binding:"required"`
    Unit       Unit     `json:"unit"`
    CostPrice  float64  `gorm:"type:DECIMAL(18,2);not null"             json:"cost_price"      binding:"required"`
    SellPrice  float64  `gorm:"type:DECIMAL(18,2);not null"             json:"sell_price"      binding:"required"`
    Stock      int      `gorm:"not null;default:0"                      json:"stock"`
    SupplierID uint     `gorm:"not null"                                json:"supplier_id"     binding:"required"`
    Supplier   Supplier `json:"supplier"`
    Audit
}

