package controllers

import (
	"01/models"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserController which has methods attached to it like: GetUser, CreateUser, DeleteUser
type UserController struct {
	Client *mongo.Client
}

var uc UserController

// NewUserController func create a new pointer to UserController variable
func NewUserController(client *mongo.Client) *UserController {

	return &UserController{
		Client: client,
	}
}

// GetUser func is attached to UserController which gets user
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var user models.User

	id := p.ByName("id")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("cannot get object id:", err)
		return
	}

	err = uc.Client.Database("go-web-dev").Collection("users").FindOne(ctx, bson.M{"_id": ObjectID}).Decode(&user)
	if err != nil {
		fmt.Println("cannot find or decode:", err)
		return
	}

	jsonStringByte, _ := json.Marshal(user)

	fmt.Println(string(jsonStringByte))
}

// CreateUser func creates new user from sent JSON user data converts to user model struct and again back to JSON and sends it.
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	user := models.User{}

	user.ID = primitive.NewObjectID()

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("Cannot decode given JSON to user struct:", err)
		return
	}

	collection := uc.Client.Database("go-web-dev").Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, user)
	if err != nil {
		fmt.Println("cannot insert:", err)
		return
	}

	fmt.Println("inserted doc:", res)

	userJSONbyte, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Cannot marshal given user struct to JSON:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userJSONbyte)
}

// DeleteUser func deletes given user
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := uc.Client.Database("go-web-dev").Collection("users").DeleteOne(ctx, bson.M{"_id": ObjectID})
	if err != nil {
		fmt.Println("not deleted:", err)
		return
	}

	fmt.Println(res.DeletedCount)
}
