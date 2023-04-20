package main

import (
	"fmt"
	"log"
	"net/http"
	"products/routers"

)

func main() {

	r := routers.Router()
	fmt.Println("Starting server on the port 8080..")

	log.Fatal(http.ListenAndServe(":8080", r))

}