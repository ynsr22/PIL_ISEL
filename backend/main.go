package main

import (
    "backend/database"
    "backend/routes"
    "log"
    "net/http"
)

func main() {
    // Connexion à la base de données
    if err := database.ConnectDB(); err != nil {
        log.Fatalf("Erreur de connexion à la base de données: %v", err)
    }
    
    // Configuration du routeur
    r := routes.SetupRouter()
    
    // Démarrage du serveur avec http.ListenAndServe
    log.Println("Serveur démarré sur 0.0.0.0:8000")
    if err := http.ListenAndServe("0.0.0.0:8000", r); err != nil {
        log.Fatalf("Erreur lors du démarrage du serveur: %v", err)
    }
}