//go:build !production
package database

import (
    "log"
    "bufio"
    "os"
    "strings"
    "sync"
)

var loadEnvOnce sync.Once


func LoadEnv() { // LoadEnv charge les variables d'environnement
    loadEnvOnce.Do(func() { // Garantie une exécution unique
        file, err := os.Open("../.env") // Ouvre le fichier .env
        if err != nil {
            log.Println("Impossible d'ouvrir .env, utilisation des variables système")
            return
        }
        defer file.Close()

        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
            line := scanner.Text()

            // Ignorer les lignes vides et les commentaires
            if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
                continue
            }

            // Séparer la clé et la valeur
            parts := strings.SplitN(line, "=", 2)
            if len(parts) != 2 {
                continue
            }

            key := strings.TrimSpace(parts[0])
            value := strings.TrimSpace(parts[1])

            // Supprimer les guillemets autour de la valeur, si présents
            value = strings.Trim(value, `"'`)

            os.Setenv(key, value)
        }

        if err := scanner.Err(); err != nil {
            log.Println("Erreur lors de la lecture de .env:", err)
        } else {
            log.Println("Variables d'environnement chargées depuis .env")
        }
    })
}
