package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID          string    `json:"id" gorm:"type:uuid;primary_key;not null;uniqueIndex"`
	UserId      string    `json:"user_id" gorm:"index"`
	Name        string    `json:"name" gorm:"unique;varchar(100);not null"`
	Price       float64   `json:"price" gorm:"not null"`
	Description string    `json:"description" gorm:"not null;type:text"`
	Stock       int32     `json:"stock" gorm:"not null;default:0"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoCreateTime;autoUpdateTime"`
}

type ProductResponse struct {
	ID          string    `json:"id"`
	UserId      string    `json:"user_id"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Description string    `json:"description"`
	Stock       int32     `json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ProductCreate struct {
	UserId      string
	Name        string  `validate:"required,max=200,min=4"`
	Price       float64 `validate:"required"`
	Description string  `validate:"required"`
	Stock       int32   `validate:"required"`
}

type ProductUpdate struct {
	ID          string `validate:"required"`
	UserId      string
	Name        string  `validate:"required,max=200,min=4"`
	Price       float64 `validate:"required"`
	Description string  `validate:"required"`
	Stock       int32   `validate:"required"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func ToProductReponse(product Product) ProductResponse {
	return ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Stock:       product.Stock,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}

func ToProductResponses(products []Product) []ProductResponse {
	var responses []ProductResponse

	for _, product := range products {
		responses = append(responses, ToProductReponse(product))
	}
	return responses
}

func (product *Product) BeforeCreate(tx *gorm.DB) (err error) {
	UUID := uuid.New()
	product.ID = UUID.String()
	return
}
