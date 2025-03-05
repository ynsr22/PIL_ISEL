package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"

    "backend/database"
    "backend/models"
)

func GetAccessoires(c *gin.Context) {
    var accessoires []models.Accessoire

    if err := database.DB.Find(&accessoires).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Erreur lors de la récupération des accessoires",
        })
        return
    }

    c.JSON(http.StatusOK, accessoires)
}

func GetAccessoire(c *gin.Context) {
    var accessoire models.Accessoire
    id := c.Param("accessoire_id")

    if err := database.DB.First(&accessoire, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "error": "Accessoire non trouvé",
        })
        return
    }

    c.JSON(http.StatusOK, accessoire)
}

func GetAccessoiresDefauts(c *gin.Context) {
    var (
        accessoiresDefauts []models.AccessoireDefaut
        accessoires        []models.Accessoire
    )
    categorieID := c.Param("categorie_id")

    // On récupère d'abord les enregistrements de la table accessoire_defauts
    database.DB.Where("categorie_id = ?", categorieID).Find(&accessoiresDefauts)

    var accessoireIDs []uint
    for _, a := range accessoiresDefauts {
        accessoireIDs = append(accessoireIDs, a.AccessoireID)
    }

    // Puis on récupère les accessoires concernés
    database.DB.Where("id IN ?", accessoireIDs).Find(&accessoires)

    c.JSON(http.StatusOK, accessoires)
}

func GetAccessoiresForCategorie(c *gin.Context) {
    var (
        compatibilites []models.Compatibilite
        accessoires    []models.Accessoire
    )
    categorieID := c.Param("categorie_id")

    // Récupération des compatibilités de la catégorie
    database.DB.Where("categorie_id = ?", categorieID).Find(&compatibilites)

    // Extraction des IDs d'accessoires
    var accessoireIDs []uint
    for _, comp := range compatibilites {
        accessoireIDs = append(accessoireIDs, comp.AccessoireID)
    }

    // Récupération des accessoires correspondants
    database.DB.Where("id IN ?", accessoireIDs).Find(&accessoires)

    c.JSON(http.StatusOK, accessoires)
}
