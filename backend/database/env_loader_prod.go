//go:build production
package database

// LoadEnv ne fait rien en Prod, car on utilise les variables système
func LoadEnv() {
    // Aucune action à faire en prod
}