// package main

// import (
// 	// "fmt"
// 	"log"
// 	"net/http"
// 	"os"

// 	// login "github.com/Shashank-Bansal/GoLang/Day06/server/login"
// 	"github.com/go-chi/chi"
// 	"github.com/go-chi/cors"
// 	"encoding/json"
// 	"github.com/joho/godotenv"
// )

// type credentials struct {
// 	user     string `json:"user"`
// 	password string `json:"password"`
// }

// func main() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		log.Fatal("Please specify the port for the server")
// 	}

// 	log.Println("Starting server on port", port)

// 	router := chi.NewRouter()
// 	router.Use(cors.Handler(cors.Options{
// 		AllowedOrigins: []string{"http://*", "https://*"},
// 		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
// 		AllowedHeaders: []string{"*"},
// 	}))

// 	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte("Hello, world!"))
// 	})

// 	// router.Mount("/login", login.LoginRoute)

// 	LoginRoute := chi.NewRouter()
// 	LoginRoute.Post("/", func(w http.ResponseWriter, r *http.Request) {
// 		// log.Println(r)
// 		// user := r.FormValue("user")
// 		// password := r.FormValue("password")
// 		var cred credentials
// 		err := json.NewDecoder(r.Body).Decode(&cred)
// 		if err != nil {
// 			http.Error(w, "Invalid JSON", http.StatusBadRequest)
// 			log.Fatal(err)
// 			return
// 		}
// 		user := cred.user
// 		password := cred.password

// 		// r.ParseForm()

// 		// user := r.Form.Get("user")
// 		// password := r.Form.Get("password")

// 		log.Println("User: ", user, "Password: ", password, "trying to login")

// 		response := "User: " + user + ", Password: " + password
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte(response))
// 	})

// 	router.Mount("/login", LoginRoute)

// 	server := &http.Server{
// 		Handler: router,
// 		Addr:    ":" + port,
// 	}

// 	log.Println(server.Addr)

// 	errServer := server.ListenAndServe()
// 	if errServer != nil {
// 		log.Fatal("Error starting server")
// 	}

// }



package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

type credentials struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Please specify the port for the server")
	}

	log.Println("Starting server on port", port)

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*", "https://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowedHeaders: []string{"*"},
	}))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, world!"))
	})

	LoginRoute := chi.NewRouter()
	LoginRoute.Get("/", func(w http.ResponseWriter, r *http.Request) {
		var cred credentials
		err := json.NewDecoder(r.Body).Decode(&cred)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			log.Println("Error decoding JSON:", err)
			return
		}
		user := cred.User
		password := cred.Password

		log.Println("User:", user, "Password:", password, "trying to login")

		response := "User: " + user + ", Password: " + password
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	})

	router.Mount("/login", LoginRoute)

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