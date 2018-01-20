package main

import (
	"io/ioutil"
	"encoding/json"
	"crypto-jobs/pkg/models"
	"crypto-jobs/pkg/database"
	"log"
	"net/http"
	"strconv"
	"crypto-jobs/pkg/routes"
	"crypto-jobs/pkg/engine/job"
)

var config models.Config

func main() () {
	configure()

	//TODO: Some way to resume all jobs from last run. Possibly by storing them in a json file or something.
	go job.StartEngine(config)

	router := models.CreateRouter()
	router = routes.CreateRoutes(router)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.Port), router))
}

func configure() () {
	populateConfig()
	err := database.InitializeDatabase(config)
	if err != nil {
		panic(err)
	}
}

func populateConfig() () {
	jsonFile, err := ioutil.ReadFile("../config/config.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(jsonFile, &config)
	if err != nil {
		panic(err)
	}
}
