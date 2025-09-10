package dto

import "app-penjualan/internal/models"

// =====================
// VOUCHER DTOs
// =====================
type CreateVoucherReq struct {
	Code   string            `json:"code" validate:"required"`
	Type   models.VoucherType `json:"type" validate:"required"`
	Value  float64           `json:"value" validate:"required"`
	Active *bool             `json:"active" validate:"required"`
}

type UpdateVoucherReq struct {
	Code   *string            `json:"code"`
	Type   *models.VoucherType `json:"type"`
	Value  *float64           `json:"value"`
	Active *bool             `json:"active"`
}

// =====================
// UNIT DTOs
// =====================
type CreateUnitReq struct {
	Name string `json:"name" validate:"required"`
}

type UpdateUnitReq struct {
	Name *string `json:"name"`
}

// =====================
// SUPPLIER DTOs
// =====================
type CreateSupplierReq struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
	Phone   string `json:"phone" validate:"required"`
}

type UpdateSupplierReq struct {
	Name    *string `json:"name"`
	Address *string `json:"address"`
	Phone   *string `json:"phone"`
}

// =====================
// CATEGORY DTOs
// =====================
type CreateCategoryReq struct {
	Name string `json:"name" validate:"required"`
}

type UpdateCategoryReq struct {
	Name *string `json:"name"`
}