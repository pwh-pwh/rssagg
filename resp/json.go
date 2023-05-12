package resp

import (
	"encoding/json"
	"net/http"
)

func RespondWithJSON(writer http.ResponseWriter, status int, payload any) {
	bytes, err := json.Marshal(payload)
	if err != nil {
		RespondWithJSON(writer, 500, err.Error())
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	_, _ = writer.Write(bytes)
}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	type MyErr struct {
		Error string `json:"error"`
	}
	myErr := MyErr{msg}
	bytes, _ := json.Marshal(myErr)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(bytes)
}
