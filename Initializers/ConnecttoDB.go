package initializers

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Username string `json:"username,omitempty" bson:"username,omitempty"`
	Email string `json:"email,omitempty" bson:"email,omitempty"`
	Pass string `json:"pass,omitempty" bson:"pass,omitempty"` 
}

var client *mongo.Client

func ConnecttoDB(){
	fmt.Println("Connecting... to DB")

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. " +
			"See: " +
			"www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Failed to connect to MongoDB: ", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB: ", err)
	}

	fmt.Println("Successfully connected and pinged.")
}

func AddData(loginData User){
	coll := client.Database("LoginDB").Collection("Users")
	result, err := coll.InsertOne(context.TODO(), loginData)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func FindData(loginData User) User{
	coll := client.Database("LoginDB").Collection("Users")
	var result User
	if err := coll.FindOne(context.TODO(), loginData).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return User{}
		}
		panic(err)
	}
	return result
}