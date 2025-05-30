package webutils

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

func RespondWithError(rw http.ResponseWriter, code int, msg string) {
	errBody, _ := json.Marshal(errorResponse{Error: msg})
	rw.WriteHeader(code)
	rw.Write(errBody)
}

func RespondWithJSON(rw http.ResponseWriter, code int, payload interface{}) {
	rw.WriteHeader(code)
	responseBody, _ := json.Marshal(payload)
	rw.Write(responseBody)
}
