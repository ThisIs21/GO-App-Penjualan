package models

type RoleName string

const (
	RoleOwner        RoleName = "OWNER"
	RoleKasir        RoleName = "KASIR"
	RolePembelian    RoleName = "PEMBELIAN"
	RoleGudang       RoleName = "GUDANG"
	RoleKepalaGudang RoleName = "KEPALA_GUDANG"
)

type Role struct {
	ID   uint     `gorm:"primaryKey"`
	Name RoleName `gorm:"uniqueIndex;size:32"`
	Audit
}

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:100;not null"`
	Email    string `gorm:"size:120;uniqueIndex;not null"`
	Password string `gorm:"size:255;not null"`
	RoleID   uint
	Role     Role
	Status	 string
	Audit
}
