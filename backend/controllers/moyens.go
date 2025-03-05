package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"

    "backend/database"
    "backend/models"
)

func GetMoyens(c *gin.Context) {
    var moyens []models.MoyenRoulant

    if err := database.DB.Find(&moyens).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Erreur lors de la récupération des moyens roulants",
        })
        return
    }
    c.JSON(http.StatusOK, moyens)
}

func GetMoyen(c *gin.Context) {
    var moyen models.MoyenRoulant
    id := c.Param("moyen_id")

    if err := database.DB.First(&moyen, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "error": "Moyen roulant non trouvé",
        })
        return
    }
    c.JSON(http.StatusOK, moyen)
}

func GetAccessoiresForMoyen(c *gin.Context) {
    var (
        moyen         models.MoyenRoulant
        compatibilites []models.Compatibilite
        accessoires   []models.Accessoire
    )
    id := c.Param("moyen_id")

    // Vérification de l'existence du moyen
    if err := database.DB.First(&moyen, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "error": "Moyen roulant non trouvé",
        })
        return
    }

    // Récupération des compatibilités liées à la catégorie du moyen
    database.DB.Where("categorie_id = ?", moyen.CategorieID).Find(&compatibilites)

    // Extraction des IDs des accessoires
    var accessoireIDs []uint
    for _, comp := range compatibilites {
        accessoireIDs = append(accessoireIDs, comp.AccessoireID)
    }

    // Récupération des accessoires correspondants
    database.DB.Where("id IN ?", accessoireIDs).Find(&accessoires)

    c.JSON(http.StatusOK, accessoires)
}
