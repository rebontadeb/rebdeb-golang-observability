package server

import (
	"net/http"
	"sample_webserver/controllers"
	"sample_webserver/routes"
	"sample_webserver/selflogger"

	"github.com/gorilla/mux"
)

func Server(ServerPort string) {
	port := ":" + ServerPort

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/healthz", routes.Healtz).Methods("GET")
	r.HandleFunc("/api/redis/healthz", controllers.RedisHealthz).Methods("GET")
	r.HandleFunc("/api/data/employee/create", routes.CreateEmployeeData).Methods("POST")
	http.Handle("/", r)
	selflogger.InfoLogger.Println("Server Started At ", port)
	http.ListenAndServe(port, r)
}
