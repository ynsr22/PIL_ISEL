package main

import (
    "backend/database"
    "backend/routes"
)

func main() {
    // Connexion à la base de données
    database.ConnectDB()

    // Configuration du routeur et démarrage du serveur
    r := routes.SetupRouter()
    r.Run("0.0.0.0:8000")
}
