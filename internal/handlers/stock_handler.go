package handlers

import (
	"github.com/gin-gonic/gin"
	"app-penjualan/internal/dto"
	"app-penjualan/internal/services"
	"app-penjualan/internal/utils"
)

type StockHandler struct{ svc *services.StockService }
func NewStockHandler(s *services.StockService) *StockHandler { return &StockHandler{svc: s} }

func (h *StockHandler) Register(rg *gin.RouterGroup) {
	rg.POST("/stock-opnames", h.Create) // role GUDANG
}

func (h *StockHandler) Create(c *gin.Context) {
	var req dto.CreateOpnameReq
	if err := c.ShouldBindJSON(&req); err != nil { utils.BadRequest(c, err.Error()); return }
	if err := utils.Validate.Struct(req); err != nil { utils.BadRequest(c, err.Error()); return }
	uid := c.GetUint("uid")
	out, err := h.svc.DoStockOpname(uid, req)
	if err != nil { utils.BadRequest(c, err.Error()); return }
	utils.Created(c, out)
}
