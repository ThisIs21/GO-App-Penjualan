package dto

import "time"

type OpnameItemReq struct {
	ProductID   uint `json:"product_id" validate:"required"`
	QtyPhysical int  `json:"qty_fisik" validate:"gte=0"`
}

type CreateOpnameReq struct {
	Date  time.Time      `json:"date" validate:"required"`
	Note  string         `json:"note"`
	Items []OpnameItemReq `json:"items" validate:"required,dive"`
}
