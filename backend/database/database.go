package database

import (
    "fmt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
    dsn := "host=db user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
    var err error

    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic(fmt.Sprintf("Erreur de connexion à la base de données : %v", err))
    }
}
