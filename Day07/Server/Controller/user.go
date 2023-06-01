package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"

	server "github.com/Shashank-Bansal/GoLang/Day07/Server"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	PersonalDetails PersonalDetails `bson:"personal_details"`
	Residence       Residence       `bson:"residence"`
	Pwd_hash        string          `bson:"pwd_hash"`
}

type Residence struct {
	Address string `bson:"address"`
	Pincode string `bson:"pincode"`
	Country string `bson:"country"`
	State   string `bson:"state"`
	City    string `bson:"city"`
}

type PersonalDetails struct {
	Name       string `bson:"name"`
	Email      string `bson:"email"`
	MobileNo   string `bson:"mobile_no"`
	DOB        string `bson:"dob"`
	Gender     string `bson:"gender"`
	BloodGrp   string `bson:"blood_grp"`
	AadhaarUID string `bson:"aadhaar_uid"`
}

func SetupRoutes() {
	log.Println("Setting up routes...")
	userIndexes()
	loginRoute := chi.NewRouter()
	loginRoute.Get("/", getUser)
	loginRoute.Post("/", postUser)
	loginRoute.Delete("/", deleteUser)

	router := server.GetInstance().Route
	router.Mount("/login", loginRoute)
}

func userIndexes() {
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"personal_details.email": 1},
		Options: options.Index().SetUnique(true),
	}

	database := server.GetMongoInstance().Database
	userCollection := database.Collection("users")
	userCollection.Indexes().CreateOne(context.TODO(), indexModel)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	database := server.GetMongoInstance().Database
	userCollection := database.Collection("users")

	cursor, err := userCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println("Failed to retrieve users:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	var users []User // Assuming User is the struct representing a user document

	for cursor.Next(context.TODO()) {
		var user User
		if err := cursor.Decode(&user); err != nil {
			log.Println("Failed to decode user:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	usersJSON, err := json.Marshal(users)
	if err != nil {
		log.Println("Failed to marshal users to JSON:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(usersJSON)
}

func postUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	var newUser User
	err = json.Unmarshal(body, &newUser)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	database := server.GetMongoInstance().Database
	userCollection := database.Collection("users")

	_, err = userCollection.InsertOne(context.TODO(), newUser)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			http.Error(w, "User with the same email already exists", http.StatusBadRequest)
		} else {
			http.Error(w, "Failed to insert user", http.StatusInternalServerError)
		}
		return
	}

	response := fmt.Sprintf("User '%s' created successfully", newUser.PersonalDetails.Name)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	var userEmail struct {
		Email string `json:"email"`
	}
	json.Unmarshal(body, &userEmail)

	database := server.GetMongoInstance().Database
	userCollection := database.Collection("users")
	result := userCollection.FindOneAndDelete(context.TODO(), bson.M{"personal_details.email": userEmail.Email})
	if result.Err() != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	response := fmt.Sprintf("User with Email '%s' deleted successfully", userEmail.Email)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}