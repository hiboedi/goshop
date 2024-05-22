package repository

import (
	"context"

	"github.com/hiboedi/zenshop/app/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(ctx context.Context, db *gorm.DB, product models.Product) (models.Product, error)
	Update(ctx context.Context, db *gorm.DB, product models.Product) (models.Product, error)
	Delete(ctx context.Context, db *gorm.DB, product models.Product) error
	FindById(ctx context.Context, db *gorm.DB, productId string) (models.Product, error)
	FindAll(ctx context.Context, db *gorm.DB) ([]models.Product, error)
}
