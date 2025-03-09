package controllers

import (
	"encoding/json"
	"net/http"
	"backend/database"
	"backend/models"
	"strings"
	"strconv"
)

// GetCategories récupère et renvoie toutes les catégories depuis PostgreSQL.
func GetCategories(w http.ResponseWriter, r *http.Request) {
	query := "SELECT id, nom FROM categories"
	rows, err := database.DB.Query(r.Context(), query)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des catégories", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var categories []models.Categorie
	for rows.Next() {
		var c models.Categorie
		if err := rows.Scan(&c.ID, &c.Nom); err != nil {
			http.Error(w, "Erreur lors du scan des données", http.StatusInternalServerError)
			return
		}
		categories = append(categories, c)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)
}

func GetCategorie(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "ID manquant", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	query := "SELECT id, nom FROM categories WHERE id = $1"
	rows, err := database.DB.Query(r.Context(), query, id)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération de la catégorie", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var category models.Categorie
	for rows.Next() {
		var c models.Categorie
		if err := rows.Scan(&c.ID, &c.Nom); err != nil {
			http.Error(w, "Erreur lors du scan des données", http.StatusInternalServerError)
			return
		}
		category = c
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(category)
}