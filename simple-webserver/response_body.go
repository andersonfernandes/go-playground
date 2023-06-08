package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type responseBody struct {
	Message string            `json:"message"`
	Data    map[string]string `json:"data"`
}

func (rb *responseBody) writeJsonResponse(w http.ResponseWriter) {
	jsonResp, err := json.Marshal(rb)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}
