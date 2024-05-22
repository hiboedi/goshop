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

type ProductServiceImpl struct {
	ProductRepo repository.ProductRepository
	DB          *gorm.DB
	Validate    *validator.Validate
}

func NewProductService(productRepo repository.ProductRepository, db *gorm.DB, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepo: productRepo,
		DB:          db,
		Validate:    validate,
	}
}

func (s *ProductServiceImpl) Create(ctx context.Context, request models.ProductCreate) models.ProductResponse {
	err := s.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	product := models.Product{
		Name:        request.Name,
		Price:       request.Price,
		Stock:       request.Stock,
		Description: request.Description,
	}

	data, err := s.ProductRepo.Create(ctx, tx, product)
	helpers.PanicIfError(err)

	return models.ToProductReponse(data)
}

func (s *ProductServiceImpl) Update(ctx context.Context, request models.ProductUpdate) models.ProductResponse {
	err := s.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	product, err := s.ProductRepo.FindById(ctx, tx, request.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	product.Name = request.Name
	product.Price = request.Price
	product.Stock = request.Stock
	product.Description = request.Description

	data, err := s.ProductRepo.Update(ctx, tx, product)
	helpers.PanicIfError(err)

	return models.ToProductReponse(data)
}

func (s *ProductServiceImpl) Delete(ctx context.Context, prroductId string) {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	product, err := s.ProductRepo.FindById(ctx, tx, prroductId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	err = s.ProductRepo.Delete(ctx, tx, product)
	helpers.PanicIfError(err)
}

func (s *ProductServiceImpl) FindById(ctx context.Context, productId string) models.ProductResponse {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	product, err := s.ProductRepo.FindById(ctx, tx, productId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return models.ToProductReponse(product)
}

func (s *ProductServiceImpl) FindAll(ctx context.Context) []models.ProductResponse {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	products, err := s.ProductRepo.FindAll(ctx, tx)
	helpers.PanicIfError(err)
	return models.ToProductResponses(products)
}
