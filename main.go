package main

import(
	"os"
	"github.com/gin-gonic/gin"
	"restaurant-backend/database"
	"restaurant-backend/routes"
	"restaurant-backend/middlewares"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main(){
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middlewares.AuthMiddleware())

	routes.FoodRoutes(router)
	routes.OrderRoutes(router)
	routes.TableRoutes(router)
	routes.MenuRoutes(router) 
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)
}