package main

import (
	"log"
	"net/http"

	stub_contacts "github.com/wshaman/contacts-stub"
	"github.com/wshaman/course-rest/handlers/user"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v0/user", user.List).Methods(http.MethodGet)
	router.HandleFunc("/api/v0/user", user.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/v0/user/{id}", user.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/api/v0/user/{id}", user.View).Methods(http.MethodGet)
	router.HandleFunc("/api/v0/user/{id}", user.Update).Methods(http.MethodPut)
	router.Path("/api/v0/search").Queries("first_name", "{first_name}").HandlerFunc(user.SearchFirstName).Name("SearchFirstName").Methods(http.MethodGet)
	log.Println("Populating data")
	p := stub_contacts.LoadPersistent()
	if err := stub_contacts.Populate(p); err != nil {
		log.Fatal(err)
	}
	log.Println("Starting API server on 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
