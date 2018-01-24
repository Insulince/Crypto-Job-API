package main

import (
	"crypto-jobs/pkg/models"
	"crypto-jobs/pkg/database"
	"log"
	"net/http"
	"strconv"
	"crypto-jobs/pkg/routes"
	"crypto-jobs/pkg/engine/job"
	"crypto-jobs/pkg/models/configuration"
	"fmt"
)

func main() () {
	config := configuration.GetConfiguration()

	err := database.InitializeDatabase(config)
	if err != nil {
		panic(err)
	}

	go job.RunEngine()

	fmt.Println("asdf")
	router := models.CreateRouter()
	router = routes.CreateRoutes(router)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.Port), router))
}
