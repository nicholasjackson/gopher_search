package main

import (
	"log"
	"os"

	"github.com/nicholasjackson/gopher_search/actions"
)

func main() {
	log.Println("db_url", os.Getenv("DATABASE_URL"))
	log.Println("env", os.Getenv("GO_ENV"))

	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
