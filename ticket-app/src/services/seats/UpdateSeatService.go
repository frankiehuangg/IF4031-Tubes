package seats

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"microservices/ticket/src/clients"
	"microservices/ticket/src/models"
	"net/http"
	"os"
	"strconv"
	"time"
)

func OrderSeat(w http.ResponseWriter, r *http.Request) {
	clientID := r.FormValue("client_id")
	seatNumber := r.FormValue("seat_number")
	eventID := r.FormValue("event_id")

	var response = models.InvoiceJSONResponse{}

	if clientID == "" {
		response = models.InvoiceJSONResponse{
			Type:    "error",
			Message: "Form value client_id is missing!",
		}
	} else if seatNumber == "" {
		response = models.InvoiceJSONResponse{
			Type:    "error",
			Message: "Form value seat_number is missing!",
		}
	} else if eventID == "" {
		response = models.InvoiceJSONResponse{
			Type:    "error",
			Message: "Form value event_id is missing!",
		}
	} else {
		random := rand.New(rand.NewSource(time.Now().UnixNano()))

		num := random.Intn(100)
		log.Println(num)

		if num <= 80 {
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

			val, err := strconv.Atoi(clientID)

			fmt.Println(val, retrievedSeatID)

			postBody, _ := json.Marshal(map[string]int{
				"client_id": val,
				"seat_id":   retrievedSeatID,
			})

			fmt.Println(postBody)

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

			out, err := ioutil.ReadAll(resp.Body)

			if err != nil {
				panic(err)
			}

			type Response struct {
				Status  bool           `json:"success"`
				Message string         `json:"message"`
				Data    models.Invoice `json:"data"`
			}

			var res Response

			fmt.Println(res)

			err = json.Unmarshal([]byte(out), &res)

			if err != nil {
				panic(err)
			}

			var invoices []models.Invoice

			invoices = append(invoices, res.Data)

			if res.Status == true {
				response = models.InvoiceJSONResponse{
					Type:    "success",
					Data:    invoices,
					Message: "Seat status has been updated successfully!",
				}
			} else {
				response = models.InvoiceJSONResponse{
					Type:    "error",
					Message: "Invoice ID not found",
				}
			}
		} else {
			pdf := gofpdf.New("P", "mm", "A4", "")
			pdf.AddPage()
			pdf.SetFont("Arial", "B", 16)
			pdf.Text(40, 10, "Booking failed: cannot reach external event host")

			err := pdf.OutputFileAndClose("./out.pdf")

			if err != nil {
				log.Println("ERROR", err.Error())
			}

			file, err := os.Open("./out.pdf")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer file.Close()

			content, _ := ioutil.ReadFile("./out.pdf")

			resp, err := http.Post("http://client-app:8000/api/invoices", "application/json", bytes.NewBuffer(content))

			fmt.Println(resp)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			response = models.InvoiceJSONResponse{
				Type:    "error",
				Message: "Error fetching seat",
			}
		}
	}

	err := json.NewEncoder(w).Encode(response)

	if err != nil {
		panic(err)
	}
}
