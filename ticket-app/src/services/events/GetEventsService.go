package events

import (
	"encoding/json"
	"fmt"
	"microservices/ticket/src/clients"
	"microservices/ticket/src/models"
	"net/http"
)

func GetEvents(w http.ResponseWriter, r *http.Request) {
	db := clients.GetDBInstance()

	fmt.Println("[!] [GET] [/events]")

	rows, err := db.Query("SELECT * FROM events")

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

	var response = models.EventJSONResponse{Type: "success", Data: events, Message: "Events retrieved successfully"}

	err = json.NewEncoder(w).Encode(response)

	if err != nil {
		panic(err)
	}
}
