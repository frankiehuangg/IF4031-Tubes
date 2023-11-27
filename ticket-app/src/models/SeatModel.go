package models

type SeatModel struct {
	EventID    int    `json:"event_id"`
	ClientID   int    `json:"client_id"`
	SeatNumber int    `json:"seat_number"`
	SeatStatus string `json:"seat_status"`
}

type SeatJSONResponse struct {
	Type    string   `json:"type"`
	Data    []Events `json:"data"`
	Message string   `json:"message"`
}
