package server

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Credentials struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type Router struct {
	Route *chi.Mux
}

type MongoDatabase struct {
	Database *mongo.Database
}

var (
	instance      *Router
	once          sync.Once
	mongoOnce     sync.Once
	mongoInstance *MongoDatabase
	mongoClient   *mongo.Client
)

func GetInstance() *Router {
	once.Do(func() {
		instance = &Router{Route: chi.NewRouter()}
	})

	return instance
}

func GetMongoInstance() *MongoDatabase {
	mongoOnce.Do(func() {
		opts := options.Client().ApplyURI(os.Getenv("MONGO_URL"))
		client, err := mongo.Connect(context.TODO(), opts)
		if err != nil {
			log.Println("Failed to connect to MongoDB:", err)
			return
		}

		mongoClient = client
		mongoInstance = &MongoDatabase{Database: client.Database("Gurukul")}
		log.Println("Connected to MongoDB!")
	})

	return mongoInstance
}

func CloseDatabase() {
	if err := mongoClient.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
