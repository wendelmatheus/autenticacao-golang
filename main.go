package main

import (
	"go-jwt/config"
	"go-jwt/database"
	"go-jwt/handlers"
	"go-jwt/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadEnv()
	database.Connect()

	r := mux.NewRouter()

	r.HandleFunc("/register", handlers.Register).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")

	r.HandleFunc("/users", middleware.AuthMiddleware(handlers.GetUsers)).Methods("GET")
	r.HandleFunc("/users/{id}", middleware.AuthMiddleware(handlers.GetUser)).Methods("GET")
	r.HandleFunc("/users/{id}", middleware.AuthMiddleware(handlers.UpdateUser)).Methods("PUT")
	r.HandleFunc("/users/{id}", middleware.AuthMiddleware(handlers.DeleteUser)).Methods("DELETE")

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
