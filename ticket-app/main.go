package main

import (
	"fmt"
	"log"
	"microservices/ticket/src/routers"
	"microservices/ticket/src/utils"
	"net/http"
	kafka "microservices/ticket/src/clients"
)

func main() {
	go kafka.StartKafka()	
	router := routers.SetupRoutes()

	port := utils.GoDotEnvVariable("PORT")
	addr := fmt.Sprintf(":%s", port)

	log.Fatal(http.ListenAndServe(addr, router))
}
