package main

import (
	"fmt"
	"github.com/eunicebjm/wordCounter/internal/parser"
	"github.com/eunicebjm/wordCounter/internal/reader"
	"github.com/eunicebjm/wordCounter/internal/service"
	transporthttp "github.com/eunicebjm/wordCounter/internal/transport/http"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// Setup service
	reader, err := reader.New()
	if err != nil {
		log.Fatal("creating reader", err)
		return
	}
	parser, err := parser.New()
	if err != nil {
		log.Fatal("creating parser", err)
		return
	}

	svc := service.New(reader, parser)
	if err != nil {
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
		HandleFunc("/count", httpHandler.CountWords).
		Methods("GET")
	http.Handle("/", router)

	http.ListenAndServe(":8080", router)
}
