package main

import (
	"fmt"
	google "github.com/eunicebjm/gc/internal/google/geocoder"
	"github.com/eunicebjm/gc/internal/service"
	transporthttp "github.com/eunicebjm/gc/internal/transport/http"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

const (
	geocoderAPIKey = "google_api_key"
)

func main() {
	// Load config
	apiKey, ok := os.LookupEnv(geocoderAPIKey)
	if !ok{
		fmt.Println("missing google api key")
		return
	}

	// Setup google client
	googleGeocoder, err := google.NewGeocoderClient(apiKey)
	if err != nil {
		fmt.Println("failed to setup google geocoder client")
		return
	}

	// Setup service
	svc, err  := service.NewService(googleGeocoder)
	if err != nil {
		fmt.Println("failed to start service")
		return
	}

	// Setup http handler
	httpHandler, err := transporthttp.NewHandler(svc)
	if err != nil {
		fmt.Println("failed to setup http handler")
		return
	}

	// setup router
	router := mux.NewRouter()

	// endpoints, handler functions & HTTP method
	router.
		HandleFunc("/geocode", httpHandler.GeocodeOne).
		Methods("GET")
	http.Handle("/", router)

	http.ListenAndServe(":8080", router)
}