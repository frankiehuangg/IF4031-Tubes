package events

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"microservices/ticket/src/clients"
	"microservices/ticket/src/models"
	"net/http"
)

func GetEventDetail(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["event_id"]

	var response = models.EventJSONResponse{}

	if eventID == "" {
		response = models.EventJSONResponse{
			Type:    "error",
			Message: "URL params event_id is not set!",
		}
	} else {
		db := clients.GetDBInstance()

		fmt.Printf("[!] [GET] [/events/%s]\n", eventID)

		rows, err := db.Query("SELECT * FROM events WHERE event_id = $1", eventID)

		if err != nil {
			panic(err)
		}
		var events []models.Events

		for rows.Next() {
			var retrievedEventID int
			var retrievedEventName string
			var retrievedEventDate string
			var retrievedTotalSeat int

			err = rows.Scan(&retrievedEventID, &retrievedEventName, &retrievedEventDate, &retrievedTotalSeat)

			if err != nil {
				panic(err)
			}

			events = append(events, models.Events{
				EventID:   retrievedEventID,
				EventName: retrievedEventName,
				EventDate: retrievedEventDate,
				TotalSeat: retrievedTotalSeat,
			})
		}

		response = models.EventJSONResponse{Type: "success", Data: events, Message: "Event data retrieved successfully"}
	}

	err := json.NewEncoder(w).Encode(response)

	if err != nil {
		return
	}
}
