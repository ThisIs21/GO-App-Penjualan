package dto

import "time"

// tambahkan PaidAmount ke sini
type CreateSaleReq struct {
	CustomerID *uint      `json:"customer_id,omitempty"`
	CustomerName *string    `json:"customer_name,omitempty"`
	VoucherCode *string    `json:"voucher_code,omitempty"`
	Date         *time.Time `json:"date,omitempty"`
	Items        []SaleItemReq `json:"items" validate:"required"`
	PaidAmount   *float64   `json:"paid_amount" validate:"required,gt=0"`
}

type SaleItemReq struct {
	ProductID uint `json:"product_id" validate:"required"`
	Qty       int  `json:"qty" validate:"required,gt=0"`
}

type GetSaleReq struct {
	StartDate string `query:"start_date"`
	EndDate   string `query:"end_date"`
}
