package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"app-penjualan/internal/services"
)

type CorrectionHandler struct{ Svc *services.CorrectionService }
func NewCorrectionHandler(s *services.CorrectionService)*CorrectionHandler{ return &CorrectionHandler{Svc:s} }

func (h *CorrectionHandler) Create(c *gin.Context){
	var in services.CorrectionInput
	if err := c.ShouldBindJSON(&in); err!=nil { c.JSON(400, gin.H{"error":err.Error()}); return }
	if uid, ok := c.Get("uid"); ok { in.UserID = uint(uid.(uint)) }
	res, err := h.Svc.Create(in)
	if err!=nil { c.JSON(500, gin.H{"error":err.Error()}); return }
	c.JSON(http.StatusOK, gin.H{"data":res})
}
