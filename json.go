package main

import (
	"encoding/json"
	"net/http"
)

func respondWithJSON(writer http.ResponseWriter, status int, payload any) {
	writer.WriteHeader(status)
	bytes, err := json.Marshal(payload)
	if err != nil {
		respondWithJSON(writer, 500, err.Error())
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	_, _ = writer.Write(bytes)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)
	type MyErr struct {
		Error string `json:"error"`
	}
	myErr := MyErr{msg}
	bytes, _ := json.Marshal(myErr)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(bytes)
}
