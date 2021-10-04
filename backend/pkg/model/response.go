package model

import (
	"encoding/json"
	"net/http"
)

type ErrorResponseDTO struct {
	Code    int      `json:"code"`
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}

// RespondWithError ...
func RespWithError(w http.ResponseWriter, code int, message string) {
	RespWithJSON(w, code, ErrorResponseDTO{Code: code, Status: "Error", Message: message})
}

// RespondWithJSON write json
func RespWithJSON(w http.ResponseWriter, code int, data interface{}) {
	response, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
