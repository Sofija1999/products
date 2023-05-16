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

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func createConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to postgres..")
	return db
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert string into int, %v", err)
	}

	product, err := getProduct(int64(id))
	if err != nil {
		log.Fatalf("Unable to get product. %v", err)
	}

	json.NewEncoder(w).Encode(product)
}

func getProduct(id int64) (models.Product, error) {
	db := createConnection()
	defer db.Close()
	sqlStatement := `SELECT * FROM products WHERE id=$1`

	var product models.Product

	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&product.Id, &product.Name, &product.ShortDescription, &product.Description, &product.Price, &product.Created, &product.Updated, &product.Quantity, &product.Category_id)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Now rows were returned.")
		return product, nil
	case nil:
		return product, nil
	default:
		log.Fatalf("Unable to scan the row %v", err)
	}
	return product, err
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := getAllProducts()
	if err != nil {
		log.Fatalf("Unable to get all the products %v", err)
	}
	json.NewEncoder(w).Encode(products)
}

func getAllProducts() ([]models.Product, error) {
	db := createConnection()
	defer db.Close()
	sqlStatement := `SELECT * FROM products`

	var products []models.Product

	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Unable to execute the query, %v", err)
	}

	defer rows.Close()
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.Id, &product.Name, &product.ShortDescription, &product.Description, &product.Price, &product.Created, &product.Updated, &product.Quantity, &product.Category_id)
		if err != nil {
			log.Fatalf("Unable to scan the row %v", err)
		}
		products = append(products, product)
	}
	return products, err
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Fatalf("Unable to decode the request body, %v", err)
	}

	insertID := insertProduct(product)

	res := response{
		ID:      insertID,
		Message: "Product create successfully",
	}
	json.NewEncoder(w).Encode(res)
}

func insertProduct(product models.Product) int64 {
	db := createConnection()
	defer db.Close()
	sqlStatement := `INSERT INTO products(name, shortDescription, description, price, created, updated, quantity, category_id) VALUES($1,$2,$3,$4,Now(),Now(),$5, $6) RETURNING id`

	var id int64

	err := db.QueryRow(sqlStatement, product.Name, product.ShortDescription, product.Description, product.Price, product.Quantity, product.Category_id).Scan(&id)
	if err != nil {
		log.Fatalf("Unable to execute the query %v", err)
	}
	return id
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert string into int %v", err)
	}

	var product models.Product

	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Fatalf("Unable to decode the request %v", err)
	}

	updatedRows := updateProduct(int64(id), product)
	msg := fmt.Sprintf("Product updated successfully %v", updatedRows)
	res := response{
		ID:      int64(id),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}

func updateProduct(id int64, product models.Product) int64 {
	db := createConnection()
	defer db.Close()
	sqlStatement := `UPDATE products SET name=$2, shortdescription=$3, description=$4, price=$5, updated=Now(), quantity=$6, category_id=$7 WHERE id=$1`

	res, err := db.Exec(sqlStatement, id, product.Name, product.ShortDescription, product.Description, product.Price, product.Quantity, product.Category_id)
	if err != nil {
		log.Fatalf("Unable to execute the query %v", err)
	}

	rowAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows %v", err)
	}
	return rowAffected
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert string into int, %v", err)
	}

	deletedRow := deleteProduct(int64(id))
	msg := fmt.Sprintf("Product deleted successfully %v", deletedRow)
	res := response{
		ID:      int64(id),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}

func deleteProduct(id int64) int64 {
	db := createConnection()
	defer db.Close()
	sqlStatement := `DELETE FROM products WHERE id=$1`

	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		log.Fatalf("Unable to execute the query %v", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows %v", err)
	}
	return rowsAffected
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category

	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		log.Fatalf("Unable to decode the request body %v", err)
	}

	insertID := insertCategory(category)

	res := response{
		ID:      insertID,
		Message: "Category create successfully",
	}

	json.NewEncoder(w).Encode(res)
}

func insertCategory(category models.Category) int64 {
	db := createConnection()
	defer db.Close()
	sqlStatement := `INSERT INTO categories(category_name,created_at,updated_at) VALUES ($1, Now(), Now()) RETURNING category_id`

	var id int64

	err := db.QueryRow(sqlStatement, category.Category_name).Scan(&id)
	if err != nil {
		log.Fatalf("Unable to execute the query %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	return id
}

func GetCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert string into int %v", err)
	}

	category, err := getCategory(int64(id))
	if err != nil {
		log.Fatalf("Unable to get category %v", err)
	}

	json.NewEncoder(w).Encode(category)
}

func getCategory(id int64) (models.Category, error) {
	db := createConnection()
	defer db.Close()
	sqlStatement := `SELECT * FROM categories WHERE category_id=$1`

	var category models.Category

	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&category.Category_id, &category.Category_name, &category.Created_at, &category.Updated_at)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Now rows were returned")
		return category, nil
	case nil:
		return category, nil
	default:
		log.Fatalf("Unable to scan the row %v", err)
	}
	return category, err
}

func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := getAllCategories()
	if err != nil {
		log.Fatalf("Unable to get all the categories %v", err)
	}

	json.NewEncoder(w).Encode(categories)
}

func getAllCategories() ([]models.Category, error) {
	db := createConnection()
	defer db.Close()
	sqlStatement := `SELECT * FROM categories`

	var categories []models.Category

	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Unable to execute the query %v", err)
	}

	defer rows.Close()
	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.Category_id, &category.Category_name, &category.Created_at, &category.Updated_at)
		if err != nil {
			log.Fatalf("Unable to scan the row %v", err)
		}
		categories = append(categories, category)
	}
	return categories, err
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert strig into int %v", err)
	}

	var category models.Category

	err = json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		log.Fatalf("Unable to decode the request %v", err)
	}

	updatedRow := updateCategory(int64(id), category)
	msg := fmt.Sprintf("Category updated successfully  %v", updatedRow)
	res := response{
		ID:      int64(id),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}

func updateCategory(id int64, category models.Category) int64 {
	db := createConnection()
	defer db.Close()
	sqlStatement := `UPDATE categories SET category_name=$2, updated_at=Now() WHERE category_id=$1`

	res, err := db.Exec(sqlStatement, id, category.Category_name)
	if err != nil {
		log.Fatalf("Unable to execute the query %v", err)
	}

	rowAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows %v", err)
	}
	return rowAffected
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert string into int %v", err)
	}

	deletedRows := deleteCategory(int64(id))
	msg := fmt.Sprintf("Category deleted successfully %v", deletedRows)
	res := response{
		ID:      int64(id),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}

func deleteCategory(id int64) int64 {
	db := createConnection()
	defer db.Close()
	sqlStatement := `DELETE FROM categories WHERE category_id=$1`

	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		log.Fatalf("Unable to execute the query %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows, %v", rowsAffected)
	}

	return rowsAffected
}

func UserRegister(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatalf("Unable to decode the request body %v", err)
	}

	userID := insertUser(user)

	res := response{
		ID:      userID,
		Message: "User create successfully",
	}
	json.NewEncoder(w).Encode(res)
}

func insertUser(user models.User) int64 {
	db := createConnection()
	defer db.Close()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user.Password = string(hashedPassword)
	sqlStatement := `INSERT INTO users(first_name, last_name, email, password, created_at) 
	VALUES ($1, $2, $3, $4, Now() ) RETURNING id`

	var id int64

	err = db.QueryRow(sqlStatement, user.First_name, user.Last_name, user.Email, user.Password).Scan(&id)
	if err != nil {
		log.Fatalf("Unable to execute the query %v", err)
	}

	return id

}

var sampleSecretKey = []byte("SecretYouShouldHide")

func UserLogin(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Fatalf("Unable to decode the request body %v", err)
	}

	checkEmail := checkEmail(req.Email)
	if checkEmail == false {
		log.Fatalf("In database we dont have user with this email.")
	}
	fmt.Println("email ispravan")

	/*user,err := getUserByEmail(req.Email)
	if err!=nil{
		log.Fatalf("user dont exist %v", err)
	}

	if !validPassword(req.Password, user){
		log.Fatalf("In database we dont have user with this password.")
	}*/
 

	tokenString, err := createJWT(req.Email, req.Password)
	if err != nil {
		log.Fatalf("Unable to create jwt %v", err)
	}
	fmt.Println(tokenString)

}

// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOiIyMDIzLTA1LTE2VDExOjUzOjI0LjAyMDQ1NTYyNSswMjowMCIsInVzZXIiOiJzb2ZpamFAZ21haWwuY29tIn0.hHYX9xD7kT1ngxzJFEIeNHDbjWdEH7NsCgjOJB8GNx8

func validPassword(password string, user models.User) bool {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))==nil
}

func checkEmail(email string) bool {
	db := createConnection()
	defer db.Close()

	sqlStatement := `SELECT email FROM users WHERE email=$1`
	row:= db.QueryRow(sqlStatement, email)
	
	switch err := row.Scan(&email); err{
	case sql.ErrNoRows:
		fmt.Println("No rows were returned")
		return false
	default:
		return true
	}
	
}

func createJWT(email string, password string) (string, error) {
	claims := &jwt.MapClaims{
		"expiresAt": 15000,
		"user": email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	

	return token.SignedString([]byte(sampleSecretKey))
}

func validateJWT(tokenString string) (*jwt.Token, error) {

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected singing method %v", token.Header["alg"])
		}
		return []byte(sampleSecretKey), nil
	})
}

func WithJWTAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("x-jwt-token")

		token, err := validateJWT(tokenString)
		if err != nil {
			log.Fatalf("token invalid %v", err)
		}
		if !token.Valid{
			log.Fatalf("token invalid %v", err)
		}

		params := mux.Vars(r)
		userid, err := strconv.Atoi(params["id"])
		if err!=nil {
			log.Fatalf("Unable to convert string into int %v", err)
		}
		user, err := getUserByID(int64(userid))

		claims := token.Claims.(jwt.MapClaims)
		if user.Email != claims["user"] {
			log.Fatalf("unable user %v", err)
		}

		handlerFunc(w, r)

	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert strig into int %v", err)
	}

	var user models.User

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatalf("Unable to decode the request %v", err)
	}

	updatedRow := updateUser(int64(id), user)
	msg := fmt.Sprintf("User updated successfully  %v", updatedRow)
	res := response{
		ID:      int64(id),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}

func updateUser(id int64, user models.User) int64 {
	db := createConnection()
	defer db.Close()
	sqlStatement := `UPDATE users SET first_name=$2, last_name=$3 WHERE id=$1`

	res, err := db.Exec(sqlStatement, id, user.First_name, user.Last_name)
	if err != nil {
		log.Fatalf("Unable to execute the query %v", err)
	}

	rowAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows %v", err)
	}
	return rowAffected
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert string into int %v", err)
	}

	user, err := getUserByID(int64(id))
	if err != nil {
		log.Fatalf("Unable to get user %v", err)
	}

	json.NewEncoder(w).Encode(user)
}

func getUserByID(id int64) (models.User, error) {
	db := createConnection()
	defer db.Close()
	sqlStatement := `SELECT * FROM users WHERE id=$1`

	var user models.User

	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&user.Id, &user.First_name, &user.Last_name, &user.Email, &user.Password, &user.Created_at)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Now rows were returned")
		return user, nil
	case nil:
		return user, nil
	default:
		log.Fatalf("Unable to scan the row %v", err)
	}
	return user, err
}

func GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	email := params["email"]

	user, err := getUserByEmail(email)
	if err != nil {
		log.Fatalf("Unable to get user %v", err)
	}

	json.NewEncoder(w).Encode(user)
}

func getUserByEmail(email string) (models.User, error) {
	db := createConnection()
	defer db.Close()
	sqlStatement := `SELECT * FROM users WHERE email=$1`

	var user models.User

	row := db.QueryRow(sqlStatement, email)
	err := row.Scan(&user.Id, &user.First_name, &user.Last_name, &user.Email, &user.Password, &user.Created_at)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Now rows were returned")
		return user, nil
	case nil:
		return user, nil
	default:
		log.Fatalf("Unable to scan the row %v", err)
	}
	return user, err
}
