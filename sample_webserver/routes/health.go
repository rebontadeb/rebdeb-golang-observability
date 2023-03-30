package routes

import (
	"encoding/json"
	"net/http"
	"sample_webserver/selflogger"
)

type healthstatus struct {
	Status string
}

func Healtz(w http.ResponseWriter, req *http.Request) {
	healthstatus := healthstatus{Status: "Healthy"}
	jsonResponse, jsonError := json.Marshal(healthstatus)
	if jsonError != nil {
		selflogger.ErrorLogger.Println("Unable to encode JSON")
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}
}
