// gorilla mux: fpr creating routes and http handlers.package main
// go get github.com/gorilla/mux.
// gorm: an orm tool for mysql.
// go get github.com/jinzhu/gorm.
// mysql: the mysql driver.
// go get github.com/go-sql-driver/mysql.
// math/rand: to make random id when create (must string type not int).

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// lihat apakah db error (define db error).
var err error

// struct (class).
type User struct {
	ID       string `json:"id"`
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

// homepage function on browser (go run main.go after that open browser and write localhost:8080, result: welcome).
func homePage(w http.ResponseWriter, r *http.Request) {
	// Fprintf to show output in browser not on terminal.
	fmt.Fprintf(w, "welcome")
}

// crud test using postman.
// create user.
func createUser(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Create User!")
	// variable to get payloads (_ is to catch error but we don't use).
	payloads, _ := ioutil.ReadAll(r.Body)

	// user variable and User struct.
	var user User
	// casting payload to User.
	json.Unmarshal(payloads, &user)

	// create random id.
	user.ID = strconv.Itoa(rand.Intn(1000000))
	// user table.
	db.Create(&user)

	// data (result from struct).
	res := Result{Code: 200, Data: user, Message: "Success Create User!"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// response.
	w.Header().Set("Content-Type", "application/json")
	// http status.
	w.WriteHeader(http.StatusOK)
	// res.body.
	w.Write(result)
}

// get all user (show all user).
func getAllUser(w http.ResponseWriter, r *http.Request) {
	// make table.
	user := []User{}
	// get data.
	db.Find(&user)
	// result.
	res := Result{Code: 200, Data: user, Message: "Success Get All User!"}
	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// response.
	w.Header().Set("Content-Type", "application/json")
	// http status.
	w.WriteHeader(http.StatusOK)
	// res.body.
	w.Write(results)
}

// get user by id.
func getUserById(w http.ResponseWriter, r *http.Request) {
	// get id.
	vars := mux.Vars(r)
	userID := vars["id"]

	var user User
	// get data.
	db.First(&user, userID)
	// result.
	res := Result{Code: 200, Data: user, Message: "Success Get User By ID!"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// response.
	w.Header().Set("Content-Type", "application/json")
	// http status.
	w.WriteHeader(http.StatusOK)
	// res.body.
	w.Write(result)
}

// update (edit) user by id. it was same with create user.
func updateUserById(w http.ResponseWriter, r *http.Request) {
	// get id.
	vars := mux.Vars(r)
	userID := vars["id"]

	// variable to get payloads (_ is to catch error but we don't use).
	payloads, _ := ioutil.ReadAll(r.Body)

	// userUpdate variable and User struct.
	var userUpdate User
	// casting payload to User.
	json.Unmarshal(payloads, &userUpdate)

	// existing user data.
	var user User
	// get userID from User table from the existing data.
	db.First(&user, userID)
	// update db.
	db.Model(&user).Update(userUpdate)

	// data (result from struct).
	res := Result{Code: 200, Data: user, Message: "Success Update User!"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// response.
	w.Header().Set("Content-Type", "application/json")
	// http status.
	w.WriteHeader(http.StatusOK)
	// res.body.
	w.Write(result)
}

// delete user by id.
func deleteUserById(w http.ResponseWriter, r *http.Request) {
	// get id.
	vars := mux.Vars(r)
	userID := vars["id"]

	// get user from table.
	var user User

	// delete from database.
	db.First(&user, userID)
	db.Delete(&user)

	// data (result from struct).
	// we don't need send data because it will destroy from database.
	res := Result{Code: 200, Message: "Success Delete User By ID!"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// response.
	w.Header().Set("Content-Type", "application/json")
	// http status.
	w.WriteHeader(http.StatusOK)
	// res.body.
	w.Write(result)
}

// in nodejs we call it routes.
func handleRequests() {
	// log.Println (console.log on javascript) to see if the routes running.
	log.Println("Start the development server at localhost:8080!")

	// package mux for library routes (in nodejs we use const Router = req(router)).
	Router := mux.NewRouter().StrictSlash(true)
	// homePage.
	Router.HandleFunc("/", homePage)
	// create user.
	Router.HandleFunc("/user/create", createUser).Methods("POST")
	// get user.
	Router.HandleFunc("/user/show", getAllUser).Methods("GET")
	// get user by id (in nodejs we defined user at app.js and endpoint on routes).
	Router.HandleFunc("/user/show/{id}", getUserById).Methods("GET")
	// update user by id.
	Router.HandleFunc("/user/update/{id}", updateUserById).Methods("PUT")
	// delete user by id.
	Router.HandleFunc("/user/delete/{id}", deleteUserById).Methods("DELETE")

	// port.
	log.Fatal(http.ListenAndServe(":8080", Router))
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
