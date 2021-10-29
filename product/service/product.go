package service

import (
	"context"

	"github.com/ducbm95/go-clean-arch-template/domain"
	"github.com/ducbm95/go-clean-arch-template/product/model"
	"gorm.io/gorm"
)

type productService struct {
	db *gorm.DB
}

func NewProductService(db *gorm.DB) domain.ProductService {
	return &productService{
		db: db,
	}
}

func (ps *productService) Create(ctx context.Context, product model.Product) error {
	tx := ps.db.Begin()
	defer tx.Commit()
	return nil
}

func (ps *productService) GetByID(ctx context.Context, productID int64) (*model.Product, error) {
	return nil, nil
}

func (ps *productService) Update(ctx context.Context, product model.Product) error {
	return nil
}

func (ps *productService) Delete(ctx context.Context, productID int64) error {
	return nil
}
