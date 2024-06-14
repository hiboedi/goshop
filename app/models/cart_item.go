package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartItem struct {
	ID              string    `json:"id" gorm:"size:36;not null;uniqueIndex;primary_key"`
	Cart            Cart      `json:"cart" gorm:"foreignKey:CartID"`
	CartID          string    `json:"cart_id" gorm:"size:36;index"`
	Product         Product   `json:"product" gorm:"foreignKey:ProductID"`
	ProductID       string    `json:"product_id" gorm:"size:36;index"`
	Qty             int       `json:"qty"`
	BasePrice       float64   `json:"base_price" gorm:"not null"`
	BaseTotal       float64   `json:"total_price" gorm:"not null"`
	TaxAmount       float64   `json:"tax_amount" gorm:"not null"`
	TaxPercent      float64   `json:"tax_percent" gorm:"not null"`
	DiscountAmount  float64   `json:"discount_amount" gorm:"not null"`
	DiscountPercent float64   `json:"discount_percent" gorm:"not null"`
	SubTotal        float64   `json:"subtotal" gorm:"not null"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"autoCreateTime;autoUpdateTime"`
}

func (c *CartItem) BeforeCreate(tx *gorm.DB) error {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}

	return nil
}
