package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/stephendatascientist/go-postgres-api/database"
	"github.com/stephendatascientist/go-postgres-api/models"
	"github.com/stephendatascientist/go-postgres-api/routers"
)

func init() {

	// Load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

}

func main() {

	// Connect the database
	db, err := database.NewConnection()
	if err != nil {
		log.Fatal("could not load the database")
	}

	// Migrate the database
	err = models.MigrateEmployee(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}

	r := routers.Routers(db)

	fmt.Println("Server is listening at port 8080 ...")
	log.Fatal(http.ListenAndServe(":8080", r))

}
