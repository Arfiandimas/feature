package main

import (
	log "github.com/sirupsen/logrus"
	"mandiri_payment_gateway_service/routes"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	/*
		LOGRUS INIT
	*/
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	println("App started")
	err := godotenv.Load()
	if err != nil {
		return
	}
	routes.Init()
}
