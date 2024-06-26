package controllers

import (
	"net/http"
)

type OrderController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	FindAllOrByPaymentStatus(w http.ResponseWriter, r *http.Request)
}
