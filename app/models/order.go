package models

import (
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	ID            string    `json:"id" gorm:"unique;primary_key;not null"`
	UserID        string    `json:"user_id"`
	Code          string    `json:"code" gorm:"unique"`
	Note          string    `json:"note"`
	BasePrice     float64   `json:"base_price" gorm:"not null"`
	TotalPrice    float64   `json:"total_price" gorm:"not null"`
	OrderDate     time.Time `json:"order_date" gorm:"autoCreateTime"`
	PaymentStatus string    `json:"payment_status" gorm:"default:unpay"`
	PaymentDue    time.Time `json:"payment_due"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoCreateTime;autoUpdateTime"`
}

type OrderResponse struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	Code          string    `json:"code"`
	Note          string    `json:"note"`
	BasePrice     float64   `json:"base_price"`
	TotalPrice    float64   `json:"total_price"`
	OrderDate     time.Time `json:"order_date"`
	PaymentStatus string    `json:"payment_status"`
	PaymentDue    time.Time `json:"payment_due"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type OrderCreate struct {
	ID            string
	UserID        string `validate:"required"`
	Code          string
	Note          string  `validate:"required"`
	BasePrice     float64 `validate:"required" json:"base_price"`
	TotalPrice    float64 `validate:"required" json:"total_price"`
	OrderDate     time.Time
	PaymentStatus string
	PaymentDue    time.Time
}

type OrderUpdate struct {
	ID            string `validate:"required"`
	Note          string
	BasePrice     float64 `validate:"required" json:"base_price"`
	TotalPrice    float64 `validate:"required" json:"total_price"`
	PaymentStatus string  `json:"payment_status"`
	UpdatedAt     time.Time
}

func (order *Order) BeforeCreate(db *gorm.DB) error {
	if order.ID == "" {
		order.ID = uuid.New().String()
	}

	order.Code = generateOrderNumber(db)

	return nil
}

func ToOrderReponse(order Order) OrderResponse {
	return OrderResponse{
		ID:            order.ID,
		UserID:        order.UserID,
		Code:          order.Code,
		Note:          order.Note,
		BasePrice:     order.BasePrice,
		TotalPrice:    order.TotalPrice,
		PaymentDue:    order.PaymentDue,
		PaymentStatus: order.PaymentStatus,
		CreatedAt:     order.CreatedAt,
		UpdatedAt:     order.UpdatedAt,
	}
}

func ToOrderReponses(orders []Order) []OrderResponse {
	var responses []OrderResponse

	for _, order := range orders {
		responses = append(responses, ToOrderReponse(order))
	}
	return responses
}

func generateOrderNumber(db *gorm.DB) string {
	now := time.Now()
	month := now.Month()
	year := strconv.Itoa(now.Year())

	dateCode := "/ORDER/" + intToRoman(int(month)) + "/" + year

	var latestOrder Order

	err := db.Debug().Order("created_at DESC").Find(&latestOrder).Error

	latestNumber, _ := strconv.Atoi(strings.Split(latestOrder.Code, "/")[0])
	if err != nil {
		latestNumber = 1
	}

	number := latestNumber + 1

	invoiceNumber := strconv.Itoa(number) + dateCode

	return invoiceNumber
}

func intToRoman(num int) string {
	values := []int{
		1000, 900, 500, 400,
		100, 90, 50, 40,
		10, 9, 5, 4, 1,
	}

	symbols := []string{
		"M", "CM", "D", "CD",
		"C", "XC", "L", "XL",
		"X", "IX", "V", "IV",
		"I"}
	roman := ""
	i := 0

	for num > 0 {
		// calculate the number of times this num is completly divisible by values[i]
		// times will only be > 0, when num >= values[i]
		k := num / values[i]
		for j := 0; j < k; j++ {
			// buildup roman numeral
			roman += symbols[i]

			// reduce the value of num.
			num -= values[i]
		}
		i++
	}
	return roman
}
