package main

import (
	"net/http"
	"fmt"
	"log"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type User struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Pass string `json:"pass"`
}

var Users []User = []User{{Username: "Shreyansh Netam",Email: "netams2000@gmail.com",Pass: "Temp@2000"}}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/",HomeHandler).Methods("POST")
	r.HandleFunc("/login",LoginHandler).Methods("POST")
	r.HandleFunc("/signup",SignupHandler).Methods("POST")
	// r.HandleFunc("/forget",Forgethandler).Methods("GET")
	handler := cors.Default().Handler(r)
	fmt.Printf("Starting the server at 8000\n")
	log.Fatal(http.ListenAndServe(":8000",handler))
}

func HomeHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var temp User
	if err := json.NewDecoder(r.Body).Decode(&temp); err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return 
	}
	for _,val := range Users {
		if val.Email==temp.Email {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(User{Username: val.Username,Email: val.Email})
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func LoginHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var temp User
	if err := json.NewDecoder(r.Body).Decode(&temp); err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return 
	}
	for _,val := range Users {
		if val.Email==temp.Email && val.Pass==temp.Pass {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(User{Username: val.Username,Email: val.Email})
			return
		}else if val.Email==temp.Email {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(User{Email: "netams2000@gmail.com"})
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(User{Email: "netams2000@gmail.com"})
	// fmt.Fprintf(w,"Not_Found")
}

func SignupHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/plain")
	var temp User
	if err := json.NewDecoder(r.Body).Decode(&temp); err != nil {
		log.Fatal(err)
		fmt.Fprintf(w,"error")
		return
	}
	for _,val := range Users {
		if val.Email==temp.Email {
			w.WriteHeader(http.StatusConflict)
			fmt.Fprintf(w,"Already_Exists")
			return
		}
	}
	Users = append(Users, temp)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w,"Success")
}