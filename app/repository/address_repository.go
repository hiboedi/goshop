package repository

import (
	"context"

	"github.com/hiboedi/zenshop/app/models"
	"gorm.io/gorm"
)

type AdrressRepository interface {
	Create(ctx context.Context, db *gorm.DB, address models.Address) (models.Address, error)
	Update(ctx context.Context, db *gorm.DB, address models.Address) (models.Address, error)
	Delete(ctx context.Context, db *gorm.DB, address models.Address) error
	FindById(ctx context.Context, db *gorm.DB, addressId string) (models.Address, error)
	FindByUserId(ctx context.Context, db *gorm.DB, userId string) ([]models.Address, error)
}
