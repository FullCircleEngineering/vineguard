package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"vineguard/internal/cache"
	"vineguard/internal/registrations"

	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGKILL, syscall.SIGTERM)

	// setup CORS
	corsManager := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		Debug:            true,
		AllowCredentials: true,
	})

	// setup and run API server
	RegHandler := corsManager.Handler(registrations.NewHander())

	httpMux := http.NewServeMux()
	httpMux.Handle("/api/v0/registrations", RegHandler)

	svr := &http.Server{
		Addr:         ":8080",
		Handler:      httpMux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	// run server
	go func() {
		err := svr.ListenAndServe()
		if err != nil {
			logrus.Errorf("server closed. %s", err)
			c <- syscall.SIGINT
		}
	}()

	// run cache updater
	go func() {
		svc := cache.NewService()
		err := svc.Run()
		if err != nil {
			logrus.Errorf("Cache updater failed. %s", err)
			c <- syscall.SIGINT
		}
	}()

	<-c
}
