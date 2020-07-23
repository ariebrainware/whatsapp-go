// gorilla mux: fpr creating routes and http handlers.package main
// go get github.com/gorilla/mux.
// gorm: an orm tool for mysql.
// go get github.com/jinzhu/gorm.
// mysql: the mysql driver.
// go get github.com/go-sql-driver/mysql.

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// lihat apakah db error (define db error).
var err error

// struct (class).
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// result is an array of user (for migration).
type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func main() {
	// db connection.
	// root is username, Hallo123$ is a password, @api is a database name.
	db, err = gorm.Open("mysql", "root:Hallo123$@/api?charset=utf8&parseTime=True")

	// check connection null (nil).
	if err != nil {
		log.Println("Connection Failed!", err)
	} else {
		// connection success.
		log.Println("Connection established")
	}

	// auto migrate User table (same like nodejs we run sequelize db:migrate)
	db.AutoMigrate(&User{})
	// check database: use api, show tables, desc users, select *from users (it will be empty because have not create data.).

	// call handle request function (route / routing).
	handleRequests()
}

// in nodejs we call it routes.
func handleRequests() {
	// log.Println (console.log on javascript) to see if the routes running.
	log.Println("Start the development server at localhost:8080!")

	// package mux for library routes (in nodejs we use const Router = req(router)).
	Router := mux.NewRouter().StrictSlash(true)
	// homePage.
	Router.HandleFunc("/", homePage)
	Router.HandleFunc("/user/create", createUser).Methods("POST")
	// port.
	log.Fatal(http.ListenAndServe(":8080", Router))
}

// homepage function on browser (go run main.go after that open browser and write localhost:8080, result: welcome).
func homePage(w http.ResponseWriter, r *http.Request) {
	// Fprintf to show output in browser not on terminal.
	fmt.Fprintf(w, "welcome")
}

// crud test using postman.
// create user.
func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create User!")
}
