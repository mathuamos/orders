package api

import (
	"orders/service" 
	"github.com/gin-gonic/gin"

)

func SetupRoutes(r *gin.RouterGroup) {
	r.GET("/api/cards/all", service.GetAllCards)

    // Example route: Get specific card by ID
    r.GET("/api/cards/:cardId", service.GetCardByID)

    // Example route: Create a new card
    r.POST("/api/cards/create", service.CreateCard)

    // Example route: Update a card
    r.PUT("/api/cards/update/:cardId", service.UpdateCard)

    // Example route: Delete a card
    r.DELETE("/api/cards/delete/:cardId", service.DeleteCard)
}
