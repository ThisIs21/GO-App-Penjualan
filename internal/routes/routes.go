
package routes

import (
	"app-penjualan/config"
	"app-penjualan/internal/handlers"
	"app-penjualan/internal/middlewares"
	"app-penjualan/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func Register(r *gin.Engine, db *gorm.DB, cfg *config.AppConfig) {
	// =======================
	// HEALTH CHECK
	// =======================
	r.GET("/api/health", func(c *gin.Context) {
		sqlDB, err := db.DB()
		if err != nil || sqlDB.Ping() != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "error",
				"db":     "down",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"db":      "connected",
			"message": "API is running 🚀",
		})
	})

	// =======================
	// AUTH (Public)
	// =======================
	authSvc := services.NewAuthService(db, cfg)
	authH := handlers.NewAuthHandler(authSvc)
	r.POST("/api/v1/auth/login", authH.Login)

	// =======================
	// Protected Group
	// =======================
	auth := r.Group("/api/v1")
	auth.Use(middlewares.JWTAuth(cfg))

	// =======================
	// OWNER (Full Access)
	// =======================
	owner := auth.Group("/owner")
	owner.Use(middlewares.RequireRoles("OWNER"))
	{
		// User Management
		owner.POST("/users", authH.CreateUser)

		// Master Data (Kategori, Unit, Supplier, Voucher)
		masterSvc := services.NewMasterService(db)
		masterH := handlers.NewMasterHandler(masterSvc)
		masterH.RegisterAll(owner)

		// Produk
		prodH := handlers.NewProductHandler(services.NewProductService(db))
		prodH.Register(owner)

		// Semua laporan
		reportH := handlers.NewReportHandler(services.NewReportService(db))
		owner.GET("/reports", reportH.GenerateReport)
	}

	// =======================
	// KASIR
	// =======================
	kasir := auth.Group("/kasir")
	kasir.Use(middlewares.RequireRoles("KASIR", "OWNER"))
	{
		saleH := handlers.NewSaleHandler(services.NewSaleService(db))
		saleH.Register(kasir)

		returnH := handlers.NewReturnHandler()
		kasir.POST("/sale-returns", returnH.CreateSaleReturn)

		// Laporan penjualan
		reportH := handlers.NewReportHandler(services.NewReportService(db))
		kasir.GET("/reports/sales", reportH.GenerateReport)
	}

	// =======================
	// PEMBELIAN
	// =======================
	pembelian := auth.Group("/pembelian")
	pembelian.Use(middlewares.RequireRoles("PEMBELIAN", "OWNER"))
	{
		purchaseH := handlers.NewPurchaseHandler(services.NewPurchaseService(db))
		purchaseH.Register(pembelian)

		returnH := handlers.NewReturnHandler()
		pembelian.POST("/purchase-returns", returnH.CreatePurchaseReturn)

		// Laporan pembelian & retur
		reportH := handlers.NewReportHandler(services.NewReportService(db))
		pembelian.GET("/reports/purchases", reportH.GenerateReport)
		pembelian.GET("/reports/purchase-returns", reportH.GenerateReport)
	}

	// =======================
	// GUDANG
	// =======================
	gudang := auth.Group("/gudang")
	gudang.Use(middlewares.RequireRoles("GUDANG", "OWNER"))
	{
		stockH := handlers.NewStockHandler(services.NewStockService(db))
		stockH.Register(gudang)

		// Laporan stok & history stok
		reportH := handlers.NewReportHandler(services.NewReportService(db))
		gudang.GET("/reports/stocks", reportH.GenerateReport)
		gudang.GET("/reports/stock-histories", reportH.GenerateReport)
	}

	// =======================
	// KEPALA GUDANG
	// =======================
	kepalaGudang := auth.Group("/kepala-gudang")
	kepalaGudang.Use(middlewares.RequireRoles("KEPALA_GUDANG", "OWNER"))
	{
		purchaseH := handlers.NewPurchaseHandler(services.NewPurchaseService(db))
		kepalaGudang.POST("/purchases/:id/approve", purchaseH.Approve)
		kepalaGudang.POST("/purchases/:id/reject", purchaseH.Reject)

		returnH := handlers.NewReturnHandler()
		kepalaGudang.POST("/sale-returns/:id/approve", returnH.ApproveSaleReturn)
		kepalaGudang.POST("/purchase-returns/:id/approve", returnH.ApprovePurchaseReturn)

		// Laporan stok & history stok
		reportH := handlers.NewReportHandler(services.NewReportService(db))
		kepalaGudang.GET("/reports/stocks", reportH.GenerateReport)
		kepalaGudang.GET("/reports/stock-histories", reportH.GenerateReport)
	}

	// =======================
	// Produk (List Bisa diakses Semua Role)
	// =======================
	auth.GET("/products",
		middlewares.RequireRoles("OWNER", "GUDANG", "KASIR", "PEMBELIAN", "KEPALA_GUDANG"),
		handlers.NewProductHandler(services.NewProductService(db)).List)

	// =======================
	// HISTORY
	// =======================
	histH := handlers.NewHistoryHandler(services.NewHistoryService(db))
	hist := auth.Group("/history")
	hist.Use(middlewares.RequireRoles("OWNER", "PEMBELIAN", "KASIR", "GUDANG", "KEPALA_GUDANG"))
	{
		hist.GET("/purchases", histH.Purchases)
		hist.GET("/sales", histH.Sales)
		hist.GET("/purchase-returns", histH.PurchaseReturns)
		hist.GET("/sale-returns", histH.SaleReturns)
		hist.GET("/stock-opnames", histH.StockOpnames)
	}

	// =======================
	// CORRECTIONS
	// =======================
	corrH := handlers.NewCorrectionHandler(services.NewCorrectionService(db))
	corr := auth.Group("/corrections")
	corr.Use(middlewares.RequireRoles("OWNER", "PEMBELIAN", "KASIR", "GUDANG", "KEPALA_GUDANG"))
	{
		corr.POST("", corrH.Create)
	}
}
