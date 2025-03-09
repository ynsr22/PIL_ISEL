package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"backend/database"
	"backend/models"
)

// GetAccessoires récupère et renvoie la liste des accessoires.
func GetAccessoires(w http.ResponseWriter, r *http.Request) {
	query := "SELECT id, nom  FROM liste_accessoire"
	rows, err := database.DB.Query(r.Context(), query)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des accessoires", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var accessoires []models.Accessoire
	for rows.Next() {
		var a models.Accessoire
		if err := rows.Scan(&a.ID, &a.Nom); err != nil {
			http.Error(w, "Erreur lors du scan des données", http.StatusInternalServerError)
			return
		}
		accessoires = append(accessoires, a)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accessoires)
}

// GetAccessoire récupère un accessoire en fonction de son ID.
func GetAccessoire(w http.ResponseWriter, r *http.Request) {
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

	var accessoire models.Accessoire
	query := "SELECT id, nom  FROM liste_accessoire WHERE id = $1"
	err = database.DB.QueryRow(r.Context(), query, id).Scan(&accessoire.ID, &accessoire.Nom)

	if err != nil {
		http.Error(w, "Accessoire non trouvé", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accessoire)
}

// GetAccessoiresDefauts récupère les accessoires par défaut d'une catégorie.
func GetAccessoiresDefauts(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "ID de catégorie manquant", http.StatusBadRequest)
		return
	}

	categorieID, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "ID de catégorie invalide", http.StatusBadRequest)
		return
	}

	query := "SELECT la.id, la.nom, la.image, la.prix FROM liste_accessoire la JOIN accessoire_defaut ad ON la.id = ad.accessoire_id WHERE ad.categorie_id = $1"
	rows, err := database.DB.Query(r.Context(), query, categorieID)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des accessoires par défaut", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	
	var accessoires []models.Accessoire
	for rows.Next() {
		var a models.Accessoire
		if err := rows.Scan(&a.ID, &a.Nom, &a.Image, &a.Prix); err != nil {
			http.Error(w, "Erreur lors du scan des données", http.StatusInternalServerError)
			return
		}
		accessoires = append(accessoires, a)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accessoires)
}

// GetAccessoiresForCategorie récupère les accessoires compatibles pour une catégorie.
func GetAccessoiresForCategorie(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "ID de catégorie manquant", http.StatusBadRequest)
		return
	}

	categorieID, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "ID de catégorie invalide", http.StatusBadRequest)
		return
	}

	// Récupérer les IDs des accessoires compatibles
	query := "SELECT accessoire_id FROM compatibilites WHERE categorie_id = $1"
	rows, err := database.DB.Query(r.Context(), query, categorieID)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des compatibilités", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var accessoireIDs []int
	for rows.Next() {
		var accessoireID int
		if err := rows.Scan(&accessoireID); err != nil {
			http.Error(w, "Erreur lors du scan des données", http.StatusInternalServerError)
			return
		}
		accessoireIDs = append(accessoireIDs, accessoireID)
	}

	if len(accessoireIDs) == 0 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]models.Accessoire{})
		return
	}

	// Récupérer les accessoires correspondants
	query = "SELECT id, nom, image, prix FROM liste_accessoire WHERE id = ANY($1)"
	rows, err = database.DB.Query(r.Context(), query, accessoireIDs)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des accessoires", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var accessoires []models.Accessoire
	for rows.Next() {
		var a models.Accessoire
		if err := rows.Scan(&a.ID, &a.Nom, &a.Image, &a.Prix); err != nil {
			http.Error(w, "Erreur lors du scan des données", http.StatusInternalServerError)
			return
		}
		accessoires = append(accessoires, a)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accessoires)
}
