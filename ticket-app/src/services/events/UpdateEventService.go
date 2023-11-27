package events

import (
	"encoding/json"
	"fmt"
	"microservices/ticket/src/clients"
	"microservices/ticket/src/models"
	"net/http"
	"strings"
)

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	eventID := r.FormValue("event_id")
	eventName := r.FormValue("event_name")
	eventDate := r.FormValue("event_date")
	totalSeat := r.FormValue("total_seat")

	var response = models.EventJSONResponse{}

	if eventID == "" {
		response = models.EventJSONResponse{
			Type:    "error",
			Message: "Form value event_id is missing!",
		}
	} else if eventName == "" && eventDate == "" && totalSeat == "" {
		response = models.EventJSONResponse{
			Type:    "error",
			Message: "Please choose one column to update!",
		}
	} else {
		db := clients.GetDBInstance()

		var setQuery []string

		if eventName != "" {
			setQuery = append(setQuery, "event_name = '"+eventName+"'")
		}

		if eventDate != "" {
			setQuery = append(setQuery, "event_date = '"+eventDate+"'")
		}

		if totalSeat != "" {
			setQuery = append(setQuery, "total_seat = '"+totalSeat+"'")
		}

		updateQuery := strings.Join(setQuery, ", ")

		fmt.Printf("[!] [PATCH] [/events]\n")

		var retrievedEventId int
		var retrievedEventName string
		var retrievedEventDate string
		var retrievedTotalSeat int

		var events []models.Events

		eventErr := db.QueryRow(
			"UPDATE events SET "+updateQuery+" WHERE event_id = $1 RETURNING * ;", eventID,
		).Scan(
			&retrievedEventId,
			&retrievedEventName,
			&retrievedEventDate,
			&retrievedTotalSeat,
		)

		if eventErr != nil {
			panic(eventErr)
		}

		if totalSeat != "" {
			_, deleteSeatErr := db.Exec("DELETE FROM seats WHERE event_id = $1 ;", retrievedEventId)

			if deleteSeatErr != nil {
				panic(deleteSeatErr)
			}

			_, insertSeatErr := db.Exec("INSERT INTO seats (event_id, seat_number) SELECT $1, generate_series FROM generate_series(1, $2)",
				retrievedEventId,
				retrievedTotalSeat,
			)

			if insertSeatErr != nil {
				panic(insertSeatErr)
			}
		}

		events = append(events, models.Events{
			EventID:   retrievedEventId,
			EventName: retrievedEventName,
			EventDate: retrievedEventDate,
			TotalSeat: retrievedTotalSeat,
		})

		response = models.EventJSONResponse{
			Type:    "success",
			Data:    events,
			Message: "Event has been updated successfully",
		}
	}

	err := json.NewEncoder(w).Encode(response)

	if err != nil {
		panic(err)
	}
}
