package main

import (
	"fmt"
	"log"
	"microservices/ticket/src/routers"
	"microservices/ticket/src/utils"
	"net/http"
)

func main() {
	router := routers.SetupRoutes()

	port := utils.GoDotEnvVariable("PORT")
	addr := fmt.Sprintf(":%s", port)

	log.Fatal(http.ListenAndServe(addr, router))
}
