package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/nicholasjackson/gopher_search/actions"
)

func main() {
	//loadConfig()
	log.Println("env", os.Getenv("GO_ENV"))

	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}

func loadConfig() {
	path := "/etc/secrets/config"

	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	parts := strings.Split(string(file), "=")
	os.Setenv(parts[0], strings.Replace(parts[1], "\"", "", -1))
}
