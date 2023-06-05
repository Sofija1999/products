package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"products/models"
	"testing"
	"time"

	"github.com/gorilla/mux"
)

func TestCreateProduct(t *testing.T) {
	product := models.Product{
		Name:             "jabuka",
		ShortDescription: "zelena jabuka",
		Description:      "zelena jabuka, kisela, srednje velicine",
		Price:            10,
		Quantity:         100,
		Category_id:      1,
	}

	jsonProduct, err := json.Marshal(product)
	if err != nil {
		log.Fatalf("Failed to marshal product to JSON: %v", err)
	}

	req, err := http.NewRequest("POST", "/api/newproduct", bytes.NewBuffer(jsonProduct))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/newproduct", CreateProduct)

	router.ServeHTTP(rr, req)

	//CreateProduct(rr, req)

	if rr.Code != http.StatusOK {
		log.Fatalf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	expectedRes := response{
		Id:      15,
		Message: "Product create successfully",
	}

	var res response
	err = json.Unmarshal(rr.Body.Bytes(), &res)
	if err != nil {
		log.Fatalf("Failed to unmarshal response body: %v", err)
	}

	if res.Id != expectedRes.Id || res.Message != expectedRes.Message {
		log.Fatalf("Expected response %v, but got %v", expectedRes, res)
	}
}

func TestDeleteProduct(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/api/deleteproduct/11", nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/deleteproduct/{id}", DeleteProduct)

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		log.Fatalf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	expectedRes := response{
		Id:      11,
		Message: "Product deleted successfully 1",
	}

	var res response
	err = json.Unmarshal(rr.Body.Bytes(), &res)
	if err != nil {
		log.Fatalf("Failed to unmarshal response body: %v", err)
	}

	if res.Id != expectedRes.Id || res.Message != expectedRes.Message {
		log.Fatalf("Expected response %v, but got %v", expectedRes, res)
	}
}

func TestGetProduct(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/product/4", nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/product/{id}", GetProduct)

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		log.Fatalf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	expectedProduct := models.Product{
		Id:               4,
		Name:             "plava trenerka",
		ShortDescription: "plava trenerka koja ima po sebi cvetice",
		Description:      "potrebno je da se pere na 30 stepeni u masini na programu cotton",
		Price:            2000,
		Created:          time.Date(2023, 4, 25, 13, 15, 37, 0, time.UTC),
		Updated:          time.Date(2023, 4, 25, 13, 18, 45, 0, time.UTC),
		Quantity:         5,
		Category_id:      1,
	}

	var product models.Product
	err = json.Unmarshal(rr.Body.Bytes(), &product)
	if err != nil {
		log.Fatalf("Failed to unmarshal response body: %v", err)
	}

	if product.Id != expectedProduct.Id || product.Name != expectedProduct.Name {
		t.Errorf("Expected response%v, but got %v", expectedProduct, product)
	}

	fmt.Println(product)

}

func TestUpdateProduct(t *testing.T) {
	updateProduct := models.Product{
		Id:               14,
		Name:             "Updated Product",
		ShortDescription: "Updated Short Description",
		Description:      "Updated Detailed Description",
		Price:            19.99,
		Updated:          time.Now(),
		Quantity:         50,
		Category_id:      1,
	}

	product, err := json.Marshal(updateProduct)
	if err != nil {
		log.Fatalf("Failed to marshal update product: %v", err)
	}

	req, err := http.NewRequest("PUT", "/api/product/14", bytes.NewBuffer(product))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/product/{id}", UpdateProduct)

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		log.Fatalf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	var res response
	err = json.Unmarshal(rr.Body.Bytes(), &res)
	if err != nil {
		log.Fatalf("Failed to unmarshal response body: %v", err)
	}

	expectedMsg := "Product updated successfully 1"
	if res.Message != expectedMsg {
		log.Fatalf("Expected message '%s', but got '%s'", expectedMsg, res.Message)
	}

	updatedProduct, err := getProduct(updateProduct.Id)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	if updatedProduct.Name != updateProduct.Name || updatedProduct.Description != updateProduct.Description {
		t.Errorf("Expected product %+v, but got %+v", updateProduct, updatedProduct)
	}

}

func TestGetAllProduct(t *testing.T){
	req, err := http.NewRequest("GET", "/api/product", nil)
	if err!=nil{
		log.Fatalf("Failed to create request %v", err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/product", GetAllProducts)

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK{
		log.Fatalf("Expected status code %v,  got %v", http.StatusOK, rr.Code)
	}

	var products []models.Product
	err = json.Unmarshal(rr.Body.Bytes(), &products)
	if err!=nil{
		log.Fatalf("Failed to unmarshal response %v", err)
	}

	if len(products) != 5{
		log.Fatalf("Expected 5 products, but got %v", len(products))
	}

	expectedProduct := models.Product{
		Id:               4,
		Name:             "plava trenerka",
		ShortDescription: "plava trenerka koja ima po sebi cvetice",
		Description:      "potrebno je da se pere na 30 stepeni u masini na programu cotton",
		Price:            2000,
		Created:          time.Date(2023, 4, 25, 13, 15, 37, 0, time.UTC),
		Updated:          time.Date(2023, 4, 25, 13, 18, 45, 0, time.UTC),
		Quantity:         5,
		Category_id:      1,
	}

	if products[0].Id != expectedProduct.Id || products[0].Name != expectedProduct.Name {
		t.Errorf("Expected product %v, but got %v", expectedProduct, products[0])
	}
}

func TestCreateCategory(t *testing.T){
	category := models.Category{
		Category_name: "obuca",
	}

	jsonProduct, err := json.Marshal(category)
	if err != nil {
		log.Fatalf("Failed to marshal category to JSON: %v", err)
	}

	req, err := http.NewRequest("POST", "/api/newcategory", bytes.NewBuffer(jsonProduct))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/newcategory", CreateCategory)

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		log.Fatalf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	expectedRes := response{
		Id:      6,
		Message: "Category create successfully",
	}

	var res response
	err = json.Unmarshal(rr.Body.Bytes(), &res)
	if err != nil {
		log.Fatalf("Failed to unmarshal response body: %v", err)
	}

	if res.Id != expectedRes.Id || res.Message != expectedRes.Message {
		log.Fatalf("Expected response %v, but got %v", expectedRes, res)
	}
}

func TestGetCategory(t *testing.T){
	req, err := http.NewRequest("GET", "/api/category/6", nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/category/{id}", GetCategory)

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		log.Fatalf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	expectedCategory := models.Category{
		Category_id:               6,
		Category_name:          "obuca",
		Created_at:          time.Date(2023, 6, 5, 16, 11, 50, 0, time.UTC),
		Updated_at:          time.Date(2023, 6, 5, 16, 12, 18, 0, time.UTC),
	}

	var category models.Category
	err = json.Unmarshal(rr.Body.Bytes(), &category)
	if err != nil {
		log.Fatalf("Failed to unmarshal response body: %v", err)
	}

	if category.Category_id != expectedCategory.Category_id || category.Category_name != expectedCategory.Category_name {
		t.Errorf("Expected response %v, but got %v", expectedCategory, category)
	}

	fmt.Println(category)
}

func TestUpdateCategory(t *testing.T){
	updateCategory := models.Category{
		Category_id:      5,
		Category_name:   "Updated Product",
		Updated_at:          time.Now(),
	}

	category, err := json.Marshal(updateCategory)
	if err != nil {
		log.Fatalf("Failed to marshal update category: %v", err)
	}

	req, err := http.NewRequest("PUT", "/api/category/5", bytes.NewBuffer(category))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/category/{id}", UpdateCategory)

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		log.Fatalf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	var res response
	err = json.Unmarshal(rr.Body.Bytes(), &res)
	if err != nil {
		log.Fatalf("Failed to unmarshal response body: %v", err)
	}

	expectedMsg := "Category updated successfully  1"
	if res.Message != expectedMsg {
		log.Fatalf("Expected message '%s', but got '%s'", expectedMsg, res.Message)
	}

	updatedCategory, err := getCategory(updateCategory.Category_id)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	if updatedCategory.Category_name != updateCategory.Category_name {
		t.Errorf("Expected product %+v, but got %+v", updateCategory, updateCategory)
	}

}

func TestDeleteCategory(t *testing.T){
	req, err := http.NewRequest("DELETE", "/api/deletecategory/3", nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/deletecategory/{id}", DeleteCategory)

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		log.Fatalf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	expectedRes := response{
		Id:      3,
		Message: "Category deleted successfully 1",
	}

	var res response
	err = json.Unmarshal(rr.Body.Bytes(), &res)
	if err != nil {
		log.Fatalf("Failed to unmarshal response body: %v", err)
	}

	if res.Id != expectedRes.Id || res.Message != expectedRes.Message {
		log.Fatalf("Expected response %v, but got %v", expectedRes, res)
	}
}