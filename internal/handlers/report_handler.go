package handlers

import (
	"bytes"
	"encoding/csv"
	"net/http"
	"strings" // Digunakan untuk memodifikasi string judul laporan
	"fmt" // Digunakan untuk formatting string
	
	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
	"app-penjualan/internal/services"
)

type ReportHandler struct {
	Svc *services.ReportService
}

// NewReportHandler membuat dan mengembalikan instance baru dari ReportHandler
func NewReportHandler(s *services.ReportService) *ReportHandler {
	return &ReportHandler{Svc: s}
}

// helper export CSV
// Fungsi ini tetap sama karena sudah baik
func writeCSV(c *gin.Context, header []string, rows [][]string, filename string) {
	buf := &bytes.Buffer{}
	w := csv.NewWriter(buf)
	_ = w.Write(header)
	_ = w.WriteAll(rows)
	w.Flush()
	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", "attachment; filename=" + filename)
	c.String(http.StatusOK, buf.String())
}

// helper export PDF (sederhana satu tabel)
// Fungsi ini juga tetap sama
func writePDF(c *gin.Context, title string, header []string, rows [][]string, filename string) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(40, 10, title)
	pdf.Ln(12)
	pdf.SetFont("Arial", "", 10)
	for _, h := range header {
		pdf.CellFormat(40, 8, h, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(-1)
	for _, r := range rows {
		for _, col := range r {
			pdf.CellFormat(40, 8, col, "1", 0, "L", false, 0, "")
		}
		pdf.Ln(-1)
	}
	var buf bytes.Buffer
	_ = pdf.Output(&buf)
	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename=" + filename)
	c.Data(http.StatusOK, "application/pdf", buf.Bytes())
}

// toS adalah helper kecil untuk konversi tipe data ke string
// Fungsi ini tetap sama
func toS(v any) string {
	if v == nil {
		return ""
	}
	return fmt.Sprintf("%v", v)
}

// GenerateReport adalah handler umum yang dapat menangani berbagai jenis laporan
func (h *ReportHandler) GenerateReport(c *gin.Context) {
	// Ambil parameter "type" dari URL, misalnya: /api/report?type=sales
	reportType := c.DefaultQuery("type", "")
	if reportType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "report type is required"})
		return
	}

	by := c.DefaultQuery("by", "day")
	var data []map[string]any
	var err error

	// Memanggil service yang sesuai berdasarkan jenis laporan
	switch reportType {
	case "sales":
		data, err = h.Svc.SalesSummary(by)
	case "purchases":
		data, err = h.Svc.PurchaseSummary(by)
	case "sale-returns":
		data, err = h.Svc.SaleReturnSummary(by)
	case "purchase-returns":
		data, err = h.Svc.PurchaseReturnSummary(by)
	case "stock-opnames":
		data, err = h.Svc.StockOpnameSummary(by)
	case "stock-snapshot":
		data, err = h.Svc.StockSnapshot()
	case "inventory-value":
		data, err = h.Svc.InventoryValue()
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid report type"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Memeriksa apakah data perlu diekspor
	if c.Query("export") == "csv" {
		header := []string{"period", "trx", "total"}
		var rows [][]string
		for _, m := range data {
			rows = append(rows, []string{toS(m["period"]), toS(m["trx"]), toS(m["total"])})
		}
		filename := strings.ReplaceAll(reportType, "-", "_") + ".csv"
		writeCSV(c, header, rows, filename)
		return
	}

	if c.Query("export") == "pdf" {
		header := []string{"period", "trx", "total"}
		var rows [][]string
		for _, m := range data {
			rows = append(rows, []string{toS(m["period"]), toS(m["trx"]), toS(m["total"])})
		}
		// Membuat judul PDF yang lebih mudah dibaca
		title := strings.ReplaceAll(reportType, "-", " ")
		title = strings.Title(title) + " Report"
		filename := strings.ReplaceAll(reportType, "-", "_") + ".pdf"
		writePDF(c, title, header, rows, filename)
		return
	}

	// Mengirim data sebagai JSON jika tidak ada permintaan ekspor
	c.JSON(http.StatusOK, gin.H{"data": data})
}
  