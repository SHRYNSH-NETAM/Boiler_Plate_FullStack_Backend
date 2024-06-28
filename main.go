package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	endpoints "github.com/SHRYNSH-NETAM/Go-Backend/EndPoints"
	initializers "github.com/SHRYNSH-NETAM/Go-Backend/Initializers"
	// "github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	// "golang.org/x/crypto/bcrypt"
)

func init(){
	initializers.Initializers()
	initializers.ConnecttoDB()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/",endpoints.HomeHandler).Methods("POST")
	r.HandleFunc("/login",endpoints.LoginHandler).Methods("POST")
	r.HandleFunc("/signup",endpoints.SignupHandler).Methods("POST")
	r.HandleFunc("/forget",endpoints.Forgethandler).Methods("POST")
	handler := cors.Default().Handler(r)
	PORT := os.Getenv("PORT")
	fmt.Printf("Starting the server at %v\n", PORT)
	log.Fatal(http.ListenAndServe(":8000",handler))
}