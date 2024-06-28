package endpoints

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	initializers "github.com/SHRYNSH-NETAM/Go-Backend/Initializers"
)

type User struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Pass string `json:"pass"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var temp User
	if err := json.NewDecoder(r.Body).Decode(&temp); err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return 
	}

	alreadyexists := initializers.FindData(initializers.User{Email: temp.Email});
	if(alreadyexists.Username!=""){
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(User{Username: alreadyexists.Username})
		return
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

	alreadyexists := initializers.FindData(initializers.User{Email: temp.Email});
	if(alreadyexists.Pass==temp.Pass){
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(User{Username: alreadyexists.Username, Email: alreadyexists.Email})
		return
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(User{})
}

func SignupHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var temp User
	if err := json.NewDecoder(r.Body).Decode(&temp); err != nil {
		log.Fatal(err)
		fmt.Fprintf(w,"error")
		return
	}

	alreadyexists := initializers.FindData(initializers.User{Email: temp.Email});
	if(alreadyexists.Email==temp.Email){
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(User{})
		return
	}

	initializers.AddData(initializers.User{Username: temp.Username, Email: temp.Email, Pass: temp.Pass})
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(User{})
}

func Forgethandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var temp User
	if err := json.NewDecoder(r.Body).Decode(&temp); err != nil {
		log.Fatal(err)
		fmt.Fprintf(w,"error")
		return
	}

	alreadyexists := initializers.FindData(initializers.User{Email: temp.Email});
	if(alreadyexists.Email==temp.Email){
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(User{Pass: alreadyexists.Pass})
		return
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(User{})
}