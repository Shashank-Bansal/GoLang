package server

import (
	"github.com/go-chi/cors"
	"log"
	"net/http"
)

func StartSever(port string) {
	router := GetInstance().Route
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*", "https://*", "http://localhost:3000/*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowedHeaders: []string{"*"},
	}))

	router.Use(check);

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

func check(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Incoming - METHOD: [%s] - URL: [%s] - IP: [%s]", r.Method, r.URL, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}