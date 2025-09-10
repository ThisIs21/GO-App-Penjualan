package services

import (
	"errors"
	

	"gorm.io/gorm"
	"app-penjualan/internal/models"
	"app-penjualan/internal/repositories"
)

type ProductService struct{ repo *repositories.ProductRepo }

func NewProductService(db *gorm.DB) *ProductService {
	return &ProductService{repositories.NewProductRepo(db)}
}

// ambil semua produk (tanpa filter)
func (s *ProductService) List() ([]models.Product, error) {
	return s.repo.FindAll()
}

// ambil produk by ID
func (s *ProductService) Get(id uint) (*models.Product, error) {
	return s.repo.FindByID(id)
}

// create produk
func (s *ProductService) Create(p *models.Product) error {
	return s.repo.Create(p)
}

// update produk
func (s *ProductService) Update(id uint, in *models.Product) (*models.Product, error) {
	p, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	p.Name = in.Name
	p.CategoryID = in.CategoryID
	p.UnitID = in.UnitID
	p.CostPrice = in.CostPrice
	p.SellPrice = in.SellPrice
	p.SupplierID = in.SupplierID
	if err := s.repo.Update(p); err != nil {
		return nil, err
	}
	return p, nil
}

// delete produk
func (s *ProductService) Delete(id uint) error {
	return s.repo.Delete(id)
}

// sesuaikan stok (tambah/kurang)
func (s *ProductService) AdjustStock(productID uint, delta int) error {
	if delta == 0 {
		return nil
	}
	if delta > 0 {
		return s.repo.AddStock(productID, delta)
	}
	// minus
	if err := s.repo.ReduceStock(productID, -delta); err != nil {
		return errors.New("insufficient stock or product not found")
	}
	return nil
}

// ✅ fitur baru: search & filter
func (s *ProductService) ListFiltered(search, category, minPrice, maxPrice string) ([]models.Product, error) {
	return s.repo.FindFiltered(search, category, minPrice, maxPrice)}