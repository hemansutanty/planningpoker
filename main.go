package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/hemansutanty/planningpoker/handlers"
)

func main() {
	planningPokerLogger := log.New(os.Stdout, "Collaborative Planning App", log.LstdFlags)
	planningPokerHandler := handlers.NewPlanningPokerMeta(planningPokerLogger)
	planningPokerMux := mux.NewRouter()

	// Register POST route for out endpoint
	postRouter := planningPokerMux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/planningpoker/v1/create-poll", planningPokerHandler.CreatePoll)
	postRouter.Use(planningPokerHandler.MiddlewareCreatePollRequestValidation)

	//serve swagger for the api
	getRouter := planningPokerMux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", planningPokerHandler.Welcome)
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)
	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	server := http.Server{
		Handler:      planningPokerMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			planningPokerLogger.Fatal(err)
		}
	}()

	// Graceful shutdown of our server on any kind of OS kill/interrupt signals
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	planningPokerLogger.Println("Recieved terminate, graceful shutdown", sig)
	server.ListenAndServe()
	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	server.Shutdown(tc)
}
