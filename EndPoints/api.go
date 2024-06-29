package endpoints

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	initializers "github.com/SHRYNSH-NETAM/Go-Backend/Initializers"
	"github.com/SHRYNSH-NETAM/Go-Backend/utils"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Pass string `json:"pass"`
}

type res struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var temp res
	if err := json.NewDecoder(r.Body).Decode(&temp); err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return 
	}

	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(User{})
		return
	}

	tokenString = tokenString[len("Bearer "):]
	
	err := utils.VerifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(User{})
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
	err := bcrypt.CompareHashAndPassword([]byte(alreadyexists.Pass),[]byte(temp.Pass))
	if(err==nil){
		tokenString, error := utils.CreateToken(temp.Email)
		if error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(res{})
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res{Username: alreadyexists.Username, Email: alreadyexists.Email, Token: tokenString})
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

	hashpass, _ := bcrypt.GenerateFromPassword([]byte(temp.Pass), 14)

	initializers.AddData(initializers.User{Username: temp.Username, Email: temp.Email, Pass: string(hashpass)})
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