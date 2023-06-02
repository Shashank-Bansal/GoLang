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
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	PersonalDetails PersonalDetails `bson:"personal_details" json:"personal_details"`
	Residence       Residence       `bson:"residence" json:"residence"`
	Pwd_hash        string          `bson:"pwd_hash" json:"pwd_hash"`
}

type Residence struct {
	Address string `bson:"address" json:"address"`
	Pincode string `bson:"pincode" json:"pincode"`
	Country string `bson:"country" json:"country"`
	State   string `bson:"state" json:"state"`
	City    string `bson:"city" json:"city"`
}

type PersonalDetails struct {
	Name       string `bson:"name" json:"name"`
	Email      string `bson:"email" json:"email"`
	MobileNo   string `bson:"mobile_no" json:"mobile_no"`
	DOB        string `bson:"dob" json:"dob"`
	Gender     string `bson:"gender" json:"gender"`
	BloodGrp   string `bson:"blood_grp" json:"blood_grp"`
	AadhaarUID string `bson:"aadhaar_uid" json:"aadhaar_uid"`
}


func SetupRoutes() {
	log.Println("Setting up routes...")
	userIndexes()
	loginRoute := chi.NewRouter()
	loginRoute.Post("/admin", checkUser)
	loginRoute.Get("/", getUser)
	loginRoute.Patch("/", updateUser)
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

func updateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var updatedFields map[string]interface{}
	err = json.Unmarshal(body, &updatedFields)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	database := server.GetMongoInstance().Database
	userCollection := database.Collection("users")

	var email = updatedFields["email"]

	log.Println(email)
	result := userCollection.FindOneAndUpdate(context.TODO(), bson.M{"personal_details.email": email}, bson.M{"$set": updatedFields})
	if result.Err() != nil {
		log.Println(result.Err())
		log.Println(updatedFields)
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user updated successfully"))
}

func checkUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	var loginUser struct {
		Email    string `json:"email"`
		Pwd_hash string `json:"pwd"`
	}

	err = json.Unmarshal(body, &loginUser)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	log.Println(loginUser)
	database := server.GetMongoInstance().Database
	userCollection := database.Collection("users")

	var user User
	// err = userCollection.FindOne(context.TODO(), bson.M{"personal_details.email": loginUser.Email, "pwd_hash": loginUser.Pwd_hash}).Decode(&user)
	err = userCollection.FindOne(context.TODO(), bson.M{"personal_details.email": loginUser.Email).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Either email or password is incorrect", http.StatusBadRequest)
		} else {
			http.Error(w, "Failed to find user", http.StatusInternalServerError)
		}
		return
	}

	err = bcrypt.CompareHashAndPassword(user.Pwd_hash, loginUser.pwd_hash)
	if err != nil {
		http.Error(w, "Either email or password is incorrect", http.StatusBadRequest)
		return
	}
	userJSON, _ := json.Marshal(user)
	log.Println(userJSON, user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(userJSON))
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

	pass := "Gurukuk1"
	pass, err = bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
    if err != nil {
        log.Println(err)
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
    }
	
	newUser.Pwd_hash = string(pass)

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
