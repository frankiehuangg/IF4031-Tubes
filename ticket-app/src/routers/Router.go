package routers

import (
	"fmt"
	"github.com/gorilla/mux"
	"microservices/ticket/src/services/events"
	"os"
)

func SetupRoutes() *mux.Router {
	// Initialize the mux router
	router := mux.NewRouter()

	// List of routes

	// events
	router.HandleFunc("/events", events.CreateEvent).Methods("POST")
	router.HandleFunc("/events", events.GetEvents).Methods("GET")
	router.HandleFunc("/events/{event_id}", events.GetEventDetail).Methods("GET")
	router.HandleFunc("/events", events.UpdateEvent).Methods("PATCH")
	router.HandleFunc("/events", events.DeleteEvent).Methods("DELETE")

	fmt.Printf("Server started at %s\n", os.Getenv("PORT"))

	return router
}
