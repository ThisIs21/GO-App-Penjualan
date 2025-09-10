
package dto

import "time"

type PurchaseItemReq struct {
	ProductID uint    `json:"product_id" validate:"required"`
	Qty       int     `json:"qty" validate:"gt=0"`
	Price     float64 `json:"price" validate:"gt=0"`
}

type CreatePurchaseReq struct {
	SupplierID uint              `json:"supplier_id" validate:"required"`
	Date       time.Time         `json:"date" validate:"required"`
	Items      []PurchaseItemReq `json:"items" validate:"required,dive"`
	Submit     bool              `json:"submit"`
}
