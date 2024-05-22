package services

import (
	"context"

	"github.com/hiboedi/zenshop/app/models"
)

type OrderService interface {
	Create(ctx context.Context, request models.OrderCreate) models.OrderResponse
	Update(ctx context.Context, request models.OrderUpdate) models.OrderResponse
	Delete(ctx context.Context, orderId string)
	FindById(ctx context.Context, orderId string) models.OrderResponse
	FindAll(ctx context.Context) []models.OrderResponse
	FindByPaymentStatus(ctx context.Context, status string) []models.OrderResponse
}
