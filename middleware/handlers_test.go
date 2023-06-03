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
		log.Fatalf("Failed to marshal update payload: %v", err)
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
