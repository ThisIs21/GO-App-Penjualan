package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"app-penjualan/internal/models"
	"app-penjualan/internal/repositories"
	"app-penjualan/internal/services"
	"app-penjualan/internal/utils"
)

type MasterHandler struct{
	db *gorm.DB
}
func NewMasterHandler(s *services.MasterService) *MasterHandler { return &MasterHandler{db: s.DB()} }

// registerCRUD adalah fungsi biasa (bukan method) yang menggunakan type parameters
func registerCRUD[T any](db *gorm.DB, rg *gin.RouterGroup, path string) {
	crud := repositories.NewCRUD[T](db)
	rg.GET(path, func(c *gin.Context) {
		var list []T; if err := crud.List(&list); err != nil { utils.ServerError(c, err); return }
		utils.OK(c, list)
	})
	rg.GET(path+"/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id")); var m T
		if err := crud.Get(uint(id), &m); err != nil { utils.NotFound(c, "not found"); return }
		utils.OK(c, m)
	})
	rg.POST(path, func(c *gin.Context) {
		var m T; if err := c.ShouldBindJSON(&m); err != nil { utils.BadRequest(c, err.Error()); return }
		if err := crud.Create(&m); err != nil { utils.ServerError(c, err); return }
		utils.Created(c, m)
	})
	rg.PUT(path+"/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id")); var m T
		if err := c.ShouldBindJSON(&m); err != nil { utils.BadRequest(c, err.Error()); return }
		if err := crud.Update(uint(id), &m); err != nil { utils.ServerError(c, err); return }
		utils.OK(c, m)
	})
	rg.DELETE(path+"/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id")); var m T
		if err := crud.Delete(uint(id), &m); err != nil { utils.ServerError(c, err); return }
		utils.OK(c, gin.H{"deleted": id})
	})
}

// RegisterAll adalah method yang memanggil fungsi registerCRUD
func (h *MasterHandler) RegisterAll(rg *gin.RouterGroup) {
	registerCRUD[models.Category](h.db, rg, "/categories")
	registerCRUD[models.Unit](h.db, rg, "/units")
	registerCRUD[models.Supplier](h.db, rg, "/suppliers")
	registerCRUD[models.Customer](h.db, rg, "/customers")
	registerCRUD[models.Voucher](h.db, rg, "/vouchers")
}