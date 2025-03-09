package database

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/jackc/pgx/v5/pgxpool"
)

// DB représente le pool de connexions PostgreSQL
var DB *pgxpool.Pool

// ConnectDB initialise un pool de connexions
func ConnectDB() error {
    LoadEnv() // Charge `.env` en Dev, ne fait rien en Prod

    dsn := fmt.Sprintf(
        "postgres://%s:%s@%s:%s/%s?sslmode=%s",
        os.Getenv("POSTGRES_USER"),
        os.Getenv("POSTGRES_PASSWORD"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("POSTGRES_DB"),
        os.Getenv("DB_SSLMODE"),
    )

    // Création du pool de connexions
    pool, err := pgxpool.New(context.Background(), dsn)
    if err != nil {
        return fmt.Errorf("❌ Erreur de connexion à PostgreSQL : %w", err)
    }

    // Vérification de la connexion
    if err := pool.Ping(context.Background()); err != nil {
        return fmt.Errorf("❌ Impossible de se connecter à PostgreSQL : %w", err)
    }

    DB = pool
    log.Println("✅ Connexion réussie à PostgreSQL (pgxpool) !")
    return nil
}

// CloseDB ferme le pool de connexions PostgreSQL
func CloseDB() {
    if DB != nil {
        DB.Close()
        log.Println("✅ Connexion PostgreSQL fermée.")
    }
}
