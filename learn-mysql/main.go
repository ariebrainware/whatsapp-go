// gorilla mux for creating routes and http handlers.
// go get github.com/go-sql-driver/mysql.
// gorm an orm tool for mysql.
// go get github.com/jinzhu/gorm or we can use go get -u github.com/jinzhu/gorm.
// mysql is the driver.
// go get github.com/go-sql-driver/mysql.
// before start, we have to create database manually.
// mysql -u root -p (Hallo123$).
// create database chat;
// show databases (to see all databases on mysql).
// create file name main.go.
// we want to connect our application to our database.
// db, err = gorm.Open(“mysql”, “user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8&parseTime=True”).

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

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

// create a simple user struct that features id, username, phone, password.
type User struct {
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
	myRouter.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}