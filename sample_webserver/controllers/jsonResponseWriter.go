package controllers

import (
	"encoding/json"
	"net/http"
)

type jsonMessage struct {
	Message string
	Status  string
}

func SuccessResponseWriter(w http.ResponseWriter, req *http.Request, responseString []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseString)
}

func FailureResponseWriter(w http.ResponseWriter, req *http.Request, responseString []byte) {
	jsonMessage := jsonMessage{Message: string(responseString), Status: "Failed"}
	marshalledJson, _ := json.Marshal(jsonMessage)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(marshalledJson)
}
