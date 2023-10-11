package auxiliar

import (
	"encoding/json"
	"log"
	"net/http"
)

// -------------------- Auxiliar Functions -----------------------

// ConfigureHttpResponse is a function that taken a Response writer
// a status code and a message will send a proper JSON answer back
func ConfigureHttpResponse(rw http.ResponseWriter, statusCode int, msg string) {
	rw.WriteHeader(statusCode)
	rw.Header().Set("Content-Type", "application/json")
	resp := map[string]string{
		"message": msg,
	}

	jsonResp, err := json.Marshal(resp)

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	rw.Write(jsonResp)
}
