package main

import (
	"fmt"
	"log"
	"net/http"
	"web_adv/handler"
	"web_adv/middleware"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	router := mux.NewRouter()
	router.Handle("/metrics", promhttp.Handler())
	router.HandleFunc("/random", handler.RandomRequest).Methods(http.MethodGet)

	fmt.Println("HTTP server listening in 8080")
	server := &http.Server{
		Addr:    ":8080",
		Handler: middleware.WithMetrics(router),
	}
	log.Fatal(server.ListenAndServe())
}
