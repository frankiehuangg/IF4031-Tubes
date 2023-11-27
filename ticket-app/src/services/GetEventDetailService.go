package services

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"microservices/ticket/src/clients"
	"microservices/ticket/src/models"
	"net/http"
	"time"
)

func GetEventDetail(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["event_id"]

	var response = models.EventJSONResponse{}

	if eventID == "" {
		response = models.EventJSONResponse{Type: "error", Message: "event id parameter is missing!"}
	} else {
		db := clients.GetDBInstance()

		fmt.Printf("[!] [GET] [/events/%s]\n", eventID)

		rows, err := db.Query("SELECT * FROM events WHERE event_id = $1", eventID)

		if err != nil {
			panic(err)
		}
		var events []models.Events

		for rows.Next() {
			var EventID int
			var EventName string
			var EventDate time.Time

			err = rows.Scan(&EventID, &EventName, &EventDate)

			if err != nil {
				panic(err)
			}

			events = append(events, models.Events{
				EventID:   EventID,
				EventName: EventName,
				EventDate: EventDate,
			})
		}

		response = models.EventJSONResponse{Type: "success", Data: events, Message: "Event data retrieved successfully"}
	}

	err := json.NewEncoder(w).Encode(response)

	if err != nil {
		return
	}
}
