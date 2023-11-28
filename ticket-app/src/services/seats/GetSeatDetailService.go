package seats

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"microservices/ticket/src/clients"
	"microservices/ticket/src/models"
	"net/http"
)

func GetSeatDetail(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["event_id"]
	seatNumber := r.URL.Query().Get("seat_number")

	var response = models.SeatJSONResponse{}

	if eventID == "" {
		response = models.SeatJSONResponse{
			Type:    "error",
			Message: "URL Params event_id is missing!",
		}
	} else if seatNumber == "" {
		response = models.SeatJSONResponse{
			Type:    "error",
			Message: "URL Params seat_number is missing!",
		}
	} else {
		db := clients.GetDBInstance()

		fmt.Printf("[!] [GET] [/seats/%s]\n", eventID)

		rows, err := db.Query("SELECT * FROM seats WHERE event_id = $1 AND seat_number = $2 ;", eventID, seatNumber)

		if err != nil {
			panic(err)
		}

		var seats []models.Seats

		for rows.Next() {
			var retrievedSeatID int
			var retrievedEventID int
			var retrievedSeatNumber int
			var retrievedSeatStatus string

			err = rows.Scan(&retrievedSeatID, &retrievedEventID, &retrievedSeatNumber, &retrievedSeatStatus)

			if err != nil {
				panic(err)
			}

			seats = append(seats, models.Seats{
				SeatID:     retrievedSeatID,
				EventID:    retrievedEventID,
				SeatNumber: retrievedSeatNumber,
				SeatStatus: retrievedSeatStatus,
			})
		}

		response = models.SeatJSONResponse{
			Type:    "success",
			Data:    seats,
			Message: "Seat data retrieved successfully",
		}
	}

	err := json.NewEncoder(w).Encode(response)

	if err != nil {
		panic(err)
	}
}
