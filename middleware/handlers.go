package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"products/models"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_"github.com/lib/pq"
)

type response struct {
	ID int64 `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func createConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err!=nil{
		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err!=nil {
		panic(err)
	}

	err = db.Ping()

	if err!=nil {
		panic(err)
	}

	fmt.Println("Successfully connected to postgres..")
	return db
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err!=nil {
		log.Fatalf("Unable to convert string into int, %v", err)
	}

	product, err := getProduct(int64(id))
	if err!=nil {
		log.Fatalf("Unable to get product. %v", err)
	}

	json.NewEncoder(w).Encode(product)
}

func getProduct(id int64)(models.Product, error) {
	db := createConnection()
	defer db.Close()
	sqlStatement := `SELECT * FROM products WHERE id=$1`

	var product models.Product

	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&product.Id, &product.Name, &product.ShortDescription, &product.Description, &product.Price, &product.Created,&product.Updated)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Now rows were returned.")
		return product, nil
	case nil:
		return product,nil
	default:
		log.Fatalf("Unable to scan the row %v", err)
	}
	return product, err
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := getAllProducts()
	if err!=nil {
		log.Fatalf("Unable to get all the products %v",err)
	}
	json.NewEncoder(w).Encode(products)
}

func getAllProducts()([]models.Product, error) {
	db := createConnection()
	defer db.Close()
	sqlStatement := `SELECT * FROM products`

	var products []models.Product

	rows, err := db.Query(sqlStatement)
	if err!=nil {
		log.Fatalf("Unable to execute the query, %v",err)
	}

	defer rows.Close()
	for rows.Next(){
		var product models.Product
		err := rows.Scan(&product.Id, &product.Name, &product.ShortDescription,&product.Description, &product.Price, &product.Created, &product.Updated)
		if err!=nil {
			log.Fatalf("Unable to scan the row %v", err)
		}
		products = append(products, product)
	}
	return products, err
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err!=nil {
		log.Fatalf("Unable to decode the request body, %v", err)
	}

	insertID := insertProduct(product)

	res := response{
		ID: insertID,
		Message: "Product create successfully",
	}
	json.NewEncoder(w).Encode(res)
}

func insertProduct(product models.Product) int64{
	db := createConnection()
	defer db.Close()
	sqlStatement := `INSERT INTO products(name, shortDescription, description, price, created, updated) VALUES($1,$2,$3,$4,Now(),Now()) RETURNING id`

	var id int64

	err := db.QueryRow(sqlStatement, product.Name, product.ShortDescription, product.Description, product.Price).Scan(&id)
	if err!=nil{
		log.Fatalf("Unable to execute the query %v", err)
	}
	return id
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {

}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err!=nil {
		log.Fatalf("Unable to convert string into int, %v", err)
	}

	deletedRow := deleteProduct(int64(id))
	msg := fmt.Sprintf("Product deleted successfully %v", deletedRow)
	res := response {
		ID: int64(id),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}

func deleteProduct(id int64) int64 {
	db := createConnection()
	defer db.Close()
	sqlStatement := `DELETE FROM products WHERE id=$1`

	res, err := db.Exec(sqlStatement, id)
	if err!=nil {
		log.Fatalf("Unable to execute the query %v", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err!=nil {
		log.Fatalf("Error while checking the affected rows %v", err)
	}
	return rowsAffected
}