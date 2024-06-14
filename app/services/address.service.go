package services

import (
	"context"

	"github.com/hiboedi/zenshop/app/models"
)

type AddressService interface {
	Create(ctx context.Context, req models.AddressCreate) models.AddressResponse
	Update(ctx context.Context, req models.AddressUpdate) models.AddressResponse
	Delete(ctx context.Context, addressId string)
	FindById(ctx context.Context, addressId string) models.AddressResponse
	FindByUserId(ctx context.Context, userId string) []models.AddressResponse
}
