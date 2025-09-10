package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"app-penjualan/internal/middlewares"
	"app-penjualan/internal/services"
)

type ReturnHandler struct {
	svc *services.ReturnService
}

func NewReturnHandler() *ReturnHandler {
	return &ReturnHandler{svc: services.NewReturnService()}
}

/* ===== Retur Penjualan ===== */

func (h *ReturnHandler) CreateSaleReturn(c *gin.Context) {
	userID := middlewares.MustUserID(c)
	var dto services.CreateSaleReturnDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ret, err := h.svc.CreateSaleReturn(userID, dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, ret)
}

func (h *ReturnHandler) ApproveSaleReturn(c *gin.Context) {
	approver := middlewares.MustUserID(c)
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.svc.ApproveSaleReturn(uint(id), approver); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "retur penjualan disetujui"})
}

/* ===== Retur Pembelian ===== */

func (h *ReturnHandler) CreatePurchaseReturn(c *gin.Context) {
	userID := middlewares.MustUserID(c)
	var dto services.CreatePurchaseReturnDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ret, err := h.svc.CreatePurchaseReturn(userID, dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, ret)
}

func (h *ReturnHandler) ApprovePurchaseReturn(c *gin.Context) {
	approver := middlewares.MustUserID(c)
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.svc.ApprovePurchaseReturn(uint(id), approver); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "retur pembelian disetujui"})
}
