package models

import "time"

type Address struct {
	ID         string    `json:"id" gorm:"size:36;not null;uniqueIndex;primary_key"`
	UserID     string    `json:"user_id" gorm:"size:36;not null;index"`
	User       User      `json:"user" gorm:"foreignKey:UserID"`
	Name       string    `json:"name" gorm:"size:100"`
	Address    string    `json:"address" gorm:"size:255"`
	CityID     string    `json:"city_id" gorm:"size:100"`
	ProvinceID string    `json:"province_id" gorm:"size:100"`
	Phone      string    `json:"phone" gorm:"size:100"`
	Email      string    `json:"email" gorm:"size:100"`
	PostCode   string    `json:"post_code" gorm:"size:100"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type AddressResponse struct {
	ID         string `json:"id"`
	User       User
	UserID     string    `json:"user_id"`
	Address    string    `json:"address"`
	Name       string    `json:"name"`
	CityID     string    `json:"city_id"`
	ProvinceID string    `json:"province_id"`
	Phone      string    `json:"phone"`
	Email      string    `json:"email"`
	PostCode   string    `json:"post_code"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type AddressCreate struct {
	User       User
	UserID     string    `json:"user_id" validate:"required"`
	Name       string    `json:"name" validate:"required"`
	CityID     string    `json:"city_id" validate:"required"`
	ProvinceID string    `json:"province_id" validate:"required"`
	Address    string    `json:"address" validate:"required"`
	Phone      string    `json:"phone" validate:"required"`
	Email      string    `json:"email" validate:"required,email"`
	PostCode   string    `json:"post_code" validate:"required"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type AddressUpdate struct {
	ID         string    `json:"id" validate:"required"`
	UserID     string    `json:"user_id"`
	Name       string    `json:"name" validate:"required"`
	CityID     string    `json:"city_id" validate:"required"`
	ProvinceID string    `json:"province_id" validate:"required"`
	Address    string    `json:"address" validate:"required"`
	Phone      string    `json:"phone" validate:"required"`
	Email      string    `json:"email" validate:"required,email"`
	PostCode   string    `json:"post_code" validate:"required"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func ToAddressReponse(address Address) AddressResponse {
	return AddressResponse{
		ID:         address.ID,
		User:       address.User,
		UserID:     address.UserID,
		Name:       address.Name,
		Phone:      address.Phone,
		Email:      address.Email,
		Address:    address.Address,
		CityID:     address.CityID,
		ProvinceID: address.ProvinceID,
		PostCode:   address.PostCode,
		CreatedAt:  address.CreatedAt,
		UpdatedAt:  address.UpdatedAt,
	}
}

func ToAddressResponses(addresses []Address) []AddressResponse {
	var responses []AddressResponse

	for _, address := range addresses {
		responses = append(responses, ToAddressReponse(address))
	}
	return responses
}
