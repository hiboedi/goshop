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

type AddressServiceImpl struct {
	AddressRepo repository.AdrressRepository
	DB          *gorm.DB
	Validate    *validator.Validate
}

func NewAddressService(addressRepo repository.AdrressRepository, db *gorm.DB, validate *validator.Validate) AddressService {
	return &AddressServiceImpl{
		AddressRepo: addressRepo,
		DB:          db,
		Validate:    validate,
	}
}

func (s *AddressServiceImpl) Create(ctx context.Context, request models.AddressCreate) models.AddressResponse {
	err := s.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	address := models.Address{
		UserID:     request.UserID,
		Name:       request.Name,
		Address:    request.Address,
		CityID:     request.CityID,
		ProvinceID: request.ProvinceID,
		Phone:      request.Phone,
		PostCode:   request.PostCode,
		Email:      request.Email,
	}

	data, err := s.AddressRepo.Create(ctx, tx, address)
	helpers.PanicIfError(err)

	return models.ToAddressReponse(data)
}

func (s *AddressServiceImpl) Update(ctx context.Context, request models.AddressUpdate) models.AddressResponse {
	err := s.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	address, err := s.AddressRepo.FindById(ctx, tx, request.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	address.UserID = request.UserID
	address.Name = request.Name
	address.Address = request.Address
	address.CityID = request.CityID
	address.Email = request.Email
	address.Phone = request.Phone
	address.ProvinceID = request.ProvinceID
	address.PostCode = request.PostCode

	data, err := s.AddressRepo.Update(ctx, tx, address)
	helpers.PanicIfError(err)

	return models.ToAddressReponse(data)
}

func (s *AddressServiceImpl) Delete(ctx context.Context, addressId string) {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	address, err := s.AddressRepo.FindById(ctx, tx, addressId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	err = s.AddressRepo.Delete(ctx, tx, address)
	helpers.PanicIfError(err)
}

func (s *AddressServiceImpl) FindById(ctx context.Context, addressId string) models.AddressResponse {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	address, err := s.AddressRepo.FindById(ctx, tx, addressId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return models.ToAddressReponse(address)
}

func (s *AddressServiceImpl) FindByUserId(ctx context.Context, userId string) []models.AddressResponse {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	addresses, err := s.AddressRepo.FindByUserId(ctx, tx, userId)
	helpers.PanicIfError(err)
	return models.ToAddressResponses(addresses)
}
