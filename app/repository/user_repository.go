package repository

import (
	"context"

	"github.com/hiboedi/zenshop/app/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, db *gorm.DB, user models.User) (models.User, error)
	Update(ctx context.Context, db *gorm.DB, user models.User) (models.User, error)
	FindById(ctx context.Context, db *gorm.DB, userId string) (models.User, error)
	Delete(ctx context.Context, db *gorm.DB, user models.User) error
	FindByEmail(ctx context.Context, db *gorm.DB, email string) (models.User, error)
}
