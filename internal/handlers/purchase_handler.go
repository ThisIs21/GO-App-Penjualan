package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"app-penjualan/internal/dto"
	"app-penjualan/internal/services"
	"app-penjualan/internal/utils"
)

type PurchaseHandler struct{ svc *services.PurchaseService }
func NewPurchaseHandler(s *services.PurchaseService) *PurchaseHandler { return &PurchaseHandler{svc: s} }

func (h *PurchaseHandler) Register(rg *gin.RouterGroup) {
	rg.POST("/purchases", h.Create) // role PEMBELIAN
	rg.POST("/purchases/:id/approve", h.Approve) // role KEPALA_GUDANG
	rg.POST("/purchases/:id/reject", h.Reject)
}

func (h *PurchaseHandler) Create(c *gin.Context) {
	var req dto.CreatePurchaseReq
	if err := c.ShouldBindJSON(&req); err != nil { utils.BadRequest(c, err.Error()); return }
	if err := utils.Validate.Struct(req); err != nil { utils.BadRequest(c, err.Error()); return }
	uid := c.GetUint("uid")
	out, err := h.svc.Create(uid, req)
	if err != nil { utils.ServerError(c, err); return }
	utils.Created(c, out)
}
func (h *PurchaseHandler) Approve(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	out, err := h.svc.Approve(c.GetUint("uid"), uint(id), true)
	if err != nil { utils.BadRequest(c, err.Error()); return }
	utils.OK(c, out)
}
func (h *PurchaseHandler) Reject(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	out, err := h.svc.Approve(c.GetUint("uid"), uint(id), false)
	if err != nil { utils.BadRequest(c, err.Error()); return }
	utils.OK(c, out)
}
