	package handlers

	import (
		"strconv"

		"github.com/gin-gonic/gin"
		"app-penjualan/internal/models"
		"app-penjualan/internal/services"
		"app-penjualan/internal/utils"
	)

	type ProductHandler struct{ svc *services.ProductService }

	func NewProductHandler(s *services.ProductService) *ProductHandler {
		return &ProductHandler{svc: s}
	}

	func (h *ProductHandler) Register(rg *gin.RouterGroup) {
		rg.GET("/products", h.List)       // pakai yg sudah ada filter
		rg.GET("/products/:id", h.Get)
		rg.POST("/products", h.Create)
		rg.PUT("/products/:id", h.Update)
		rg.DELETE("/products/:id", h.Delete)
	}

	// ✅ hanya satu List, tapi sudah support filter & search
	func (h *ProductHandler) List(c *gin.Context) {
		search := c.Query("search")
		category := c.Query("category")
		minPrice := c.Query("min_price")
		maxPrice := c.Query("max_price")

		out, err := h.svc.ListFiltered(search, category, minPrice, maxPrice)
		if err != nil {
			utils.ServerError(c, err)
			return
		}
		utils.OK(c, out)
	}

	func (h *ProductHandler) Get(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		out, err := h.svc.Get(uint(id))
		if err != nil {
			utils.NotFound(c, "not found")
			return
		}
		utils.OK(c, out)
	}

	func (h *ProductHandler) Create(c *gin.Context) {
		var m models.Product
		if err := c.ShouldBindJSON(&m); err != nil {
			utils.BadRequest(c, err.Error())
			return
		}
		if err := h.svc.Create(&m); err != nil {
			utils.ServerError(c, err)
			return
		}
		utils.Created(c, m)
	}

	func (h *ProductHandler) Update(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var in models.Product
		if err := c.ShouldBindJSON(&in); err != nil {
			utils.BadRequest(c, err.Error())
			return
		}
		out, err := h.svc.Update(uint(id), &in)
		if err != nil {
			utils.ServerError(c, err)
			return
		}
		utils.OK(c, out)
	}

	func (h *ProductHandler) Delete(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		if err := h.svc.Delete(uint(id)); err != nil {
			utils.ServerError(c, err)
			return
		}
		utils.OK(c, gin.H{"deleted": id})
	}
