package main

import (
	"net/http"
	"log"
	"github.com/njackson/handlers"
	"os"
	"time"
	"context"
	"os/signal"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)
	
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)
	
	s := &http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	}
	
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()
	
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	
	sig := <-sigChan
	l.Println("\nReceived terminate, graceful shutdown", sig)
	
	
	tc, _ := context.WithTimeout(context.Background(),30*time.Second)
	s.Shutdown(tc)
}