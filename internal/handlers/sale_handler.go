	package handlers

	import (
		"github.com/gin-gonic/gin"
		"app-penjualan/internal/dto"
		"app-penjualan/internal/services"
		"app-penjualan/internal/utils"
	)  

	type SaleHandler struct{ svc *services.SaleService }
	func NewSaleHandler(s *services.SaleService) *SaleHandler { return &SaleHandler{svc: s} }

	func (h *SaleHandler) Register(rg *gin.RouterGroup) {
		rg.POST("/sales", h.Create) // role KASIR
	}

	func (h *SaleHandler) Create(c *gin.Context) {
		var req dto.CreateSaleReq
		if err := c.ShouldBindJSON(&req); err != nil { utils.BadRequest(c, err.Error()); return }
		if err := utils.Validate.Struct(req); err != nil { utils.BadRequest(c, err.Error()); return }
		uid := c.GetUint("uid")
		out, err := h.svc.Create(uid, req)
		if err != nil { utils.BadRequest(c, err.Error()); return }
		utils.Created(c, out)
	}
