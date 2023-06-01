package server

import (
	"github.com/go-chi/cors"
	"log"
	"net/http"
)

func StartSever(port string) {
	router := GetInstance().Route
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*", "https://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowedHeaders: []string{"*"},
	}))

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Println(server.Addr)

	errServer := server.ListenAndServe()
	if errServer != nil {
		log.Fatal("Error starting server:", errServer)
	}

}
