package database

import (
    "fmt"
    "log"
    "os"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
    LoadEnv() // Charge `.env` en Dev, mais ne fait rien en Prod (variables système)

    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
        os.Getenv("DB_HOST"),
        os.Getenv("POSTGRES_USER"),
        os.Getenv("POSTGRES_PASSWORD"),
        os.Getenv("POSTGRES_DB"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_SSLMODE"),
    )

    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("❌ Erreur de connexion à la base de données : %v", err)
    }

    log.Println("✅ Connexion réussie à la base de données !")
}
