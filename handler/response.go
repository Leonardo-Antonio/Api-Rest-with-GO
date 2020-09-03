package handler

import (
	"encoding/json"
	"net/http"
)

const (
	// Error .
	Error = "error"
	// Message .
	Message = "messaje"
)

type response struct {
	MessageType string      `json:"message-type"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
}

func newRespone(mjsT, mjs string, data interface{}) *response {
	return &response{mjsT, mjs, data}
}

func (r *response) responseJSON(w http.ResponseWriter, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	data := response{r.MessageType, r.Message, r.Data}
	err := json.NewEncoder(w).Encode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
