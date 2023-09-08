package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rammyblog/rent-movie/database"
	"github.com/rammyblog/rent-movie/router"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	// Connect to db
	db, err := database.Init()
	if err != nil {
		log.Fatal("Could not connect to db")
		panic(err)
	}
	fmt.Println("db conected", db.Name())

	port := fmt.Sprintf(":%v", os.Getenv("PORT"))
	routersInit := router.InitRouter()

	server := &http.Server{
		Addr:    port,
		Handler: routersInit,
	}

	log.Printf("[info] start http server listening %s", port)

	server.ListenAndServe() // listen and serve on 0.0.0.0:8080
}
