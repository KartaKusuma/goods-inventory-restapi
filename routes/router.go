package routes

import (
	"goods-inventory-restapi/handlers"

	"github.com/gin-gonic/gin"
)

func Init() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	v1.GET("/goods", handlers.GetAllGoods)
	v1.GET("/goods/:id", handlers.GetGoodsByID)
	v1.POST("/goods", handlers.StoreGoods)
	v1.PUT("/goods/:id", handlers.UpdateGoods)
	v1.DELETE("/goods/:id", handlers.DeleteGoods)

	router.Run(":8080")
}
