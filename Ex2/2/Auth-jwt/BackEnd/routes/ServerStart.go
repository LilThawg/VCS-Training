package routes

import (
	"auth-jwt/controllers"
	"auth-jwt/middlewares"
	"log"
	"net/http"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//start the server
func ServerStart() {
	router := mux.NewRouter()
	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "token"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	
	router.HandleFunc("/signup", controllers.SignUp).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/logout", controllers.LogOut).Methods("POST")

	router.HandleFunc("/getalluser", middlewares.IsAuthorized(controllers.GetAllUsers)).Methods("GET")
	router.HandleFunc("/adduser", middlewares.IsAuthorized(controllers.AddUser)).Methods("POST")
	router.HandleFunc("/deleteuser", middlewares.IsAuthorized(controllers.DeleteUser)).Methods("DELETE")

	http.ListenAndServe(":4000",handlers.CORS(header, methods, origins)(router));
	log.Println("API is running!")
}

