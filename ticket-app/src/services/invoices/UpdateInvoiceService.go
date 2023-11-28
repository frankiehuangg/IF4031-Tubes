package invoices

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"microservices/ticket/src/models"
	"net/http"
)

func UpdateInvoice(w http.ResponseWriter, r *http.Request) {
	invoiceID := r.FormValue("invoice_id")
	paymentStatus := r.FormValue("payment_status")

	var response = models.InvoiceJSONResponse{}

	if paymentStatus == "" {
		response = models.InvoiceJSONResponse{
			Type:    "error",
			Message: "URL params payment_status is not set!",
		}
	} else if invoiceID == "" {
		fmt.Println("[!] [PATCH] [/invoice]")

		postBody, _ := json.Marshal(map[string]string{
			"invoice_id":     invoiceID,
			"payment_status": paymentStatus,
		})

		responseBody := bytes.NewBuffer(postBody)

		req, err := http.NewRequest("PATCH", "http://localhost:7000/invoices", responseBody)

		if err != nil {
			panic(err)
		}

		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}

		resp, err := client.Do(req)
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

		response = models.InvoiceJSONResponse{
			Type:    "success",
			Message: "Invoice is sent successfully",
		}
	} else {

	}

	err := json.NewEncoder(w).Encode(response)

	if err != nil {
		panic(err)
	}
}
