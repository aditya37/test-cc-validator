package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aditya37/test-cc-validator/repository"
	"github.com/aditya37/test-cc-validator/service"
	"github.com/aditya37/test-cc-validator/transport"
	"github.com/gorilla/mux"
)

func main() {
	repo := repository.NewCardReaderWriter()
	svc := service.NewService(repo)
	handler := transport.NewTransport(svc)

	// router or path
	router := mux.NewRouter()
	router.Methods(http.MethodGet).Path("/").HandlerFunc(handler.HealthCheck)
	router.Methods(http.MethodPost).Path("/api/check-card").HandlerFunc(handler.GetCardNumber)
	router.Methods(http.MethodPost).Path("/api/card").HandlerFunc(handler.InsertCard)

	// server
	errs := make(chan error)
	log.Printf("Running on Port %d", 1234)
	errs <- server(router, 1234)

	log.Fatal(<-errs)
}

func server(handler *mux.Router, port int) error {
	addr := fmt.Sprintf(":%d", port)
	serve := &http.Server{
		Handler: handler,
		Addr:    addr,
	}
	return serve.ListenAndServe()
}
