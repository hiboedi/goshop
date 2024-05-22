package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/hiboedi/zenshop/app/helpers"
	"github.com/hiboedi/zenshop/app/models"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

func NewUserRepo() UserRepository {
	return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) Create(ctx context.Context, db *gorm.DB, user models.User) (models.User, error) {
	hashedPassword, _ := helpers.MakePassword(user.Password)
	modelUser := models.User{
		ID:       uuid.New().String(),
		Name:     user.Name,
		Email:    user.Email,
		Password: hashedPassword,
	}
	err := db.WithContext(ctx).Create(&modelUser).Error
	helpers.PanicIfError(err)

	return modelUser, nil
}

func (r *UserRepositoryImpl) Update(ctx context.Context, db *gorm.DB, user models.User) (models.User, error) {
	var modelUser = &models.User{}
	hashPassword, err := helpers.MakePassword(user.Password)
	helpers.PanicIfError(err)
	err = db.WithContext(ctx).Model(modelUser).Where("id = ?", user.ID).Updates(map[string]interface{}{
		"name":     user.Name,
		"email":    user.Email,
		"password": hashPassword,
	}).Error
	helpers.PanicIfError(err)

	return user, nil
}

func (r *UserRepositoryImpl) FindById(ctx context.Context, db *gorm.DB, userId string) (models.User, error) {
	var user models.User
	err := db.WithContext(ctx).Where("id = ?", userId).Take(&user).Error
	helpers.PanicIfError(err)

	return user, nil
}

func (r *UserRepositoryImpl) FindByEmail(ctx context.Context, db *gorm.DB, email string) (models.User, error) {
	var user models.User
	err := db.WithContext(ctx).Where("email = ?", email).Take(&user).Error
	helpers.PanicIfError(err)

	return user, nil
}
