package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"backend/database"
	"backend/models"
)

// GetMoyens renvoie la liste des moyens roulants.
func GetMoyens(w http.ResponseWriter, r *http.Request) {
	query := "SELECT id, nom, categorie_id, roues, emplacement, type_base, taille, departement, image, prix FROM moyens_roulants"
	rows, err := database.DB.Query(r.Context(), query)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des moyens roulants", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var moyens []models.MoyenRoulant
	for rows.Next() {
		var m models.MoyenRoulant
		if err := rows.Scan(&m.ID, &m.Nom, &m.CategorieID, &m.Roues, &m.Emplacement, &m.TypeBase, &m.Taille, &m.Departement, &m.Image, &m.Prix); err != nil {
			http.Error(w, "Erreur lors du scan des données", http.StatusInternalServerError)
			return
		}
		moyens = append(moyens, m)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(moyens)
}

// GetMoyen renvoie un moyen roulant en fonction de son ID.
func GetMoyen(w http.ResponseWriter, r *http.Request) {
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

	var moyen models.MoyenRoulant
	query := "SELECT id, nom, categorie_id, roues, emplacement, type_base, taille, departement, image, prix FROM moyens_roulants WHERE id = $1"
	err = database.DB.QueryRow(r.Context(), query, id).Scan(
		&moyen.ID, &moyen.Nom, &moyen.CategorieID, &moyen.Roues,
		&moyen.Emplacement, &moyen.TypeBase, &moyen.Taille,
		&moyen.Departement, &moyen.Image, &moyen.Prix,
	)
	if err != nil {
		http.Error(w, "Moyen roulant non trouvé", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(moyen)
}

// GetAccessoiresForMoyen renvoie les accessoires associés à un moyen roulant.
func GetAccessoiresForMoyen(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 || parts[3] != "accessoires" {
		http.Error(w, "Chemin invalide", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	// Vérification de l'existence du moyen
	var categorieID int
	query := "SELECT categorie_id FROM moyens_roulants WHERE id = $1"
	err = database.DB.QueryRow(r.Context(), query, id).Scan(&categorieID)
	if err != nil {
		http.Error(w, "Moyen roulant non trouvé", http.StatusNotFound)
		return
	}

	// Récupération des compatibilités liées à la catégorie du moyen
	query = "SELECT accessoire_id FROM compatibilites WHERE categorie_id = $1"
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

	// Récupération des accessoires correspondants
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
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessoires)
}
