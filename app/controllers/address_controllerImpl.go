package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hiboedi/zenshop/app/helpers"
	"github.com/hiboedi/zenshop/app/models"
	"github.com/hiboedi/zenshop/app/services"
	"github.com/hiboedi/zenshop/app/web"
)

type AddressControllerImpl struct {
	AddressService services.AddressService
}

func NewAddressController(addressService services.AddressService) AddressController {
	return &AddressControllerImpl{
		AddressService: addressService,
	}
}

func (c *AddressControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	addressCreateRequest := models.AddressCreate{}

	cookie, err := helpers.GetCookie(w, r)
	if err != nil {
		http.Redirect(w, r, "/api/users/login", http.StatusPermanentRedirect)
		return
	}

	userID := cookie.Value
	addressCreateRequest.UserID = userID

	helpers.ToRequestBody(r, &addressCreateRequest)

	addressResponse := c.AddressService.Create(r.Context(), addressCreateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   addressResponse,
	}

	helpers.WriteResponseBody(w, webResponse)
}

func (c *AddressControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	addressUpdateRequest := models.AddressUpdate{}
	helpers.ToRequestBody(r, &addressUpdateRequest)

	vars := mux.Vars(r)
	addressId := vars["addressId"]
	addressUpdateRequest.ID = addressId

	cookie, err := helpers.GetCookie(w, r)
	if err != nil {
		http.Redirect(w, r, "/api/users/login", http.StatusPermanentRedirect)
		return
	}

	userID := cookie.Value
	addressUpdateRequest.UserID = userID

	addressResponse := c.AddressService.Update(r.Context(), addressUpdateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   addressResponse,
	}

	helpers.WriteResponseBody(w, webResponse)
}

func (c *AddressControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	addressId := vars["addressId"]

	c.AddressService.Delete(r.Context(), addressId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
	}
	helpers.WriteResponseBody(w, webResponse)
}

func (c *AddressControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	addressId := vars["addressId"]

	addressResponse := c.AddressService.FindById(r.Context(), addressId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   addressResponse,
	}
	helpers.WriteResponseBody(w, webResponse)
}

func (c *AddressControllerImpl) FindByUserId(w http.ResponseWriter, r *http.Request) {
	cookie, err := helpers.GetCookie(w, r)
	if err != nil {
		http.Redirect(w, r, "/api/users/login", http.StatusPermanentRedirect)
		return
	}

	userId := cookie.Value

	addressResponse := c.AddressService.FindByUserId(r.Context(), userId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   addressResponse,
	}
	helpers.WriteResponseBody(w, webResponse)
}
