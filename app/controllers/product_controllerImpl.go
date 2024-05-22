package controllers

import (
	"net/http"

	"github.com/hiboedi/zenshop/app/helpers"
	"github.com/hiboedi/zenshop/app/models"
	"github.com/hiboedi/zenshop/app/services"
	"github.com/hiboedi/zenshop/app/web"
	"github.com/julienschmidt/httprouter"
)

type ProductControllerImpl struct {
	ProductService services.ProductService
}

func NewProductController(productService services.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func (c *ProductControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	productCreateRequest := models.ProductCreate{}
	helpers.ToRequestBody(r, &productCreateRequest)

	producrResponse := c.ProductService.Create(r.Context(), productCreateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   producrResponse,
	}

	helpers.WriteResponseBody(w, webResponse)
}

func (c *ProductControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	productUpdateRequest := models.ProductUpdate{}
	helpers.ToRequestBody(r, &productUpdateRequest)

	productId := params.ByName("productId")

	productUpdateRequest.ID = productId

	producrResponse := c.ProductService.Update(r.Context(), productUpdateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   producrResponse,
	}

	helpers.WriteResponseBody(w, webResponse)
}

func (c *ProductControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	productId := params.ByName("productId")

	c.ProductService.Delete(r.Context(), productId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
	}
	helpers.WriteResponseBody(w, webResponse)
}

func (c *ProductControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	productId := params.ByName("productId")

	producrResponse := c.ProductService.FindById(r.Context(), productId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   producrResponse,
	}
	helpers.WriteResponseBody(w, webResponse)
}

func (c *ProductControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	productResponses := c.ProductService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   productResponses,
	}
	helpers.WriteResponseBody(w, webResponse)
}
