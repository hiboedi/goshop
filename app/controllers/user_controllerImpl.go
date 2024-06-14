package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hiboedi/zenshop/app/helpers"
	"github.com/hiboedi/zenshop/app/models"
	"github.com/hiboedi/zenshop/app/services"
	"github.com/hiboedi/zenshop/app/web"
)

type UserControllerImpl struct {
	UserService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (c *UserControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	userCreateRequest := models.UserCreate{}
	helpers.ToRequestBody(r, &userCreateRequest)

	userResponse := c.UserService.Create(r.Context(), userCreateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   userResponse,
	}

	helpers.WriteResponseBody(w, webResponse)
	http.Redirect(w, r, "/api/users/login", http.StatusOK)
}

func (c *UserControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	userUpdateRequest := models.UserUpdate{}
	helpers.ToRequestBody(r, &userUpdateRequest)

	vars := mux.Vars(r)
	userId := vars["userId"]
	userUpdateRequest.ID = userId

	userResponse := c.UserService.Update(r.Context(), userUpdateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   userResponse,
	}

	helpers.WriteResponseBody(w, webResponse)
}

func (c *UserControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userId := vars["userId"]

	c.UserService.Delete(r.Context(), userId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
	}

	helpers.WriteResponseBody(w, webResponse)
}

func (c *UserControllerImpl) Login(w http.ResponseWriter, r *http.Request) {
	userLogin := models.UserLogin{}
	helpers.ToRequestBody(r, &userLogin)

	userResponse, loggedIn := c.UserService.Login(r.Context(), userLogin)
	if loggedIn {
		helpers.SetUserCookie(w, r, userResponse)
		webResponse := web.WebResponse{
			Code:   http.StatusOK,
			Status: "Ok",
			Data:   userResponse,
		}
		helpers.WriteResponseBody(w, webResponse)
	} else {
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}
		helpers.WriteResponseBody(w, webResponse)
	}
}

func (c *UserControllerImpl) Logout(w http.ResponseWriter, r *http.Request) {
	helpers.DeleteCookieHandler(w, r)

	webResponse := web.WebResponse{
		Code:   http.StatusPermanentRedirect,
		Status: "Redirect",
	}
	helpers.WriteResponseBody(w, webResponse)
	http.Redirect(w, r, "/api/users/login", http.StatusPermanentRedirect)
}
