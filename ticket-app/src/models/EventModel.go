package models

type Events struct {
	EventID   int    `json:"event_id"`
	EventName string `json:"event_name"`
	EventDate string `json:"event_date"`
	TotalSeat int    `json:"total_seat"`
}

type EventJSONResponse struct {
	Type    string   `json:"type"`
	Data    []Events `json:"data"`
	Message string   `json:"message"`
}
