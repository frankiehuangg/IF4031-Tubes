package models

type Invoice struct {
	InvoiceID     int    `json:"invoice_id"`
	ClientID      int    `json:"client_id"`
	Priority      int    `json:"priority"`
	PaymentStatus string `json:"payment_status"`
	SeatID        int    `json:"seat_id"`
}

type InvoiceJSONResponse struct {
	Type    string    `json:"type"`
	Data    []Invoice `json:"data"`
	Message string    `json:"message"`
}
