package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"

    "backend/database"
    "backend/models"
)

func GetCategories(c *gin.Context) {
    var categories []models.Categorie

    if err := database.DB.Find(&categories).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Erreur lors de la récupération des catégories",
        })
        return
    }

    c.JSON(http.StatusOK, categories)
}
