package models

type Cart struct {
	ID              string     `json:"id" gorm:"size:36;not null;uniqueIndex;primary_key"`
	CartItems       []CartItem `json:"cart_items"`
	BaseTotalPrice  float64    `json:"base_total_price" gorm:"not null"`
	TaxAmount       float64    `json:"tax_amount" gorm:"not null"`
	TaxPercent      float64    `json:"tax_percent" gorm:"not null"`
	DiscountAmount  float64    `json:"discount_amount" gorm:"not null"`
	DiscountPercent float64    `json:"discount_percent" gorm:"not null"`
	GrandTotal      float64    `json:"grand_total" gorm:"not null"`
	TotalWeight     int        `json:"total_weight" gorm:"-"`
}

type CartResponse struct {
	ID              string     `json:"id"`
	CartItems       []CartItem `json:"cart_items"`
	BaseTotalPrice  float64    `json:"base_total_price"`
	TaxAmount       float64    `json:"tax_amount"`
	TaxPercent      float64    `json:"tax_percent"`
	DiscountAmount  float64    `json:"discount_amount"`
	DiscountPercent float64    `json:"discount_percent"`
	GrandTotal      float64    `json:"grand_total"`
	TotalWeight     int        `json:"total_weight"`
}

type CartCreate struct {
	ID              string     `json:"id"`
	CartItems       []CartItem `json:"cart_items"`
	BaseTotalPrice  float64    `json:"base_total_price" validate:"required"`
	TaxAmount       float64    `json:"tax_amount" validate:"required"`
	TaxPercent      float64    `json:"tax_percent" validate:"required"`
	DiscountAmount  float64    `json:"discount_amount" validate:"required"`
	DiscountPercent float64    `json:"discount_percent" validate:"required"`
	GrandTotal      float64    `json:"grand_total" validate:"required"`
	TotalWeight     int        `json:"total_weight"`
}

type CartUpdate struct {
	ID              string     `json:"id"`
	CartItems       []CartItem `json:"cart_items"`
	BaseTotalPrice  float64    `json:"base_total_price" validate:"required"`
	TaxAmount       float64    `json:"tax_amount" validate:"required"`
	TaxPercent      float64    `json:"tax_percent" validate:"required"`
	DiscountAmount  float64    `json:"discount_amount" validate:"required"`
	DiscountPercent float64    `json:"discount_percent" validate:"required"`
	GrandTotal      float64    `json:"grand_total" validate:"required"`
	TotalWeight     int        `json:"total_weight"`
}

func ToCartReponse(cart Cart) CartResponse {
	return CartResponse{
		ID:              cart.ID,
		CartItems:       cart.CartItems,
		BaseTotalPrice:  cart.BaseTotalPrice,
		TaxAmount:       cart.TaxAmount,
		TaxPercent:      cart.TaxPercent,
		DiscountAmount:  cart.DiscountAmount,
		DiscountPercent: cart.DiscountPercent,
		GrandTotal:      cart.GrandTotal,
	}
}

func ToCartResponses(carts []Cart) []CartResponse {
	var responses []CartResponse

	for _, cart := range carts {
		responses = append(responses, ToCartReponse(cart))
	}
	return responses
}
