package routes

import(
	"github.com/gin-gonic/gin"
	"restaurant-backend/controllers"
)

func FoodRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/foods", controllers.GetFoods)
	incomingRoutes.GET("/foods/:id", controllers.GetFood)
	incomingRoutes.POST("/foods", controllers.CreateFood)
	incomingRoutes.PATCH("/foods/:id", controllers.UpdateFood)
	// incomingRoutes.DELETE("/foods/:id", controllers.DeleteFood)
}