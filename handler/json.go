package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, statusCode int, msg string) {
	if statusCode > 499 {
		log.Println("Responging with 5XX error:", msg)
	}

	type errRes struct {
		Error string `json:"error"`
	}
	RespondWithJSON(w, statusCode, errRes{
		Error: msg,
	})
}

func RespondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Println("failed to marshal json response", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(dat)
}
