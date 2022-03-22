package main

import (
	"log"
	"net/http"
	"os"

	"github.com/andrelrg/go-greet/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")

	s := server.Server{}

	s.Initialize()

	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(port, s.Router))
}
