package service

import (
    "github.com/gin-gonic/gin"
    "net/http"
	"strconv"
	"time"
	"orders/entities"
	"orders/dto"
    "orders/repositories"
    "github.com/jinzhu/gorm"
	"github.com/go-playground/validator/v10"
)



//CartService represents the service for cart-related operations.
type CartService struct {
    cartRepo *repositories.CartRepository
}

func NewCartService(db *gorm.DB) *CartService {
    return &CartService{
        cartRepo: repositories.NewCartRepository(db), // Initialize the cart repository with the database connection.
    }
}



func GetAllCards(c *gin.Context) {

   db := c.MustGet("db").(*gorm.DB)
    // Create a card repository using the database connection.
   cartRepo := repositories.NewCartRepository(db)

    // Fetch all cards from the database using the repository.
    cards, err := cartRepo.GetAll() // Implement the GetAll method in your repository.

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to fetch cards",
        })
        return
    }

    // Respond with JSON containing the card data.
    c.JSON(http.StatusOK, gin.H{
        "cards": cards,
    })
}


func GetCardByID(c *gin.Context) {
    // Extract cardId from the request URL
    cardID := c.Param("cardId")

    // Validate cardID (You can use additional validation logic as needed)
    if cardID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "cardId is required"})
        return
    }

    // Convert cardID to a suitable data type (e.g., uint) based on your database schema
    // Here, we assume cardID is a uint
    cardIDUint, err := strconv.ParseUint(cardID, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cardId"})
        return
    }

    // Initialize a database connection (you can use your own initialization logic)
    
	db := c.MustGet("db").(*gorm.DB)
    // Create a CartRepository instance to interact with the database
    cartRepo := repositories.NewCartRepository(db)

    // Retrieve the card by cardID
    card, err := cartRepo.GetByID(uint(cardIDUint))
    if err != nil {
        // Handle database errors
        if gorm.IsRecordNotFoundError(err) {
            c.JSON(http.StatusNotFound, gin.H{"error": "Card not found"})
            return
        }

        // Handle other database errors
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
        return
    }

    // Check if the card is nil (e.g., not found in the database)
    if card == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Card not found"})
        return
    }

    // Card found, return it as JSON
    c.JSON(http.StatusOK, card)
}



func CreateCard(c *gin.Context) {
    // Initialize a database connection (you can use your own initialization logic)
	db := c.MustGet("db").(*gorm.DB)

    // Parse the JSON request body into a CartDTO
    var cartDTO dto.CartDTO
    if err := c.ShouldBindJSON(&cartDTO); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Validate the CartDTO
	if err := cartDTO.Validate(); err != nil {
        // Extract and return custom error messages
        validationErrors := err.(validator.ValidationErrors)
        errorMessages := make(map[string]string)

        for _, fieldErr := range validationErrors {
            field := fieldErr.Field()
            message := fieldErr.Tag()
            errorMessages[field] = field + " " + message
        }

        c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
        return
    }


	currentTime := time.Now().Format(time.RFC3339)

    // Convert CartDTO to a Cart entity
    cart := entities.Cart{
        Status:      cartDTO.Status,
        ItemsCount:  cartDTO.ItemsCount,
        UserID:      cartDTO.UserID,
		CreatedAt:   currentTime, 
		UpdatedAt:   currentTime, 
        TotalAmount: cartDTO.TotalAmount,
    }

    // Create a CartRepository instance to interact with the database
    cartRepo := repositories.NewCartRepository(db)

    // Create the card in the database
    if err := cartRepo.Create(&cart); err != nil {
        // Handle database errors
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
        return
    }


    // Respond with the created card as JSON
    c.JSON(http.StatusCreated, cart)
}


// UpdateCard updates a card by its ID.
func UpdateCard(c *gin.Context) {
    // Initialize a database connection (you can use your own initialization logic)
	db := c.MustGet("db").(*gorm.DB)

	cardIDStr := c.Param("cardId")
    // Parse the card ID from the request URL
    cardID, err := strconv.ParseUint(cardIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid card ID"})
		return
	}

    // Check if the card exists in the database
    cardRepo := repositories.NewCartRepository(db)
    existingCard, err := cardRepo.GetByID(uint(cardID))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Card not found"})
        return
    }

    // Parse the JSON request body into a CartDTO
    var cartDTO dto.CartDTO
    if err := c.ShouldBindJSON(&cartDTO); err != nil {
        validationErrors := err.(validator.ValidationErrors)
        errorMessages := make(map[string]string)

        for _, fieldErr := range validationErrors {
            field := fieldErr.Field()
            message := fieldErr.Tag()
            errorMessages[field] = field + " " + message
        }

        c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
        return
    }

    // Validate the CartDTO
    if err := cartDTO.Validate(); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Update the existing card with the data from CartDTO
    existingCard.Status = cartDTO.Status
    existingCard.ItemsCount = cartDTO.ItemsCount
    existingCard.UserID = cartDTO.UserID
    existingCard.TotalAmount = cartDTO.TotalAmount

    // Save the updated card in the database
    if err := cardRepo.Update(existingCard); err != nil {
        // Handle database errors
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
        return
    }

    // Respond with the updated card as JSON
    c.JSON(http.StatusOK, existingCard)
}


// DeleteCard deletes a card by its ID.
func DeleteCard(c *gin.Context) {
    // Initialize a database connection (you can use your own initialization logic)
	db := c.MustGet("db").(*gorm.DB)

    // Parse the card ID from the request URL as a string
    cardIDStr := c.Param("cardId")

    // Convert the cardIDStr to a uint64
    cardID, err := strconv.ParseUint(cardIDStr, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid card ID"})
        return
    }

    // Check if the card exists in the database
    cardRepo := repositories.NewCartRepository(db)
    existingCard, err := cardRepo.GetByID(uint(cardID))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Card not found"})
        return
    }

    // Delete the card from the database
    if err := cardRepo.Delete(existingCard); err != nil {
        // Handle database errors
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
        return
    }

    // Respond with a success message
    c.JSON(http.StatusOK, gin.H{"message": "Card deleted successfully"})
}




// CreateCart creates a new cart in the database based on the provided cart data.
func (s *CartService) CreateCart(cart *entities.Cart) error {
    return s.cartRepo.Create(cart) // Call the Create method of the cart repository to create the cart.
}

// GetCartByID retrieves a cart by its unique identifier from the database.
func (s *CartService) GetCartByID(id uint) (*entities.Cart, error) {
    return s.cartRepo.GetByID(id) // Call the GetByID method of the cart repository to retrieve the cart.
}


