// before start, we should be in this directory: GOPATH/src/github.com/<Github username>. 
// is just a function name but when used with package main, it serves as the application entry point.
package main

// fmt is golang package that implements formatted input and output.
// set up tbe HTTP server using gorilla mux
// gorilla mux is package that implements a request router and dispatcher for matching incoming request to their respective handle.
// go get -u github.com/gorilla/mux. 
// we use gorilla mux to help us make our fisrt home endpoint "/".
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

// create a dummy database.
// create a struct and a slice.
type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}
type allEvents []event

// we only use title and description fields for event struct.
var events = allEvents{
	{
	ID          : "1",
	Title       : "Introduction to Golang",
	Description : "Come to join for a chance to learn how golang work and get to eventually try it out",
	},
}

// for homeLink to show "Welcome Home" at vs code terminal.
// stuck to show on postman.
func homeLink(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

// create an event.
// user enters data which is in the form of http Request data.
func createEvent(w http.ResponseWriter, r *http.Request)  {
	var newEvent event
	// the request data is not is a human-readable format hence we use the package ioutil to convert it into a slice.
	// Convert r.Body into a readable format. r.Body like req.body on Node.js.
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(w, "Kindly enter data with the event and description only in order to update")
	}
	json.Unmarshal(reqBody, &newEvent)
	// Add the newly created event to the array of events.
	events = append(events, newEvent)
	w.WriteHeader(http.StatusCreated)
	// Return the newly created event
	json.NewEncoder(w).Encode(newEvent)
	// when test on postman: http Responsewith 201 Created Status Code.
}

// get one event by id.
// endpoint: /event/{id} and use get method.
func getOneEvent(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the url.
	eventID := mux.Vars(r)["id"]

	// Get the details from an existing event.
	// Use the blank identifier to avoid creating a value that will not be used.
	for _, singleEvent := range events {
		if singleEvent.ID == eventID {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

// get all events.
func getAllEvents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(events)
}

// update an event.
// endpoint: /events/{id} and use patch method.
func updateEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]
	var updatedEvent event

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	json.Unmarshal(reqBody, &updatedEvent)

	for i, singleEvent := range events {
		if singleEvent.ID == eventID {
			singleEvent.Title = updatedEvent.Title
			singleEvent.Description = updatedEvent.Description
			events = append(events[:i], singleEvent)
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

// Delete event.
// endpoint: /events/{id} and uses delete method.
func deleteEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	// Get the details from an existing event.
	// Use the blank identifier to avoid creating a value that will not be used.
	for i, singleEvent := range events {
		if singleEvent.ID == eventID {
			events = append(events[:i], events[i+1:]...)
			fmt.Fprintf(w, "The event with ID %v has been deleted successfully", eventID)
		}
	}
}

func main()  {
	// fmt.Println("Hello World!")
	// we created a function that will return the "welcome home."
	router := mux.NewRouter().StrictSlash(true)
	// this is when the "/" endpoint is hit.
	router.HandleFunc("/", homeLink)
	// result on terminal: Welcome Home!.
	router.HandleFunc("/event", createEvent).Methods("POST")
	router.HandleFunc("/events", getAllEvents).Methods("GET")
	router.HandleFunc("/events/{id}", getOneEvent).Methods("GET")
	router.HandleFunc("/events/{id}", updateEvent).Methods("PATCH")
	router.HandleFunc("/events/{id}", deleteEvent).Methods("DELETE")
	// we also created a server that runs on http://localhost:8080.
	log.Fatal(http.ListenAndServe(":8080", router))
}