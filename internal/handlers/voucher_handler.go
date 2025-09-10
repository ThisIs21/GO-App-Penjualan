package handlers

import (
	"app-penjualan/internal/dto"
	"app-penjualan/internal/services"
	"app-penjualan/internal/utils"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type VoucherHandler struct {
	Svc *services.VoucherService
}

func NewVoucherHandler(svc *services.VoucherService) *VoucherHandler {
	return &VoucherHandler{Svc: svc}
}

func (h *VoucherHandler) Create(c *gin.Context) {
	var req dto.CreateVoucherReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	voucher, err := h.Svc.Create(req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusCreated, "Voucher berhasil ditambahkan", voucher)
}

func (h *VoucherHandler) List(c *gin.Context) {
	vouchers, err := h.Svc.List()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Daftar voucher berhasil diambil", vouchers)
}

func (h *VoucherHandler) Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "ID tidak valid")
		return
	}
	voucher, err := h.Svc.Get(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Voucher tidak ditemukan")
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Detail voucher berhasil diambil", voucher)
}

func (h *VoucherHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "ID tidak valid")
		return
	}
	var req dto.UpdateVoucherReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	voucher, err := h.Svc.Update(uint(id), req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Voucher berhasil diperbarui", voucher)
}

func (h *VoucherHandler) Delete(c *gin.Context) {
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
	utils.SuccessResponse(c, http.StatusOK, "Voucher berhasil dihapus", nil)
}