package core

import (
	"log"
	"net/http"
)

func GoServerNoFrontend() {
	// Ajoute ici tes handlers/routes si besoin
	log.Println("Serveur HTTP sans SSL lancé sur le port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Erreur serveur HTTP:", err)
	}
}
