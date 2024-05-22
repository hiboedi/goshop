package repository

import (
	"context"
	"time"

	"github.com/hiboedi/zenshop/app/helpers"
	"github.com/hiboedi/zenshop/app/models"
	"gorm.io/gorm"
)

type OrderRepositoryImpl struct {
}

func NewOrderRepo() OrderRepository {
	return &OrderRepositoryImpl{}
}
func (r *OrderRepositoryImpl) Create(ctx context.Context, db *gorm.DB, order models.Order) (models.Order, error) {
	now := time.Now()
	modelOrder := models.Order{
		UserID:        order.UserID,
		Note:          order.Note,
		BasePrice:     order.BasePrice,
		TotalPrice:    order.TotalPrice,
		OrderDate:     now,
		PaymentStatus: "Unpaid",
		PaymentDue:    now.Add(3 * time.Hour),
	}

	err := db.WithContext(ctx).Create(&modelOrder).Error
	if err != nil {
		return models.Order{}, err
	}

	return modelOrder, nil
}

func (r *OrderRepositoryImpl) Update(ctx context.Context, db *gorm.DB, order models.Order) (models.Order, error) {
	var modelOrder = &models.Order{}
	err := db.WithContext(ctx).Model(modelOrder).Where("id = ?", order.ID).Updates(map[string]interface{}{
		"base_price":     order.BasePrice,
		"total_price":    order.TotalPrice,
		"note":           order.Note,
		"payment_status": order.PaymentStatus,
	}).Error
	helpers.PanicIfError(err)

	return order, nil
}

func (r *OrderRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, order models.Order) error {
	err := db.WithContext(ctx).Where("id = ?", order.ID).Delete(&order).Error
	helpers.PanicIfError(err)

	return nil
}

func (r *OrderRepositoryImpl) FindById(ctx context.Context, db *gorm.DB, orderId string) (models.Order, error) {
	var order models.Order
	err := db.WithContext(ctx).Where("id = ?", orderId).Take(&order).Error

	helpers.PanicIfError(err)
	return order, nil
}

func (r *OrderRepositoryImpl) FindByPaymentStatus(ctx context.Context, db *gorm.DB, status string) ([]models.Order, error) {
	var orders []models.Order
	err := db.WithContext(ctx).Where("payment_status = ?", status).Find(&orders).Error

	helpers.PanicIfError(err)
	return orders, nil
}

func (r *OrderRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB) ([]models.Order, error) {
	var orders []models.Order

	err := db.WithContext(ctx).Model(&models.Order{}).Find(&orders).Error
	helpers.PanicIfError(err)

	return orders, nil
}
