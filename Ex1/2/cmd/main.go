package main

import (
	"log"
	"net/http"

	"github.com/VCS-trainning/Ex1/2/crud/pkg/db"
	"github.com/VCS-trainning/Ex1/2/crud/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()
	
	router.HandleFunc("/users", h.AddUser).Methods(http.MethodPost)
	router.HandleFunc("/users", h.GetAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{name}", h.GetUserByName).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", h.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", h.DeleteUser).Methods(http.MethodDelete)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)

}
