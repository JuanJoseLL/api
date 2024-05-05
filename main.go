package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == ""{
		log.Fatal("PORT is not found in the envioroment")
	}

	router := chi.NewRouter()
	srv := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT"},
		AllowCredentials: false,
		AllowedHeaders: []string{"*"},
	}))

	v1Router := chi.NewRouter()

	v1Router.Get("/ready", handlerReadiness)
	v1Router.Get("/err", handleError)


	router.Mount("/v1", v1Router)


	log.Printf("Server starting on port: %v", portString)
	err := srv.ListenAndServe()

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("PORT: ", portString)
}