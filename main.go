package main

import (
	"log"
	"net/http"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/daywon03/Kiln-monitoring-microservice/internal/config"
	"github.com/daywon03/Kiln-monitoring-microservice/internal/collector"
	"github.com/daywon03/Kiln-monitoring-microservice/internal/kiln"
)

func main() {

	
	cfg, err := config.Load("configs/config.example.yaml")
	if err != nil {
		log.Fatal("Erreur lors du chargement de la config:", err)
	}
	log.Println("Configuration chargée avec succès!")
	log.Printf("Base URL: %s", cfg.Kiln.BaseUrl)
	log.Printf("Intervalle: %v", cfg.Monitoring.Interval)
	
	//init kiln
	client := kiln.NewClient(cfg.Kiln.BaseUrl, cfg.Kiln.Token, 10*time.Second)
	log.Printf("kiln client inittialized with baseUrl : %s", cfg.Kiln.BaseUrl)

	//create collector
	collecteur := collector.New(client, cfg)


	//context for shutdown
	context, cancel := context.WithCancel(context.Background())
	defer cancel()
	
	
	//start collector
	if err := collecteur.Start(context); err != nil {
		log.Fatal("Failed to start collecteur: ", err)
	}
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok","service: "kiln-monitoring"}`))
	})
	log.Println("listening on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))


	go func() {
		log.Println("HTTP server started on port :8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Println("http server error %v", err)
		}
	}()

	//wait for shutdown signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	log.Println("Microservice started")
	<-sigChan

	log.Println("Shutdown signal received")
    collecteur.Stop()
    cancel()
    log.Println("Microservice stopped")


}