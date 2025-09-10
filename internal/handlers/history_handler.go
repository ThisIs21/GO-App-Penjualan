package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"app-penjualan/internal/services"
	"app-penjualan/internal/utils"
)

type HistoryHandler struct{ Svc *services.HistoryService }
func NewHistoryHandler(s *services.HistoryService) *HistoryHandler { return &HistoryHandler{Svc: s} }

func getPaging(c *gin.Context) (int, int, string) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "20"))
	search := c.DefaultQuery("q", "")
	return page, size, search
}

func (h *HistoryHandler) Purchases(c *gin.Context) {
	dr, err := utils.ParseDateRange(c); if err != nil { c.JSON(400, gin.H{"error":"invalid date format (use YYYY-MM-DD)"}); return }
	page,size,search := getPaging(c)
	data,total,err := h.Svc.Purchases(dr,page,size,search)
	if err!=nil { c.JSON(500, gin.H{"error":err.Error()}); return }
	c.JSON(http.StatusOK, gin.H{"data":data,"total":total,"page":page,"size":size})
}
// duplikasi pola yang sama:
func (h *HistoryHandler) Sales(c *gin.Context){ dr,_:=utils.ParseDateRange(c); p,s,q:=getPaging(c); d,t,e:=h.Svc.Sales(dr,p,s,q); if e!=nil{c.JSON(500,gin.H{"error":e.Error()});return}; c.JSON(200,gin.H{"data":d,"total":t,"page":p,"size":s}) }
func (h *HistoryHandler) PurchaseReturns(c *gin.Context){ dr,_:=utils.ParseDateRange(c); p,s,q:=getPaging(c); d,t,e:=h.Svc.PurchaseReturns(dr,p,s,q); if e!=nil{c.JSON(500,gin.H{"error":e.Error()});return}; c.JSON(200,gin.H{"data":d,"total":t,"page":p,"size":s}) }
func (h *HistoryHandler) SaleReturns(c *gin.Context){ dr,_:=utils.ParseDateRange(c); p,s,q:=getPaging(c); d,t,e:=h.Svc.SaleReturns(dr,p,s,q); if e!=nil{c.JSON(500,gin.H{"error":e.Error()});return}; c.JSON(200,gin.H{"data":d,"total":t,"page":p,"size":s}) }
func (h *HistoryHandler) StockOpnames(c *gin.Context){ dr,_:=utils.ParseDateRange(c); p,s,q:=getPaging(c); d,t,e:=h.Svc.StockOpnames(dr,p,s,q); if e!=nil{c.JSON(500,gin.H{"error":e.Error()});return}; c.JSON(200,gin.H{"data":d,"total":t,"page":p,"size":s}) }
