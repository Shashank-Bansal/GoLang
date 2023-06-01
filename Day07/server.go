package main

import (
	"log"
	"os"
	"sync"
	"time"

	server "github.com/Shashank-Bansal/GoLang/Day07/Server"
	controller "github.com/Shashank-Bansal/GoLang/Day07/Server/Controller"
)

func main() {
	server.LoadEnv()

	_ = server.GetMongoInstance().Database
	defer server.CloseDatabase()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Please specify the port for the server")
	}

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Add(1)
	go server.StartSever(port)
	time.Sleep(time.Second)
	go controller.SetupRoutes()

	wg.Wait()
}

// collection := database.Collection("Test")
// document := []interface{}{
// 	bson.D{
// 		{Key: "User", Value: "Aman"},
// 		{Key: "Password", Value: "password123456"},
// 	},
// 	bson.D{
// 		{Key: "User", Value: "Ankit"},
// 		{Key: "Password", Value: "password123456"},
// 	},
// }

// _, err := collection.InsertMany(context.TODO(), document)
// if err != nil {
// 	log.Fatal(err)
// }

// _, err := collection.FindAll(context.TODO(), document)
// if err != nil {
// 	log.Fatal(err)
// }
