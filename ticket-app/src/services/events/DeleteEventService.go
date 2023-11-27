package events

import (
	"encoding/json"
	"fmt"
	"microservices/ticket/src/clients"
	"microservices/ticket/src/models"
	"net/http"
)

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	eventID := r.FormValue("event_id")

	var response = models.EventJSONResponse{}

	if eventID == "" {
		response = models.EventJSONResponse{
			Type:    "error",
			Message: "Form value event_id is missing!",
		}
	} else {
		db := clients.GetDBInstance()

		fmt.Printf("[!] [DELETE] [/events]\n")

		_, err := db.Exec("DELETE FROM events WHERE event_id = $1 ; ", eventID)

		if err != nil {
			panic(err)
		}

		response = models.EventJSONResponse{
			Type:    "success",
			Message: "Event has been deleted successfully",
		}
	}

	err := json.NewEncoder(w).Encode(response)

	if err != nil {
		panic(err)
	}
}
