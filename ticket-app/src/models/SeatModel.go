package models

type Seats struct {
	SeatID     int    `json:"seat_id"`
	EventID    int    `json:"event_id"`
	SeatNumber int    `json:"seat_number"`
	SeatStatus string `json:"seat_status"`
}

type SeatJSONResponse struct {
	Type    string  `json:"type"`
	Data    []Seats `json:"data"`
	Message string  `json:"message"`
}
