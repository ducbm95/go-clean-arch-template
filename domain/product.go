package domain

import (
	"context"

	"github.com/ducbm95/go-clean-arch-template/product/model"
)

type ProductService interface {
	Create(ctx context.Context, product model.Product) error
	GetByID(ctx context.Context, productID int64) model.Product
	Update(ctx context.Context, product model.Product) error
	Delete(ctx context.Context, productID int64) error
}
