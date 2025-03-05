package models

// Table des Catégories
type Categorie struct {
	ID                uint              `json:"id" gorm:"primaryKey"`
	Nom               string            `json:"nom"`
	MoyensRoulants    []MoyenRoulant    `json:"moyens_roulants" gorm:"foreignKey:CategorieID"`
	AccessoireDefauts []AccessoireDefaut `json:"accessoire_defauts" gorm:"foreignKey:CategorieID"`
	Compatibilites    []Compatibilite   `json:"compatibilites" gorm:"foreignKey:CategorieID"`
}

// Surcharge du nom de la table si besoin (GORM trouvera normalement "categories" par défaut)
// func (Categorie) TableName() string {
// 	return "categories"
// }

// Table des Moyens Roulants
type MoyenRoulant struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Nom         string    `json:"nom"`
	CategorieID uint      `json:"categorie_id"`
	Categorie   Categorie `json:"categorie" gorm:"foreignKey:CategorieID"`
	Roues       int       `json:"roues"`
	Emplacement string    `json:"emplacement"`
	TypeBase    string    `json:"type_base"`
	Taille      string    `json:"taille"`
	Departement string    `json:"departement"`
	Image       string    `json:"image"`
	Prix        float64   `json:"prix"` // numeric(10,2) => float64
}

// Forcer GORM à utiliser "moyens_roulants" au lieu de la table par défaut
func (MoyenRoulant) TableName() string {
	return "moyens_roulants"
}

// Table des Accessoires
type Accessoire struct {
	ID    uint    `json:"id" gorm:"primaryKey"`
	Nom   string  `json:"nom"`
	Image string  `json:"image"`
	Prix  float64 `json:"prix"` // numeric(10,2) => float64
}

// Forcer GORM à utiliser "liste_accessoire" au lieu de la table par défaut
func (Accessoire) TableName() string {
	return "liste_accessoire"
}

// Table des Compatibilités
type Compatibilite struct {
	ID          uint `json:"id" gorm:"primaryKey"`
	CategorieID uint `json:"categorie_id"`
	AccessoireID uint `json:"accessoire_id"`
}

// (Optionnel) TableName si tu veux être explicite
func (Compatibilite) TableName() string {
	return "compatibilites"
}

// Table des Accessoires par défaut
type AccessoireDefaut struct {
	ID          uint `json:"id" gorm:"primaryKey"`
	CategorieID uint `json:"categorie_id"`
	AccessoireID uint `json:"accessoire_id"`
}

func (AccessoireDefaut) TableName() string {
	return "accessoire_defaut"
}
