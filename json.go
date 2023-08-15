package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func decodeJSONData(r *http.Request) (map[string]interface{}, error) {
	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	return data, err
}

func responseWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5XX errors:", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errResponse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshan JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
