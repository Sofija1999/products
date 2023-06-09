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
	//Create a product
	product := models.Product{
		Name:             "jabuka",
		ShortDescription: "zelena jabuka",
		Description:      "zelena jabuka, kisela, srednje velicine",
		Price:            10,
		Quantity:         100,
		Category_id:      1,
	}

	//Convert the product in JSON format
	jsonProduct, err := json.Marshal(product)
	if err != nil {
		log.Fatalf("Failed to marshal product to JSON: %v", err)
	}

	//Create a new HTTP request with JSON product 
	req, err := http.NewRequest("POST", "/api/newproduct", bytes.NewBuffer(jsonProduct))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	//Create a new HTTP recorder to capture the response
	rr := httptest.NewRecorder()

	//Create a new router and handle the /api/newproduct endpoint
	router := mux.NewRouter()
	router.HandleFunc("/api/newproduct", CreateProduct)

	//Serve the HTTP request using the router and record the response
	router.ServeHTTP(rr, req)

	//Check if the response status is as expected
	if rr.Code != http.StatusOK {
		log.Fatalf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	//Unmarshal the response body into a response
	var res response
	err = json.Unmarshal(rr.Body.Bytes(), &res)
	if err != nil {
		log.Fatalf("Failed to unmarshal response body: %v", err)
	}

	//Get the product with the returned id
	product1, err := getProduct(res.Id)
	id := product1.Id

	//Define  the expected response
	expectedRes := response{
		Id:    id,
		Message: "Product create successfully",
	}

	//Check if the returned ID is less than zero
	if res.Id<0{
		log.Fatalf("Id is less than zero %v", err)
	}

	//Check if the actual response matches the expected response
	if res.Id != expectedRes.Id || res.Message != expectedRes.Message{
		log.Fatalf("Expected response %v, but got %v", expectedRes, res)
	}

}

func TestDeleteProduct(t *testing.T) {
	//Create a product
	product := models.Product{
		Name:             "jabuka",
		ShortDescription: "zelena jabuka",
		Description:      "zelena jabuka, kisela, srednje velicine",
		Price:            10,
		Quantity:         100,
		Category_id:      1,
	}
	
	//Convert the product in JSON format
	jsonProduct, err := json.Marshal(product)
	if err != nil {
		log.Fatalf("Failed to marshal product to JSON: %v", err)
	}
	
	//Create a new HTTP request with JSON product 
	req, err := http.NewRequest("POST", "/api/newproduct", bytes.NewBuffer(jsonProduct))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}
	
	//Create a new HTTP recorder to capture the response
	rr := httptest.NewRecorder()
	
	//Create a new router and handle the /api/newproduct endpoint
	router := mux.NewRouter()
	router.HandleFunc("/api/newproduct", CreateProduct)
	
	//Serve the HTTP request using the router and record the response
	router.ServeHTTP(rr, req)
	
	//Check if the response status is as expected
	if rr.Code != http.StatusOK {
		log.Fatalf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	//Unmarshal the response body into a response
	var res response
	err = json.Unmarshal(rr.Body.Bytes(), &res)
	if err != nil {
		log.Fatalf("Failed to unmarshal response body: %v", err)
	}
	
	//Get the product with the returned id
	product1, err := getProduct(res.Id)
	id := product1.Id

	//Create a endpoint for delete product with returned product id and create HTTP request 
	endpoint := fmt.Sprintf("/api/deleteproduct/%d", id)
	req1, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	//Create a new HTTP recorder to capture the response
	rr1 := httptest.NewRecorder()

	//Create a new router and handle the /api/deleteproduct/{id} endpoint
	router1 := mux.NewRouter()
	router1.HandleFunc("/api/deleteproduct/{id}", DeleteProduct)

	//Serve the HTTP request using the router and record the response
	router1.ServeHTTP(rr1, req1)

	//Check if the response status is as expected
	if rr1.Code != http.StatusOK {
		log.Fatalf("Expected status code %d, but got %d", http.StatusOK, rr1.Code)
	}

	//Define  the expected response
	var res1 response
	err = json.Unmarshal(rr1.Body.Bytes(), &res1)
	if err != nil {
		log.Fatalf("Failed to unmarshal response body: %v", err)
	}

	//Define  the expected response
	expectedRes := response{
		Id:      id,
		Message: "Product deleted successfully 1",
	}

	//Check if the actual response matches the expected response
	if res1.Id != expectedRes.Id || res1.Message != expectedRes.Message {
		log.Fatalf("Expected response %v, but got %v", expectedRes, res1)
	}

	log.Fatalf("We are successfully deleted product with id %v", id)
}

func TestGetProduct(t *testing.T) {
	//Create a product
	product := models.Product{
		Name:             "jabuka",
		ShortDescription: "zelena jabuka",
		Description:      "zelena jabuka, kisela, srednje velicine",
		Price:            10,
		Quantity:         100,
		Category_id:      1,
	}
	
	//Convert the product in JSON format
	jsonProduct, err := json.Marshal(product)
	if err != nil {
		log.Fatalf("Failed to marshal product to JSON: %v", err)
	}
	
	//Create a new HTTP request with JSON product 
	req, err := http.NewRequest("POST", "/api/newproduct", bytes.NewBuffer(jsonProduct))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}
	
	//Create a new HTTP recorder to capture the response
	rr := httptest.NewRecorder()
	
	//Create a new router and handle the /api/newproduct endpoint
	router := mux.NewRouter()
	router.HandleFunc("/api/newproduct", CreateProduct)
	
	//Serve the HTTP request using the router and record the response
	router.ServeHTTP(rr, req)
	
	//Check if the response status is as expected
	if rr.Code != http.StatusOK {
		log.Fatalf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	//Unmarshal the response body into a response
	var res response
	err = json.Unmarshal(rr.Body.Bytes(), &res)
	if err != nil {
		log.Fatalf("Failed to unmarshal response body: %v", err)
	}
	
	//Get the product with the returned id
	product1, err := getProduct(res.Id)
	id := product1.Id

	//Create a endpoint for get product with returned product id and create HTTP request 
	endpoint := fmt.Sprintf("/api/product/%d", id)
	req1, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	//Create a new HTTP recorder to capture the response
	rr1 := httptest.NewRecorder()

	//Create a new router and handle the /api/product/{id} endpoint
	router1 := mux.NewRouter()
	router1.HandleFunc("/api/product/{id}", GetProduct)

	//Serve the HTTP request using the router and record the response
	router1.ServeHTTP(rr1, req1)

	//Check if the response status is as expected
	if rr1.Code != http.StatusOK {
		log.Fatalf("Expected status code %d, but got %d", http.StatusOK, rr1.Code)
	}

	//Define the expected product
	var product2 models.Product
	err = json.Unmarshal(rr1.Body.Bytes(), &product2)
	if err != nil {
		log.Fatalf("Failed to unmarshal response body: %v", err)
	}

	//Check if the actual response matches the expected response
	if product1.Id != product2.Id || product1.Name != product2.Name {
		t.Errorf("Expected response %v, but got %v", product2, product1)
	}

}

func TestUpdateProduct(t *testing.T) {
	//Create a product
	product := models.Product{
		Name:             "jabuka",
		ShortDescription: "zelena jabuka",
		Description:      "zelena jabuka, kisela, srednje velicine",
		Price:            10,
		Quantity:         100,
		Category_id:      1,
	}
	
	//Convert the product in JSON format
	jsonProduct, err := json.Marshal(product)
	if err != nil {
		log.Fatalf("Failed to marshal product to JSON: %v", err)
	}
	
	//Create a new HTTP request with JSON product 
	req, err := http.NewRequest("POST", "/api/newproduct", bytes.NewBuffer(jsonProduct))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}
	
	//Create a new HTTP recorder to capture the response
	rr := httptest.NewRecorder()
	
	//Create a new router and handle the /api/newproduct endpoint
	router := mux.NewRouter()
	router.HandleFunc("/api/newproduct", CreateProduct)
	
	//Serve the HTTP request using the router and record the response
	router.ServeHTTP(rr, req)
	
	//Check if the response status is as expected
	if rr.Code != http.StatusOK {
		log.Fatalf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	//Unmarshal the response body into a response
	var res response
	err = json.Unmarshal(rr.Body.Bytes(), &res)
	if err != nil {
		log.Fatalf("Failed to unmarshal response body: %v", err)
	}
	
	//Get the product with the returned id
	product1, err := getProduct(res.Id)
	id := product1.Id

	//Create updateProduct
	updateProduct := models.Product{
		Name:             "Updated Product",
		ShortDescription: "Updated Short Description",
		Description:      "Updated Detailed Description",
		Price:            19.99,
		Updated:          time.Now(),
		Quantity:         50,
		Category_id:      1,
	}

	//Convert updateProduct to JSON format
	product2, err := json.Marshal(updateProduct)
	if err != nil {
		log.Fatalf("Failed to marshal update payload: %v", err)
	}

	//Create a endpoint for update product with returned product id and create HTTP request
	endpoint := fmt.Sprintf("/api/product/%d", id)
	req1, err := http.NewRequest("PUT", endpoint, bytes.NewBuffer(product2))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	//Create a new HTTP recorder to capture the response
	rr1 := httptest.NewRecorder()

	//Create a new router and handle the /api/product/{id} endpoint
	router1 := mux.NewRouter()
	router1.HandleFunc("/api/product/{id}", UpdateProduct)

	//Serve the HTTP request using the router and record the response
	router1.ServeHTTP(rr1, req1)

	//Check if the response status is as expected
	if rr1.Code != http.StatusOK {
		log.Fatalf("Expected status code %d, but got %d", http.StatusOK, rr1.Code)
	}

	//Unmarshal the response body into a response
	var res1 response
	err = json.Unmarshal(rr1.Body.Bytes(), &res1)
	if err != nil {
		log.Fatalf("Failed to unmarshal response body: %v", err)
	}

	//Check if the expected message equal got message
	expectedMsg := "Product updated successfully 1"
	if res1.Message != expectedMsg {
		log.Fatalf("Expected message '%s', but got '%s'", expectedMsg, res1.Message)
	}

	updatedProduct, err := getProduct(id)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	if updatedProduct.Name != updateProduct.Name || updatedProduct.Description != updateProduct.Description {
		t.Errorf("Expected product %+v, but got %+v", updateProduct, updatedProduct)
	}

}

func TestGetAllProduct(t *testing.T){
	//Create a new HTTP request 
	req, err := http.NewRequest("GET", "/api/product", nil)
	if err!=nil{
		log.Fatalf("Failed to create request %v", err)
	}

	//Create a new HTTP recorder to capture the response
	rr := httptest.NewRecorder()

	//Create a new router and handle the /api/product endpoint
	router := mux.NewRouter()
	router.HandleFunc("/api/product", GetAllProducts)

	//Serve the HTTP request using the router and record the response
	router.ServeHTTP(rr, req)

	//Check if the response status is as expected
	if rr.Code != http.StatusOK{
		log.Fatalf("Expected status code %v,  got %v", http.StatusOK, rr.Code)
	}

	//Unmarshal the response body into a products
	var products []models.Product
	err = json.Unmarshal(rr.Body.Bytes(), &products)
	if err!=nil{
		log.Fatalf("Failed to unmarshal response %v", err)
	}

	//Check if we dont have products in database
	if len(products) == 0{
		log.Fatalf("We dont have products in our database")
	}

	//Check if number of products is > 0
	if len(products) > 0{
		fmt.Printf("We have products in our database")
	}

}

func TestCreateCategory(t *testing.T){
	//Create a category
	category := models.Category{
		Category_name: "kupaci",
	}

	//Convert the category in JSON format
	jsonCategory, err := json.Marshal(category)
	if err != nil {
		log.Fatalf("Failed to marshal category to JSON: %v", err)
	}

	//Create a new HTTP request with JSON category 
	req, err := http.NewRequest("POST", "/api/newcategory", bytes.NewBuffer(jsonCategory))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	//Create a new HTTP recorder to capture the response
	rr := httptest.NewRecorder()

	//Create a new router and handle the /api/newcategory endpoint
	router := mux.NewRouter()
	router.HandleFunc("/api/newcategory", CreateCategory)

	//Serve the HTTP request using the router and record the response
	router.ServeHTTP(rr, req)

	//Check if the response status is as expected
	if rr.Code != http.StatusOK {
		log.Fatalf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	//Unmarshal the response body into a response
	var res response
	err = json.Unmarshal(rr.Body.Bytes(), &res)
	if err != nil {
		log.Fatalf("Failed to unmarshal response body: %v", err)
	}

	//Get the category with the returned id
	category1, err := getCategory(res.Id)
	id := category1.Category_id

	//Define  the expected response
	expectedRes := response{
		Id:    id,
		Message: "Category create successfully",
	}

	//Check if the returned ID is less than zero
	if res.Id<0{
		log.Fatalf("Id is less than zero %v", err)
	}

	//Check if the actual response matches the expected response
	if res.Id != expectedRes.Id || res.Message != expectedRes.Message{
		log.Fatalf("Expected response %v, but got %v", expectedRes, res)
	}
}

func TestGetCategory(t *testing.T){
	//Create a category
	category := models.Category{
		Category_name: "carape",
	}
	
	//Convert the category in JSON format
	jsonCategory, err := json.Marshal(category)
	if err != nil {
		log.Fatalf("Failed to marshal category to JSON: %v", err)
	}
	
	//Create a new HTTP request with JSON category 
	req, err := http.NewRequest("POST", "/api/newcategory", bytes.NewBuffer(jsonCategory))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}
	
	//Create a new HTTP recorder to capture the response
	rr := httptest.NewRecorder()
	
	//Create a new router and handle the /api/newcategory endpoint
	router := mux.NewRouter()
	router.HandleFunc("/api/newcategory", CreateCategory)
	
	//Serve the HTTP request using the router and record the response
	router.ServeHTTP(rr, req)
	
	//Check if the response status is as expected
	if rr.Code != http.StatusOK {
		log.Fatalf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	//Unmarshal the response body into a response
	var res response
	err = json.Unmarshal(rr.Body.Bytes(), &res)
	if err != nil {
		log.Fatalf("Failed to unmarshal response body: %v", err)
	}
	
	//Get the category with the returned id
	category1, err := getCategory(res.Id)
	id := category1.Category_id

	//Create a endpoint for get category with returned category id and create HTTP request 
	endpoint := fmt.Sprintf("/api/category/%d", id)
	req1, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	//Create a new HTTP recorder to capture the response
	rr1 := httptest.NewRecorder()

	//Create a new router and handle the /api/category/{id} endpoint
	router1 := mux.NewRouter()
	router1.HandleFunc("/api/category/{id}", GetCategory)

	//Serve the HTTP request using the router and record the response
	router1.ServeHTTP(rr1, req1)

	//Check if the response status is as expected
	if rr1.Code != http.StatusOK {
		log.Fatalf("Expected status code %d, but got %d", http.StatusOK, rr1.Code)
	}

	//Define the expected category
	var category2 models.Category
	err = json.Unmarshal(rr1.Body.Bytes(), &category2)
	if err != nil {
		log.Fatalf("Failed to unmarshal response body: %v", err)
	}

	//Check if the actual response matches the expected response
	if category1.Category_id != category2.Category_id ||  category1.Category_name != category2.Category_name {
		t.Errorf("Expected response %v, but got %v", category2, category1)
	}
	fmt.Println(category2)
}

func TestUpdateCategory(t *testing.T){
	//Create a category
	category := models.Category{
		Category_name: "carape",
	}
	
	//Convert the category in JSON format
	jsonCategory, err := json.Marshal(category)
	if err != nil {
		log.Fatalf("Failed to marshal category to JSON: %v", err)
	}
	
	//Create a new HTTP request with JSON category 
	req, err := http.NewRequest("POST", "/api/newcategory", bytes.NewBuffer(jsonCategory))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}
	
	//Create a new HTTP recorder to capture the response
	rr := httptest.NewRecorder()
	
	//Create a new router and handle the /api/newcategory endpoint
	router := mux.NewRouter()
	router.HandleFunc("/api/newcategory", CreateCategory)
	
	//Serve the HTTP request using the router and record the response
	router.ServeHTTP(rr, req)
	
	//Check if the response status is as expected
	if rr.Code != http.StatusOK {
		log.Fatalf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	//Unmarshal the response body into a response
	var res response
	err = json.Unmarshal(rr.Body.Bytes(), &res)
	if err != nil {
		log.Fatalf("Failed to unmarshal response body: %v", err)
	}
	
	//Get the category with the returned id
	category1, err := getCategory(res.Id)
	id := category1.Category_id

	//Create updateCategory
	updateCategory := models.Category{
		Category_name: "jakne",
		Updated_at: time.Now(),
	}

	//Convert updateCategory to JSON format
	category2, err := json.Marshal(updateCategory)
	if err != nil {
		log.Fatalf("Failed to marshal update category: %v", err)
	}

	//Create a endpoint for update category with returned category id and create HTTP request
	endpoint := fmt.Sprintf("/api/category/%d", id)
	req1, err := http.NewRequest("PUT", endpoint, bytes.NewBuffer(category2))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	//Create a new HTTP recorder to capture the response
	rr1 := httptest.NewRecorder()

	//Create a new router and handle the /api/category/{id} endpoint
	router1 := mux.NewRouter()
	router1.HandleFunc("/api/category/{id}", UpdateCategory)

	//Serve the HTTP request using the router and record the response
	router1.ServeHTTP(rr1, req1)

	//Check if the response status is as expected
	if rr1.Code != http.StatusOK {
		log.Fatalf("Expected status code %d, but got %d", http.StatusOK, rr1.Code)
	}

	//Unmarshal the response body into a response
	var res1 response
	err = json.Unmarshal(rr1.Body.Bytes(), &res1)
	if err != nil {
		log.Fatalf("Failed to unmarshal response body: %v", err)
	}

	//Check if the expected message equal got message
	expectedMsg := "Category updated successfully  1"
	if res1.Message != expectedMsg {
		log.Fatalf("Expected message '%s', but got '%s'", expectedMsg, res1.Message)
	}

	updatedCategory, err := getCategory(id)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	if updatedCategory.Category_name != updateCategory.Category_name {
		t.Errorf("Expected category %+v, but got %+v", updateCategory, updatedCategory)
	}


}

func TestDeleteCategory(t *testing.T){
	//Create a category
	category := models.Category{
		Category_name: "carape",
	}
	
	//Convert the category in JSON format
	jsonCategory, err := json.Marshal(category)
	if err != nil {
		log.Fatalf("Failed to marshal category to JSON: %v", err)
	}
	
	//Create a new HTTP request with JSON category 
	req, err := http.NewRequest("POST", "/api/newcategory", bytes.NewBuffer(jsonCategory))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}
	
	//Create a new HTTP recorder to capture the response
	rr := httptest.NewRecorder()
	
	//Create a new router and handle the /api/newcategory endpoint
	router := mux.NewRouter()
	router.HandleFunc("/api/newcategory", CreateCategory)
	
	//Serve the HTTP request using the router and record the response
	router.ServeHTTP(rr, req)
	
	//Check if the response status is as expected
	if rr.Code != http.StatusOK {
		log.Fatalf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	//Unmarshal the response body into a response
	var res response
	err = json.Unmarshal(rr.Body.Bytes(), &res)
	if err != nil {
		log.Fatalf("Failed to unmarshal response body: %v", err)
	}
	
	//Get the category with the returned id
	category1, err := getCategory(res.Id)
	id := category1.Category_id

	//Create a endpoint for delete category with returned category id and create HTTP request 
	endpoint := fmt.Sprintf("/api/deletecategory/%d", id)
	req1, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	//Create a new HTTP recorder to capture the response
	rr1 := httptest.NewRecorder()

	//Create a new router and handle the /api/deletecategory/{id} endpoint
	router1 := mux.NewRouter()
	router1.HandleFunc("/api/deletecategory/{id}", DeleteCategory)

	//Serve the HTTP request using the router and record the response
	router1.ServeHTTP(rr1, req1)

	//Check if the response status is as expected
	if rr1.Code != http.StatusOK {
		log.Fatalf("Expected status code %d, but got %d", http.StatusOK, rr1.Code)
	}

	//Define  the expected response
	var res1 response
	err = json.Unmarshal(rr1.Body.Bytes(), &res1)
	if err != nil {
		log.Fatalf("Failed to unmarshal response body: %v", err)
	}

	//Define  the expected response
	expectedRes := response{
		Id:      id,
		Message: "Category deleted successfully 1",
	}

	//Check if the actual response matches the expected response
	if res1.Id != expectedRes.Id || res1.Message != expectedRes.Message {
		log.Fatalf("Expected response %v, but got %v", expectedRes, res1)
	}

	fmt.Printf("We are successfully deleted category with id %v", id)

}