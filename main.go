package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	r := mux.NewRouter()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r.HandleFunc("/", TestHandler)
	//db := database.SetupDB()
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:" + os.Getenv("PORT"),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Listening on localhost:" + os.Getenv("PORT"))
	log.Fatal(srv.ListenAndServe())
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
