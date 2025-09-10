package services

import "gorm.io/gorm"

type ReportService struct{ DB *gorm.DB }
func NewReportService(db *gorm.DB)*ReportService{ return &ReportService{DB:db} }

// 1. Laporan Penjualan (per hari/bulan)
func (s *ReportService) SalesSummary(by string) ([]map[string]any, error){
	col := "DATE(tanggal)"; if by=="month" { col = "DATE_FORMAT(tanggal,'%Y-%m')" }
	q := s.DB.Table("sales").Select(col+" as period, COUNT(*) trx, SUM(total) total").Group(col).Order(col)
	var rows []map[string]any
	return rows, q.Find(&rows).Error
}
// 2. Laporan Retur Penjualan
func (s *ReportService) SaleReturnSummary(by string) ([]map[string]any, error){
	col := "DATE(tanggal)"; if by=="month" { col = "DATE_FORMAT(tanggal,'%Y-%m')" }
	q := s.DB.Table("sale_returns").Select(col+" as period, COUNT(*) trx").Group(col).Order(col)
	var rows []map[string]any; return rows, q.Find(&rows).Error
}
// 3. Laporan Pembelian
func (s *ReportService) PurchaseSummary(by string) ([]map[string]any, error){
	col := "DATE(tanggal)"; if by=="month" { col = "DATE_FORMAT(tanggal,'%Y-%m')" }
	q := s.DB.Table("purchases").Select(col+" as period, COUNT(*) trx, SUM(total) total").Group(col).Order(col)
	var rows []map[string]any; return rows, q.Find(&rows).Error
}
// 4. Laporan Retur Pembelian
func (s *ReportService) PurchaseReturnSummary(by string) ([]map[string]any, error){
	col := "DATE(tanggal)"; if by=="month" { col = "DATE_FORMAT(tanggal,'%Y-%m')" }
	q := s.DB.Table("purchase_returns").Select(col+" as period, COUNT(*) trx").Group(col).Order(col)
	var rows []map[string]any; return rows, q.Find(&rows).Error
}
// 5. Laporan Stok Opname
func (s *ReportService) StockOpnameSummary(by string) ([]map[string]any, error){
	col := "DATE(tanggal)"; if by=="month" { col = "DATE_FORMAT(tanggal,'%Y-%m')" }
	q := s.DB.Table("stock_opnames").Select(col+" as period, COUNT(*) opname").
		Group(col).Order(col)
	var rows []map[string]any; return rows, q.Find(&rows).Error
}
// 6. Laporan Stok (snapshot sederhana)
func (s *ReportService) StockSnapshot() ([]map[string]any, error){
	q := s.DB.Table("products").Select("id, nama_barang, stok, harga_beli, harga_jual")
	var rows []map[string]any; return rows, q.Find(&rows).Error
}
// 7. Laporan Persediaan (nilai persediaan)
func (s *ReportService) InventoryValue() ([]map[string]any, error){
	q := s.DB.Table("products").Select("SUM(stok*harga_beli) as nilai_beli, SUM(stok*harga_jual) as nilai_jual")
	var rows []map[string]any; return rows, q.Find(&rows).Error
}
