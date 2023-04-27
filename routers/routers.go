package routers

import (
	"products/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router{
	//Init router
	router := mux.NewRouter()

	//Route Handlers
	router.HandleFunc("/api/product/{id}", middleware.GetProduct).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/product", middleware.GetAllProducts).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/newproduct", middleware.CreateProduct).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/product/{id}", middleware.UpdateProduct).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deleteproduct/{id}", middleware.DeleteProduct).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/newcategory", middleware.CreateCategory).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/category/{id}", middleware.GetCategory).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/category", middleware.GetAllCategories).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/category/{id}", middleware.UpdateCategory).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deletecategory/{id}", middleware.DeleteCategory).Methods("DELETE", "OPTIONS")

	return router
}