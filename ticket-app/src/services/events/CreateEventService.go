package events

import (
	"encoding/json"
	"fmt"
	"microservices/ticket/src/clients"
	"microservices/ticket/src/models"
	"net/http"
)

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	eventName := r.FormValue("event_name")
	eventDate := r.FormValue("event_date")
	totalSeat := r.FormValue("total_seat")

	var response = models.EventJSONResponse{}

	if eventName == "" {
		response = models.EventJSONResponse{
			Type:    "error",
			Message: "event name form value is missing!",
		}
	} else if eventDate == "" {
		response = models.EventJSONResponse{
			Type:    "error",
			Message: "event date form value is missing!",
		}
	} else if totalSeat == "" {
		response = models.EventJSONResponse{
			Type:    "error",
			Message: "total seat form value is missing!",
		}
	} else {
		db := clients.GetDBInstance()

		fmt.Printf("[!] [POST] [/events]\n")

		var retrievedEventId int
		var retrievedEventName string
		var retrievedEventDate string
		var retrievedTotalSeat int

		var events []models.Events

		eventErr := db.QueryRow(
			"INSERT INTO events (event_name, event_date, total_seat) VALUES ($1, $2, $3) RETURNING * ;",
			eventName,
			eventDate,
			totalSeat,
		).Scan(
			&retrievedEventId,
			&retrievedEventName,
			&retrievedEventDate,
			&retrievedTotalSeat,
		)

		if eventErr != nil {
			panic(eventErr)
		}

		db.QueryRow(
			"INSERT INTO seats (event_id, seat_number) SELECT $1, generate_series FROM generate_series(1, $2) ;",
			retrievedEventId,
			retrievedTotalSeat,
		)

		events = append(events, models.Events{
			EventID:   retrievedEventId,
			EventName: retrievedEventName,
			EventDate: retrievedEventDate,
			TotalSeat: retrievedTotalSeat,
		})

		response = models.EventJSONResponse{
			Type:    "success",
			Data:    events,
			Message: "Event has been created successfully",
		}
	}

	err := json.NewEncoder(w).Encode(response)

	if err != nil {
		panic(err)
	}
}
