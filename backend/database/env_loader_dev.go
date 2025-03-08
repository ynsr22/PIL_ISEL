//go:build !production
package database

import (
    "log"
    "github.com/joho/godotenv"
    "path/filepath"
)

// LoadEnv charge `.env` uniquement en Dev
func LoadEnv() {
    rootPath, err := filepath.Abs("../../") // Cherche le fichier dans le dossier parent
    if err == nil {
        envPath := filepath.Join(rootPath, ".env")
        err = godotenv.Load(envPath)
        if err != nil {
            log.Println("⚠️ Impossible de charger .env, utilisation des variables d'environnement système")
        }
    }
}
