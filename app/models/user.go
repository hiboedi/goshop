package models

import "time"

type User struct {
	ID        string    `json:"id" gorm:"not null;uniqueIndex;primary_key"`
	Name      string    `json:"name" gorm:"not null;min:4;max:50"`
	Addresses []Address `json:"addresses" gorm:"foreignKey:UserID"`
	Email     string    `json:"username" gorm:"not null;unique"`
	Password  string    `json:"password" gorm:"not null;min:6;max:50"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoCreateTime;autoUpdateTime"`
}

type UserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserLoginResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type UserCreate struct {
	Name     string `validate:"required,min=4,max=50"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=6,max=50"`
}

type UserUpdate struct {
	ID       string `validate:"required"`
	Name     string `validate:"required,min=4,max=50"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=6,max=50"`
}

type UserLogin struct {
	ID       string
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

func ToUserReponse(user User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUserResponses(users []User) []UserResponse {
	var responses []UserResponse

	for _, user := range users {
		responses = append(responses, ToUserReponse(user))
	}
	return responses
}
