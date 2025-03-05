package routes

import (
    "net/http"

    "github.com/gin-gonic/gin"

    "backend/controllers"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    // Middleware CORS
    r.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(http.StatusNoContent)
            return
        }
        c.Next()
    })

    // Route simple pour tester
    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Bienvenue sur l'API Moyens Logistiques !",
        })
    })

    // Routes Catégories
    r.GET("/categories", controllers.GetCategories)

    // Routes Moyens Roulants
    r.GET("/moyens", controllers.GetMoyens)
    r.GET("/moyens/:moyen_id", controllers.GetMoyen)
    r.GET("/moyens/:moyen_id/accessoires", controllers.GetAccessoiresForMoyen)

    // Routes Accessoires
    r.GET("/accessoires", controllers.GetAccessoires)
    r.GET("/accessoires/:accessoire_id", controllers.GetAccessoire)
    r.GET("/categories/:categorie_id/accessoires_defauts", controllers.GetAccessoiresDefauts)
    r.GET("/categories/:categorie_id/accessoires", controllers.GetAccessoiresForCategorie)

    return r
}
