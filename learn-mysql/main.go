// gorilla mux : for creating routes and http handlers.
// go get github.com/go-sql-driver/mysql.
// gorm  : an orm tool for mysql.
// go get github.com/jinzhu/gorm or we can use go get -u github.com/jinzhu/gorm.
// mysql : is the driver.
// go get github.com/go-sql-driver/mysql.
// before start, we have to create database manually.
// mysql -u root -p (Hallo123$).
// create database chat;
// show databases (to see all databases on mysql).
// create file name main.go.
// we want to connect our application to our database.
// db, err = gorm.Open(“mysql”, “user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8&parseTime=True”).
// encoding/json : this package contains methods that are used to convert Go types to JSON and -
// vice-versa (This conversion is called as encode/decode in Go, serialization/de-serialization or marshall/unmarshall in other languages).
// fmt - This package implements formatted I/O functions similar to scanf and printf in C.
// log - Has methods for formatting and printing log messages.
// net/http - Contains methods for performing operations over HTTP. It provides HTTP server and client implementations and has abstractions for HTTP request, response, headers, etc.
// time - Provides methods for handling(storing/displaying/manipulating) time values.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

// create a simple user struct that features id, username, phone, password.
// default table name is User.
type User struct {
	// gorm mode.
	Id       int    `json:"id"`
	Username string `json:"username"`
	Phone    int    `json:"phone"`
	Password string `json:"password"`
}

// homepage function.
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Homepage!")
	fmt.Println("Endpoint Hit: HomePage!")
}

// create web server for handling HTTP requests.
// creating a new function named as handleRequests()
// and under this we will create function a new instance of a mux router.
func handleRequests() {
	log.Println("Starting development server at http://127.0.0.1:8080")
	log.Println("Quit the server with control-C")
	// creates a new instance of a mux router.
	myRouter := mux.NewRouter().StrictSlash(true)
	// for home page.
	myRouter.HandleFunc("/", homePage)
	// create user.
	myRouter.HandleFunc("/create", createNewUser).Methods("POST")
	// get all user.
	myRouter.HandleFunc("/show", getAllUser).Methods("GET")
	// get user by id.
	myRouter.HandleFunc("/show/{id}", getUserByID).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

// crud operation.
// create user.
func createNewUser(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user User
	json.Unmarshal(reqBody, &user)
	db.Create(&user)
	fmt.Println("Endpoint Hit: Creating New Users")
	json.NewEncoder(w).Encode(user)
}

// get all user.
func getAllUser(w http.ResponseWriter, r *http.Request) {
	users := []User{}
	db.Find(&users)
	fmt.Println("Endpoint Hit: getAllUser")
	json.NewEncoder(w).Encode(users)
}

// get user by id.
func getUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	users := []User{}
	db.Find(&users)
	for _, user := range users {
		// string to int.
		s, err := strconv.Atoi(key)
		if err == nil {
			if user.Id == s {
				fmt.Println(user)
				fmt.Println("Endpoint Hit: User No:", key)
				json.NewEncoder(w).Encode(user)
			}
		}
	}
}

func main() {
	// define username and password for mysql.
	db, err = gorm.Open("mysql", "root:Hallo123$@tcp(127.0.0.1:3306)/chat?charset=utf8&parseTime=True")
	// note: we are using = to assign the global var.
	// instead of := which would assign it only in this function.

	if err != nil {
		log.Println("Connection Failed to Open!")
	} else {
		log.Println("Connection Established!")
	}
	// test the file.
	// go build.
	// go run main.go.
	// result: 2020/07/19 14:49:01 Connection Established!.
	// go to struct.
	// after creating struct, it is time to migrate our schema.
	db.AutoMigrate(&User{})
	// AutoMigrate only create tables, missing columns and missing indexes.
	// won't change existing column's type or deleted un-used columns to protect our data.

	// call function name handleRequests().
	handleRequests()
	// now run the code by typing go run main.go.
	// result:
	// 2020/07/19 15:28:44 Connection Established!.
	// 2020/07/19 15:28:44 Starting development server at http://127.0.0.1:8080.
	// 2020/07/19 15:28:44 Quit the server with control-C.
	// open up "http://localhost:8080" in browser and we can see "Welcome to Homepage!".
}
