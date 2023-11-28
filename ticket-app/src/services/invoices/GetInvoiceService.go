package invoices

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"microservices/ticket/src/models"
	"net/http"
)

func GetInvoice(w http.ResponseWriter, r *http.Request) {
	invoiceID := mux.Vars(r)["invoice_id"]

	var response = models.InvoiceJSONResponse{}

	if invoiceID == "" {
		response = models.InvoiceJSONResponse{
			Type:    "error",
			Message: "URL params invoice_id is not set!",
		}
	} else {
		fmt.Println("[!] [GET] [/invoice]")

		resp, err := http.Get("http://localhost:7000/invoices?invoice_id=" + invoiceID)

		if err != nil {
			panic(err)
		}

		body, err := ioutil.ReadAll(resp.Body)

		response = models.InvoiceJSONResponse{
			Type:    "success",
			Message: string(body),
		}
	}

	err := json.NewEncoder(w).Encode(response)

	if err != nil {
		panic(err)
	}
}
