package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/hiboedi/zenshop/app/helpers"
	"github.com/hiboedi/zenshop/app/models"
	"gorm.io/gorm"
)

type AddressRepositoryImpl struct {
	UserRepo UserRepository
}

func NewAddressRepo(userRepo UserRepository) AdrressRepository {
	return &AddressRepositoryImpl{
		UserRepo: userRepo,
	}
}

func (r *AddressRepositoryImpl) Create(ctx context.Context, db *gorm.DB, address models.Address) (models.Address, error) {
	user, err := r.UserRepo.FindById(ctx, db, address.UserID)
	helpers.PanicIfError(err)

	modelAddress := models.Address{
		ID:         uuid.New().String(),
		UserID:     address.UserID,
		User:       user,
		Name:       address.Name,
		Address:    address.Address,
		CityID:     address.CityID,
		ProvinceID: address.ProvinceID,
		Phone:      address.Phone,
		PostCode:   address.PostCode,
		Email:      address.Email,
	}
	err = db.WithContext(ctx).Create(&modelAddress).Error
	helpers.PanicIfError(err)
	return modelAddress, nil
}

func (r *AddressRepositoryImpl) Update(ctx context.Context, db *gorm.DB, address models.Address) (models.Address, error) {

	userValid, err := r.UserRepo.FindById(ctx, db, address.UserID)

	if userValid.ID != "" {
		var modelAddress = &models.Address{
			Name:       address.Name,
			Address:    address.Address,
			CityID:     address.CityID,
			ProvinceID: address.ProvinceID,
			Phone:      address.Phone,
			PostCode:   address.PostCode,
			Email:      address.Email,
		}

		err = db.WithContext(ctx).Model(&models.Address{}).Where("id = ?", address.ID).Updates(modelAddress).Error
		helpers.PanicIfError(err)

		return address, nil
	}

	return models.Address{}, err
}

func (r *AddressRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, address models.Address) error {
	err := db.WithContext(ctx).Where("id = ?", address.ID).Delete(&address).Error
	helpers.PanicIfError(err)

	return nil
}

func (r *AddressRepositoryImpl) FindById(ctx context.Context, db *gorm.DB, addressId string) (models.Address, error) {
	var address models.Address
	err := db.WithContext(ctx).Preload("User").Where("id = ?", addressId).Take(&address).Error

	helpers.PanicIfError(err)
	return address, nil
}

func (r *AddressRepositoryImpl) FindByUserId(ctx context.Context, db *gorm.DB, userId string) ([]models.Address, error) {
	var addresses []models.Address

	err := db.WithContext(ctx).Model(&models.Address{}).Preload("User").Where("user_id = ?", userId).Find(&addresses).Error
	helpers.PanicIfError(err)

	return addresses, nil
}
