package server

import (
	"net/http"
	"sample_webserver/routes"
	"sample_webserver/selflogger"

	"github.com/gorilla/mux"
)

func Server(ServerPort string) {
	port := ":" + ServerPort

	r := mux.NewRouter()
	r.HandleFunc("/healthz", routes.Healtz).Methods("GET")
	http.Handle("/", r)
	selflogger.InfoLogger.Println("Server Started At ", port)
	http.ListenAndServe(port, r)
}
