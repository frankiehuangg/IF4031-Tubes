package routers

import (
	"fmt"
	"github.com/gorilla/mux"
	"microservices/ticket/src/services/events"
	"microservices/ticket/src/services/invoices"
	"microservices/ticket/src/services/seats"
	"microservices/ticket/src/utils"
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

	// seats
	router.HandleFunc("/seats", seats.GetSeats).Methods("GET")
	router.HandleFunc("/seats/{event_id}", seats.GetSeatDetail).Methods("GET")
	router.HandleFunc("/seats", seats.OrderSeat).Methods("PATCH")

	// invoices
	router.HandleFunc("/invoices", invoices.GetInvoice).Methods("GET")
	router.HandleFunc("/invoices", invoices.UpdateInvoice).Methods("PATCH")

	fmt.Printf("Server started at %s\n", utils.GoDotEnvVariable("TICKET_PORT"))

	return router
}
