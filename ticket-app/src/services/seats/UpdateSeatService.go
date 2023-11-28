package seats

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"microservices/ticket/src/clients"
	"microservices/ticket/src/models"
	"net/http"
	"strconv"
)

func OrderSeat(w http.ResponseWriter, r *http.Request) {
	clientID := r.FormValue("client_id")
	seatNumber := r.FormValue("seat_number")
	eventID := r.FormValue("event_id")

	var response = models.SeatJSONResponse{}

	if clientID == "" {
		response = models.SeatJSONResponse{
			Type:    "error",
			Message: "Form value client_id is missing!",
		}
	} else if seatNumber == "" {
		response = models.SeatJSONResponse{
			Type:    "error",
			Message: "Form value seat_number is missing!",
		}
	} else if eventID == "" {
		response = models.SeatJSONResponse{
			Type:    "error",
			Message: "Form value event_id is missing!",
		}
	} else {
		db := clients.GetDBInstance()

		var retrievedSeatID int
		var retrievedEventID int
		var retrievedSeatNumber int
		var retrievedSeatStatus string

		var seats []models.Seats

		err := db.QueryRow(
			"UPDATE seats SET seat_status = 'waiting' WHERE event_id = $1 AND seat_number = $2 RETURNING * ; ",
			eventID,
			seatNumber,
		).Scan(
			&retrievedSeatID,
			&retrievedEventID,
			&retrievedSeatNumber,
			&retrievedSeatStatus,
		)

		if err != nil {
			panic(err)
		}

		seats = append(seats, models.Seats{
			SeatID:     retrievedSeatID,
			EventID:    retrievedEventID,
			SeatNumber: retrievedSeatNumber,
			SeatStatus: retrievedSeatStatus,
		})

		postBody, _ := json.Marshal(map[string]string{
			"client_id": clientID,
			"seat_id":   strconv.Itoa(retrievedSeatID),
		})

		responseBody := bytes.NewBuffer(postBody)

		resp, err := http.Post("http://payment-app:8000/invoices", "application/json", responseBody)

		if err != nil {
			panic(err)
		}

		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(resp.Body)

		_, err = ioutil.ReadAll(resp.Body)

		if err != nil {
			panic(err)
		}

		response = models.SeatJSONResponse{
			Type:    "success",
			Data:    seats,
			Message: "Seat status has been updated successfully!",
		}
	}

	err := json.NewEncoder(w).Encode(response)

	if err != nil {
		panic(err)
	}
}
