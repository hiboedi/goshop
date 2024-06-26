package services

import (
	"context"

	"github.com/hiboedi/zenshop/app/models"
)

type UserService interface {
	Create(ctx context.Context, request models.UserCreate) models.UserResponse
	Update(ctx context.Context, request models.UserUpdate) models.UserResponse
	Delete(ctx context.Context, userId string)
	Login(ctx context.Context, request models.UserLogin) (models.UserLoginResponse, bool)
}
