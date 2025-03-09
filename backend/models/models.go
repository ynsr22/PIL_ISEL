package models

// Table des Catégories
type Categorie struct {
	ID                uint              `json:"id"`
	Nom               string            `json:"nom"`
	MoyensRoulants    []MoyenRoulant    `json:"moyens_roulants,omitempty"`
	AccessoireDefauts []AccessoireDefaut `json:"accessoire_defauts,omitempty"`
	Compatibilites    []Compatibilite   `json:"compatibilites,omitempty"`
}

// Table des Moyens Roulants
type MoyenRoulant struct {
	ID          uint    `json:"id"`
	Nom         string  `json:"nom"`
	CategorieID uint    `json:"categorie_id"`
	Roues       int     `json:"roues"`
	Emplacement string  `json:"emplacement"`
	TypeBase    string  `json:"type_base"`
	Taille      string  `json:"taille"`
	Departement string  `json:"departement"`
	Image       string  `json:"image"`
	Prix        float64 `json:"prix"` // numeric(10,2) => float64
}

// Table des Accessoires
type Accessoire struct {
	ID    uint    `json:"id"`
	Nom   string  `json:"nom"`
	Image string  `json:"image"`
	Prix  float64 `json:"prix"` // numeric(10,2) => float64
}

// Table des Compatibilités
type Compatibilite struct {
	ID          uint `json:"id"`
	CategorieID uint `json:"categorie_id"`
	AccessoireID uint `json:"accessoire_id"`
}

// Table des Accessoires par défaut
type AccessoireDefaut struct {
	ID          uint `json:"id"`
	CategorieID uint `json:"categorie_id"`
	AccessoireID uint `json:"accessoire_id"`
}
