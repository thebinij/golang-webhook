package main

import (
	"fmt"
	"net/http"
)

func handleWebhookEvent(w http.ResponseWriter, r *http.Request) {

	data, err := decodeJSONData(r)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	fmt.Printf("Webhook event data: %+v\n", data)

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Webhook event received"})
}
