package main

import (
	"log"
	"net/http"


	"github.com/daywon03/Kiln-monitoring-microservice/internal/config"
)

func main() {

	
	cfg, err := config.Load("configs/config.example.yaml")
	if err != nil {
		log.Fatal("Erreur lors du chargement de la config:", err)
	}
	log.Println("Configuration chargée avec succès!")
	log.Printf("Base URL: %s", cfg.Kiln.BaseUrl)
	log.Printf("Intervalle: %v", cfg.Monitoring.Interval)
	
	log.Println("Starting server on port 8080")
	http.HandleFunc("/healtz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
	log.Println("listening on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}