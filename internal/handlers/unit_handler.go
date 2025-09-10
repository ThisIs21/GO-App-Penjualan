package handlers

import (
	"app-penjualan/internal/dto"
	"app-penjualan/internal/services"
	"app-penjualan/internal/utils"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type UnitHandler struct {
	Svc *services.UnitService
}

func NewUnitHandler(svc *services.UnitService) *UnitHandler {
	return &UnitHandler{Svc: svc}
}

func (h *UnitHandler) Create(c *gin.Context) {
	var req dto.CreateUnitReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	unit, err := h.Svc.Create(req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusCreated, "Unit berhasil ditambahkan", unit)
}

func (h *UnitHandler) List(c *gin.Context) {
	units, err := h.Svc.List()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Daftar unit berhasil diambil", units)
}

func (h *UnitHandler) Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "ID tidak valid")
		return
	}
	unit, err := h.Svc.Get(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Unit tidak ditemukan")
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Detail unit berhasil diambil", unit)
}

func (h *UnitHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "ID tidak valid")
		return
	}
	var req dto.UpdateUnitReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	unit, err := h.Svc.Update(uint(id), req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Unit berhasil diperbarui", unit)
}

func (h *UnitHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "ID tidak valid")
		return
	}
	err = h.Svc.Delete(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Unit berhasil dihapus", nil)
}