package services

import (
	"context"

	"github.com/hiboedi/zenshop/app/models"
)

type ProductService interface {
	Create(ctx context.Context, req models.ProductCreate) models.ProductResponse
	Update(ctx context.Context, req models.ProductUpdate) models.ProductResponse
	Delete(ctx context.Context, productId string)
	FindById(ctx context.Context, productId string) models.ProductResponse
	FindAll(ctx context.Context) []models.ProductResponse
}
