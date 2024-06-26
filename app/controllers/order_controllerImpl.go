package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hiboedi/zenshop/app/helpers"
	"github.com/hiboedi/zenshop/app/models"
	"github.com/hiboedi/zenshop/app/services"
	"github.com/hiboedi/zenshop/app/web"
)

type OrderControllerImpl struct {
	OrderService services.OrderService
}

func NewOrderController(orderService services.OrderService) OrderController {
	return &OrderControllerImpl{
		OrderService: orderService,
	}
}

func (c *OrderControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	orderCreateRequest := models.OrderCreate{}
	helpers.ToRequestBody(r, &orderCreateRequest)
	cookie, err := helpers.GetCookie(w, r)
	if err != nil {
		http.Redirect(w, r, "/api/users/login", http.StatusPermanentRedirect)
		return
	}
	userID := cookie.Value
	orderCreateRequest.UserID = userID

	orderResponse := c.OrderService.Create(r.Context(), orderCreateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   orderResponse,
	}

	helpers.WriteResponseBody(w, webResponse)
}

func (c *OrderControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	orderUpdateRequest := models.OrderUpdate{}
	helpers.ToRequestBody(r, &orderUpdateRequest)

	vars := mux.Vars(r)
	orderId := vars["orderId"]
	orderUpdateRequest.ID = orderId

	orderResponse := c.OrderService.Update(r.Context(), orderUpdateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   orderResponse,
	}

	helpers.WriteResponseBody(w, webResponse)
}

func (c *OrderControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderId := vars["orderId"]

	c.OrderService.Delete(r.Context(), orderId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
	}
	helpers.WriteResponseBody(w, webResponse)
}

func (c *OrderControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderId := vars["orderId"]

	orderResponse := c.OrderService.FindById(r.Context(), orderId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   orderResponse,
	}
	helpers.WriteResponseBody(w, webResponse)
}

func (c *OrderControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	orderResponses := c.OrderService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   orderResponses,
	}
	helpers.WriteResponseBody(w, webResponse)
}

func (c *OrderControllerImpl) FindAllOrByPaymentStatus(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")

	if status != "" {
		if status != "Paid" && status != "Unpaid" {
			webResponse := web.WebResponse{
				Code:   http.StatusBadRequest,
				Status: "Bad Request",
				Data:   "Invalid status value",
			}
			helpers.WriteResponseBody(w, webResponse)
			return
		}

		orderResponses := c.OrderService.FindByPaymentStatus(r.Context(), status)
		if orderResponses == nil {
			webResponse := web.WebResponse{
				Code:   http.StatusInternalServerError,
				Status: "Internal Server Error",
				Data:   "Error retrieving orders",
			}
			helpers.WriteResponseBody(w, webResponse)
			return
		}

		webResponse := web.WebResponse{
			Code:   http.StatusOK,
			Status: "Ok",
			Data:   orderResponses,
		}
		helpers.WriteResponseBody(w, webResponse)
	} else {
		orderResponses := c.OrderService.FindAll(r.Context())
		webResponse := web.WebResponse{
			Code:   http.StatusOK,
			Status: "Ok",
			Data:   orderResponses,
		}
		helpers.WriteResponseBody(w, webResponse)
	}
}
