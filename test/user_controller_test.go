package test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/assert"
	"github.com/google/uuid"
	"github.com/hiboedi/zenshop/app/middleware"
	"github.com/hiboedi/zenshop/app/models"
	"github.com/hiboedi/zenshop/app/repository"
	"gorm.io/gorm"
)

func createUser(db *gorm.DB) models.User {
	tx := db.Begin()
	userRepo := repository.NewUserRepo()
	user, _ := userRepo.Create(context.Background(), tx, models.User{
		ID:       uuid.New().String(),
		Name:     "exam",
		Email:    "exam@mail.com",
		Password: "123123",
	})
	tx.Commit()

	return user
}

func truncateUser(db *gorm.DB) {
	db.Exec("TRUNCATE users")
}

func TestCreateUserSucess(t *testing.T) {
	db := DBTestSetup()
	truncateUser(db)
	router := SetUpRouter()

	requstBody := strings.NewReader(`{"name":"exam","email":"exam@mail.com","password":"123123"}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/users", requstBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "Ok", responseBody["status"])
	assert.Equal(t, "exam", responseBody["data"].(map[string]interface{})["name"])
}

func TestCreateUserFailed(t *testing.T) {
	db := DBTestSetup()
	truncateUser(db)
	router := SetUpRouter()

	requstBody := strings.NewReader(`{"name":"","email":"","password":""}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/users", requstBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BadRequest", responseBody["status"])
}

func TestLoginUserSucess(t *testing.T) {
	router := SetUpRouter()
	db := DBTestSetup()
	truncateUser(db)

	user := createUser(db)

	requstBody := strings.NewReader(`{"email":"exam@mail.com","password":"123123"}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/users/login", requstBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "Ok", responseBody["status"])
	assert.Equal(t, user.Name, responseBody["data"].(map[string]interface{})["name"])
}

func TestLoginUserFailed(t *testing.T) {
	router := SetUpRouter()
	db := DBTestSetup()
	truncateUser(db)

	createUser(db)

	requstBody := strings.NewReader(`{"email":"exam@mail.com","password":""}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/users/login", requstBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)

	assert.Equal(t, http.StatusUnauthorized, int(responseBody["code"].(float64)))
	assert.Equal(t, "Unauthorized", responseBody["status"])
}

func TestUpdateUserSuccess(t *testing.T) {
	router := SetUpRouter()
	db := DBTestSetup()
	truncateUser(db)

	user := createUser(db)

	token, _ := middleware.CreateToken(user.ID)

	requestBody := strings.NewReader(`{"name":"Budi","email":"exam@mail.com","password":"asdasd"}`)

	request := httptest.NewRequest(http.MethodPut, "http://localhost:8000/api/users/"+user.ID, requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+token)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "Ok", responseBody["status"])
	assert.Equal(t, "Budi", responseBody["data"].(map[string]interface{})["name"])
}

func TestUpdateUserFailed(t *testing.T) {
	router := SetUpRouter()
	db := DBTestSetup()
	truncateUser(db)

	user := createUser(db)

	token, _ := middleware.CreateToken(user.ID)

	requestBody := strings.NewReader(`{"name":"","email":"exam2@mail.com","password":""}`)

	request := httptest.NewRequest(http.MethodPut, "http://localhost:8000/api/users/"+user.ID, requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+token)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)

	assert.Equal(t, http.StatusBadRequest, int(responseBody["code"].(float64)))
	assert.Equal(t, "BadRequest", responseBody["status"])
}
