package main

import (
	"01/controllers"
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}
	fmt.Printf("%v\n%T\n", os.Getenv("MONGOKEY"), os.Getenv("MONGOKEY"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGOKEY")))
	if err != nil {
		fmt.Println("Cannot connect to mongoDB:", err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("Cannot ping:", err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	route := httprouter.New()
	uc := controllers.NewUserController(client)
	route.GET("/user/:id", uc.GetUser)
	route.POST("/user", uc.CreateUser)
	route.DELETE("/user/:id", uc.DeleteUser)

	http.ListenAndServe(":3000", route)
}
