package main

import (
	"log"
	"net/http"

	"github.com/edwinguerrerotech/mux-movies/routes"
)

func main() {

	router := routes.MovieRoutes()

	http.Handle("/api", router)

	log.Println("Listening...")

	// For local testing, use:
	log.Fatal(http.ListenAndServe(":8081", router))

	/* Must use .env file instead of string literal for the PORT,
	else Heroku will have issues identifying the PORT. */
	// log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
