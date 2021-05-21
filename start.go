package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/common-nighthawk/go-figure"
	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "EAG", log.LstdFlags)

	serverMux := mux.NewRouter()

	productServer := &http.Server{
		Addr:         ":8080",
		Handler:      serverMux,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		ErrorLog:     l,
	}

	go func() {
		l.Println("Starting the EAG server.....")
		myFigure := figure.NewFigure("EAG", "", true)
		myFigure.Print()

		err := productServer.ListenAndServe()

		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			//l.Fatal(err)
			os.Exit(1)

		} else {
			l.Println("EAG running.....")
		}

	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	sig := <-sigChan

	l.Println("The EAG server has received a shutdown request.. Shutting down", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	productServer.Shutdown(tc)

}
