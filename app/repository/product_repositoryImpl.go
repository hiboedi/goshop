package repository

import (
	"context"

	"github.com/hiboedi/zenshop/app/helpers"
	"github.com/hiboedi/zenshop/app/models"
	"gorm.io/gorm"
)

type ProductRepoImpl struct {
}

func NewProductRepo() ProductRepository {
	return &ProductRepoImpl{}
}

func (r *ProductRepoImpl) Create(ctx context.Context, db *gorm.DB, product models.Product) (models.Product, error) {
	modelProduct := models.Product{
		Name:        product.Name,
		Price:       product.Price,
		Stock:       product.Stock,
		Description: product.Description,
	}
	err := db.WithContext(ctx).Create(&modelProduct).Error
	helpers.PanicIfError(err)

	return modelProduct, nil
}

func (r *ProductRepoImpl) Update(ctx context.Context, db *gorm.DB, product models.Product) (models.Product, error) {
	var modelProduct = &models.Product{}
	err := db.WithContext(ctx).Model(modelProduct).Where("id = ?", product.ID).Updates(map[string]interface{}{
		"name":        product.Name,
		"price":       product.Price,
		"stock":       product.Stock,
		"description": product.Description,
	}).Error
	helpers.PanicIfError(err)

	return product, nil
}

func (r *ProductRepoImpl) Delete(ctx context.Context, db *gorm.DB, product models.Product) error {
	err := db.WithContext(ctx).Where("id = ?", product.ID).Delete(&product).Error
	helpers.PanicIfError(err)

	return nil
}

func (r *ProductRepoImpl) FindById(ctx context.Context, db *gorm.DB, productId string) (models.Product, error) {
	var product models.Product
	err := db.WithContext(ctx).Where("id = ?", productId).Take(&product).Error

	helpers.PanicIfError(err)
	return product, nil
}

func (r *ProductRepoImpl) FindAll(ctx context.Context, db *gorm.DB) ([]models.Product, error) {
	var products []models.Product

	err := db.WithContext(ctx).Model(&models.Product{}).Find(&products).Error
	helpers.PanicIfError(err)

	return products, nil
}
