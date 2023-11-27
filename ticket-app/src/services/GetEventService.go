package services

import (
	"encoding/json"
	"fmt"
	"microservices/ticket/src/clients"
	"microservices/ticket/src/models"
	"net/http"
	"time"
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

	var response = models.EventJSONResponse{Type: "success", Data: events, Message: "Events retrieved successfully"}

	err = json.NewEncoder(w).Encode(response)

	if err != nil {
		panic(err)
	}
}
