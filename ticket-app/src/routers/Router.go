package routers

import (
	"fmt"
	"github.com/gorilla/mux"
	"microservices/ticket/src/services"
	"os"
)

func SetupRoutes() *mux.Router {
	// Initialize the mux router
	router := mux.NewRouter()

	router.HandleFunc("/events", services.GetEvents).Methods("GET")
	router.HandleFunc("/events/{event_id}", services.GetEventDetail).Methods("GET")

	fmt.Printf("Server started at %s", os.Getenv("PORT"))

	return router
}
