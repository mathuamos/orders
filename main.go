package main

import (
	"log"
	"orders/api"
	"orders/common"
	"github.com/gin-gonic/gin"
	 "github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	common.Configutils()
	common.Logutils()

	db := common.DbConn()
	defer db.Close()

	log.Println("Your application started")

	r := gin.Default()


	// Convert *sql.DB to *gorm.DB
	gormDB, err := gorm.Open("mysql", db)
	if err != nil {
		panic("Failed to convert *sql.DB to *gorm.DB")
	}

	// Use the DatabaseMiddleware to set the "db" key in the Gin context
	r.Use(DatabaseMiddleware(gormDB))
	
	// Set up your API routes
	v1 := r.Group("/v1")
	{
		api.SetupRoutes(v1) // Pass v1 as a *gin.RouterGroup
	}
	
    r.Run(":8080") // You can change the port as needed
}


// Define a middleware function to set the database connection in the context
func DatabaseMiddleware(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Set("db", db)
        c.Next()
    }
}
