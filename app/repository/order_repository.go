package repository

import (
	"context"

	"github.com/hiboedi/zenshop/app/models"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(ctx context.Context, db *gorm.DB, order models.Order) (models.Order, error)
	Update(ctx context.Context, db *gorm.DB, order models.Order) (models.Order, error)
	Delete(ctx context.Context, db *gorm.DB, order models.Order) error
	FindById(ctx context.Context, db *gorm.DB, OrderId string) (models.Order, error)
	FindByPaymentStatus(ctx context.Context, db *gorm.DB, status string) ([]models.Order, error)
	FindAll(ctx context.Context, db *gorm.DB) ([]models.Order, error)
}
