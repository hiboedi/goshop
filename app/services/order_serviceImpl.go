package services

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/hiboedi/zenshop/app/exception"
	"github.com/hiboedi/zenshop/app/helpers"
	"github.com/hiboedi/zenshop/app/models"
	"github.com/hiboedi/zenshop/app/repository"
	"gorm.io/gorm"
)

type OrderServiceImpl struct {
	OrderRepo repository.OrderRepository
	DB        *gorm.DB
	Validate  *validator.Validate
}

func NewOrderService(orderRepo repository.OrderRepository, db *gorm.DB, validate *validator.Validate) OrderService {
	return &OrderServiceImpl{
		OrderRepo: orderRepo,
		DB:        db,
		Validate:  validate,
	}
}

func (s *OrderServiceImpl) Create(ctx context.Context, request models.OrderCreate) models.OrderResponse {
	err := s.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	order := models.Order{
		UserID:     request.UserID,
		Note:       request.Note,
		BasePrice:  request.BasePrice,
		TotalPrice: request.TotalPrice,
	}

	data, err := s.OrderRepo.Create(ctx, tx, order)
	helpers.PanicIfError(err)

	return models.ToOrderReponse(data)
}

func (s *OrderServiceImpl) Update(ctx context.Context, request models.OrderUpdate) models.OrderResponse {
	err := s.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	order, err := s.OrderRepo.FindById(ctx, tx, request.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	order.Note = request.Note
	order.TotalPrice = request.TotalPrice
	order.BasePrice = request.BasePrice
	order.PaymentStatus = request.PaymentStatus
	if order.PaymentStatus != "Paid" && order.PaymentStatus != "Unpaid" {
		panic("Update failed")
	}

	data, err := s.OrderRepo.Update(ctx, tx, order)
	helpers.PanicIfError(err)

	return models.ToOrderReponse(data)
}

func (s *OrderServiceImpl) Delete(ctx context.Context, prroductId string) {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	order, err := s.OrderRepo.FindById(ctx, tx, prroductId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	err = s.OrderRepo.Delete(ctx, tx, order)
	helpers.PanicIfError(err)
}

func (s *OrderServiceImpl) FindAll(ctx context.Context) []models.OrderResponse {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	orders, err := s.OrderRepo.FindAll(ctx, tx)
	helpers.PanicIfError(err)
	return models.ToOrderReponses(orders)
}

func (s *OrderServiceImpl) FindByPaymentStatus(ctx context.Context, status string) []models.OrderResponse {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	orders, err := s.OrderRepo.FindByPaymentStatus(ctx, tx, status)
	helpers.PanicIfError(err)
	return models.ToOrderReponses(orders)
}

func (s *OrderServiceImpl) FindById(ctx context.Context, orderId string) models.OrderResponse {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	order, err := s.OrderRepo.FindById(ctx, tx, orderId)
	helpers.PanicIfError(err)
	return models.ToOrderReponse(order)
}
