package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	endpoints "github.com/SHRYNSH-NETAM/Go-Backend/EndPoints"
	initializers "github.com/SHRYNSH-NETAM/Go-Backend/Initializers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func init(){
	initializers.Initializers()
	initializers.ConnecttoDB()
}

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})
	r := mux.NewRouter()
	handler := c.Handler(r)
	r.HandleFunc("/",endpoints.HomeHandler).Methods("POST")
	r.HandleFunc("/login",endpoints.LoginHandler).Methods("POST")
	r.HandleFunc("/signup",endpoints.SignupHandler).Methods("POST")
	r.HandleFunc("/forget",endpoints.Forgethandler).Methods("POST")
	PORT := os.Getenv("PORT")
	fmt.Printf("Starting the server at %v\n", PORT)
	log.Fatal(http.ListenAndServe(":8000",handler))
}