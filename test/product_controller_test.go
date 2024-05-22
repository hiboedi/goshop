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
	"github.com/hiboedi/zenshop/app/middleware"
	"github.com/hiboedi/zenshop/app/models"
	"github.com/hiboedi/zenshop/app/repository"
	"gorm.io/gorm"
)

func createProduct(db *gorm.DB) (models.Product, models.Product) {
	tx := db.Begin()
	productRepo := repository.NewProductRepo()
	product1, _ := productRepo.Create(context.Background(), tx, models.Product{
		Name:        "Macbook",
		Price:       12000000,
		Stock:       2,
		Description: "Mobile laptop",
	})
	product2, _ := productRepo.Create(context.Background(), tx, models.Product{
		Name:        "Galaxy",
		Price:       10000000,
		Stock:       2,
		Description: "Mobile phone",
	})
	tx.Commit()

	return product1, product2
}

func trucateProduct(db *gorm.DB) {
	db.Exec("TRUNCATE products")
}

func TestCreateProductSucess(t *testing.T) {
	db := DBTestSetup()
	trucateProduct(db)
	truncateUser(db)
	router := SetUpRouter()

	user := createUser(db)

	token, _ := middleware.CreateToken(user.ID)

	requstBody := strings.NewReader(`{"name":"Macbook","price":12000000,"stock":2,"description":"Mobile laptop"}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/products", requstBody)
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
	assert.Equal(t, "Macbook", responseBody["data"].(map[string]interface{})["name"])
}

func TestCreateProductFailed(t *testing.T) {
	db := DBTestSetup()
	trucateProduct(db)
	truncateUser(db)
	router := SetUpRouter()

	user := createUser(db)

	token, _ := middleware.CreateToken(user.ID)

	requstBody := strings.NewReader(`{"name":"","price":0,"stock":0,"description":""}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/products", requstBody)
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

func TestGetAllProducts(t *testing.T) {
	db := DBTestSetup()
	trucateProduct(db)
	truncateUser(db)
	router := SetUpRouter()

	user := createUser(db)
	product1, product2 := createProduct(db)

	token, _ := middleware.CreateToken(user.ID)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/api/products", nil)
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

	fmt.Println(responseBody)

	var categories = responseBody["data"].([]interface{})

	categoryResponse1 := categories[0].(map[string]interface{})
	categoryResponse2 := categories[1].(map[string]interface{})

	assert.Equal(t, product1.Name, categoryResponse1["name"])

	assert.Equal(t, product2.Name, categoryResponse2["name"])
}

func TestGetProductById(t *testing.T) {
	db := DBTestSetup()
	trucateProduct(db)
	truncateUser(db)
	router := SetUpRouter()

	user := createUser(db)
	product, _ := createProduct(db)

	token, _ := middleware.CreateToken(user.ID)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/api/products/"+product.ID, nil)
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
	assert.Equal(t, product.ID, responseBody["data"].(map[string]interface{})["id"])
	assert.Equal(t, "Macbook", responseBody["data"].(map[string]interface{})["name"])
}

func TestGetProductByIdFailed(t *testing.T) {
	db := DBTestSetup()
	trucateProduct(db)
	truncateUser(db)
	router := SetUpRouter()

	user := createUser(db)
	createProduct(db)
	productIdErr := "kajdafasd"

	token, _ := middleware.CreateToken(user.ID)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/api/products/"+productIdErr, nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+token)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusInternalServerError, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)

	assert.Equal(t, http.StatusInternalServerError, int(responseBody["code"].(float64)))
	assert.Equal(t, "Internal Server Error", responseBody["status"])
}

func TestUpdateProductSucess(t *testing.T) {
	db := DBTestSetup()
	trucateProduct(db)
	truncateUser(db)
	router := SetUpRouter()

	user := createUser(db)
	product, _ := createProduct(db)

	token, _ := middleware.CreateToken(user.ID)

	requstBody := strings.NewReader(`{"name":"Galaxy book","price":10000000,"stock":5,"description":"Mobile laptop"}`)

	request := httptest.NewRequest(http.MethodPut, "http://localhost:8000/api/products/"+product.ID, requstBody)
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
	assert.Equal(t, "Galaxy book", responseBody["data"].(map[string]interface{})["name"])
}

func TestUpdateProductFailed(t *testing.T) {
	db := DBTestSetup()
	trucateProduct(db)
	truncateUser(db)
	router := SetUpRouter()

	user := createUser(db)
	product, _ := createProduct(db)

	token, _ := middleware.CreateToken(user.ID)

	requstBody := strings.NewReader(`{"name":"","price":0,"stock":0,"description":""}`)

	request := httptest.NewRequest(http.MethodPut, "http://localhost:8000/api/products/"+product.ID, requstBody)
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

func TestUpdateDeleteSuccess(t *testing.T) {
	db := DBTestSetup()
	router := SetUpRouter()
	trucateProduct(db)
	truncateUser(db)

	user := createUser(db)
	product, _ := createProduct(db)

	token, _ := middleware.CreateToken(user.ID)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:8000/api/products/"+product.ID, nil)
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
}

func TestUpdateDeleteFailed(t *testing.T) {
	db := DBTestSetup()
	router := SetUpRouter()
	trucateProduct(db)
	truncateUser(db)

	user := createUser(db)
	product, _ := createProduct(db)

	token, _ := middleware.CreateToken(user.ID)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:8000/api/products/"+product.ID+"asd", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+token)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusInternalServerError, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)

	assert.Equal(t, http.StatusInternalServerError, int(responseBody["code"].(float64)))
	assert.Equal(t, "Internal Server Error", responseBody["status"])
}
