package services

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/hiboedi/zenshop/app/exception"
	"github.com/hiboedi/zenshop/app/helpers"
	"github.com/hiboedi/zenshop/app/middleware"
	"github.com/hiboedi/zenshop/app/models"
	"github.com/hiboedi/zenshop/app/repository"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepo repository.UserRepository
	DB       *gorm.DB
	Validate *validator.Validate
}

func NewUserService(userRepo repository.UserRepository, db *gorm.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepo: userRepo,
		DB:       db,
		Validate: validate,
	}
}

func (s *UserServiceImpl) Create(ctx context.Context, request models.UserCreate) models.UserResponse {
	err := s.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	user := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	data, err := s.UserRepo.Create(ctx, tx, user)
	helpers.PanicIfError(err)

	return models.ToUserReponse(data)
}

func (s *UserServiceImpl) Update(ctx context.Context, request models.UserUpdate) models.UserResponse {
	err := s.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	user, err := s.UserRepo.FindByEmail(ctx, tx, request.Email)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	resetPassword, _ := helpers.MakePassword(request.Password)
	user.Name = request.Name
	user.Password = resetPassword

	data, err := s.UserRepo.Update(ctx, tx, user)
	helpers.PanicIfError(err)

	return models.ToUserReponse(data)
}

func (s *UserServiceImpl) Delete(ctx context.Context, userId string) {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	user, err := s.UserRepo.FindById(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	err = s.UserRepo.Delete(ctx, tx, user)
	helpers.PanicIfError(err)
}

func (s *UserServiceImpl) Login(ctx context.Context, request models.UserLogin) (models.UserLoginResponse, bool) {
	err := s.Validate.Struct(request)
	helpers.PanicIfError(err)
	password := request.Password

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	user, err := s.UserRepo.FindByEmail(ctx, tx, request.Email)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	passwordSync := helpers.ComparePassword(password, user.Password)
	token, _ := middleware.CreateToken(user.ID)

	if !passwordSync {
		return models.UserLoginResponse{}, false
	} else {

		userLoginResponse := models.UserLoginResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Token: token,
		}
		return userLoginResponse, true
	}
}
