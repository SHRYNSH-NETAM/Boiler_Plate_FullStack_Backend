package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// var secretkey string =  os.Getenv("secretkey")
var secretKey = []byte("secret-key")


func CreateToken (email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
				jwt.MapClaims{
					"email": email,
					"exp": time.Now().Add(time.Hour * 24).Unix(),
				})

	tokenstring, err := token.SignedString(secretKey)
	if err!=nil{
		return "",err
	}
	return tokenstring, nil
}

func VerifyToken(tokenString string) error {

	
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	   return secretKey, nil
	})
   
	if err != nil {
	   return err
	}
   
	if !token.Valid {
	   return fmt.Errorf("invalid token")
	}
   
	return nil
 }