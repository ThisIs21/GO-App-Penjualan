package handlers

import (
	"app-penjualan/internal/dto"
	"app-penjualan/internal/services"
	"app-penjualan/internal/utils"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type SupplierHandler struct {
	Svc *services.SupplierService
}

func NewSupplierHandler(svc *services.SupplierService) *SupplierHandler {
	return &SupplierHandler{Svc: svc}
}

func (h *SupplierHandler) Create(c *gin.Context) {
	var req dto.CreateSupplierReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	supplier, err := h.Svc.Create(req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusCreated, "Supplier berhasil ditambahkan", supplier)
}

func (h *SupplierHandler) List(c *gin.Context) {
	suppliers, err := h.Svc.List()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Daftar supplier berhasil diambil", suppliers)
}

func (h *SupplierHandler) Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "ID tidak valid")
		return
	}
	supplier, err := h.Svc.Get(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Supplier tidak ditemukan")
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Detail supplier berhasil diambil", supplier)
}

func (h *SupplierHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "ID tidak valid")
		return
	}
	var req dto.UpdateSupplierReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	supplier, err := h.Svc.Update(uint(id), req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Supplier berhasil diperbarui", supplier)
}

func (h *SupplierHandler) Delete(c *gin.Context) {
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
	utils.SuccessResponse(c, http.StatusOK, "Supplier berhasil dihapus", nil)
}