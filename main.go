package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Mannan-Ali/RSS-Aggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	//* means pointer and DB will hold all the queris we created using  sqlc
	DB *database.Queries
}

func main() {
	fmt.Println("Hey first a project in GO")

	godotenv.Load() //to load the env file we use this function

	//to use env go has the follwing functoin
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the env")
	}
	router := chi.NewRouter() // this creates a router

	//connecting to database
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the env")
	}

	//establising connection with postgress
	cons, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to database:", err)
	}

	//this apiconfig now can be passed into different function and they will have access to our database
	//apiconfig allows us to use the functions we created using sql queires
	apiCfg := apiConfig{
		DB: database.New(cons),
	}

	//setting up cors
	//this setting allows any request from any website
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	//if handlefunc is used any kind of req can be made get,put,post but it is bad practise
	// v1Router.HandleFunc("/ready", handlerReadiness)

	// reason mount is after get here : as we arefully preparing the sub-router first, and then attaching the finished,
	//configured component to the main router.

	//telling the v1 router that if /healthz is called with get then call this function
	v1Router.Get("/healtz", handlerReadiness)

	//for error go on this route
	v1Router.Get("/err", handlerErr)
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.handlerGetUserByAPIKey)
	//mount connects the main router with v1router so if request with /v1 comes it is handed to v1router
	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	log.Printf("Server starting at port %v", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("PORT:", portString)
}
