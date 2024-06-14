package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hiboedi/zenshop/app/helpers"
	"github.com/hiboedi/zenshop/app/models"
	"github.com/hiboedi/zenshop/app/services"
	"github.com/hiboedi/zenshop/app/web"
)

type ProductControllerImpl struct {
	ProductService services.ProductService
}

func NewProductController(productService services.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func (c *ProductControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	productCreateRequest := models.ProductCreate{}
	helpers.ToRequestBody(r, &productCreateRequest)

	productResponse := c.ProductService.Create(r.Context(), productCreateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   productResponse,
	}

	helpers.WriteResponseBody(w, webResponse)
}

func (c *ProductControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	productUpdateRequest := models.ProductUpdate{}
	helpers.ToRequestBody(r, &productUpdateRequest)

	vars := mux.Vars(r)
	productId := vars["productId"]
	productUpdateRequest.ID = productId

	productResponse := c.ProductService.Update(r.Context(), productUpdateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   productResponse,
	}

	helpers.WriteResponseBody(w, webResponse)
}

func (c *ProductControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["productId"]

	c.ProductService.Delete(r.Context(), productId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
	}
	helpers.WriteResponseBody(w, webResponse)
}

func (c *ProductControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["productId"]

	productResponse := c.ProductService.FindById(r.Context(), productId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   productResponse,
	}
	helpers.WriteResponseBody(w, webResponse)
}

func (c *ProductControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	productResponses := c.ProductService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   productResponses,
	}
	helpers.WriteResponseBody(w, webResponse)
}
