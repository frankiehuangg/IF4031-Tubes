package models

import "time"

type Events struct {
	EventID   int       `json:"event_id"`
	EventName string    `json:"event_name"`
	EventDate time.Time `json:"event_date"`
}

type EventJSONResponse struct {
	Type    string   `json:"type"`
	Data    []Events `json:"data"`
	Message string   `json:"message"`
}
