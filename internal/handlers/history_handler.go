package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"app-penjualan/internal/services"
	"app-penjualan/internal/utils"
)

type HistoryHandler struct {
	Svc *services.HistoryService
}

func NewHistoryHandler(s *services.HistoryService) *HistoryHandler {
	return &HistoryHandler{Svc: s}
}

func getPaging(c *gin.Context) (int, int, string) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
	search := c.DefaultQuery("q", "")
	return page, size, search
}

func (h *HistoryHandler) Purchases(c *gin.Context) {
	dr, err := utils.ParseDateRange(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	page, size, search := getPaging(c)
	data, total, err := h.Svc.Purchases(dr, page, size, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data, "total": total, "page": page, "size": size})
}

func (h *HistoryHandler) Sales(c *gin.Context) {
	dr, err := utils.ParseDateRange(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	page, size, search := getPaging(c)
	data, total, totalValue, err := h.Svc.Sales(dr, page, size, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":        data,
		"total":       total,
		"total_value": totalValue,
		"page":        page,
		"size":        size,
	})
}

func (h *HistoryHandler) PurchaseReturns(c *gin.Context) {
	dr, err := utils.ParseDateRange(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	page, size, search := getPaging(c)
	data, total, err := h.Svc.PurchaseReturns(dr, page, size, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data, "total": total, "page": page, "size": size})
}

func (h *HistoryHandler) SaleReturns(c *gin.Context) {
	dr, err := utils.ParseDateRange(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	page, size, search := getPaging(c)
	data, total, err := h.Svc.SaleReturns(dr, page, size, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data, "total": total, "page": page, "size": size})
}

func (h *HistoryHandler) StockOpnames(c *gin.Context) {
	dr, err := utils.ParseDateRange(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	page, size, search := getPaging(c)
	data, total, err := h.Svc.StockOpnames(dr, page, size, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data, "total": total, "page": page, "size": size})
}