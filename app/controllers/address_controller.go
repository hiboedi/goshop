package controllers

import (
	"net/http"
)

type AddressController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
	FindByUserId(w http.ResponseWriter, r *http.Request)
}
