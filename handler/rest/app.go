package rest

import (
	"assignment2/database"
	"assignment2/item"
	"assignment2/order"
	"fmt"

	"github.com/gin-gonic/gin"
)

const port = ":8080"

func StartApp() {
	database.InitializeDB()

	db := database.GetDB()

	orderRepository := order.NewRepository(db)
	itemRepository := item.NewRepository(db)
	itemService := item.NewService(itemRepository)
	orderService := order.NewService(orderRepository, itemRepository)
	orderHandler := NewOrder(orderService, itemService)

	route := gin.Default()
	v1 := route.Group("/api/v1")

	order := v1.Group("/orders")
	{
		order.GET("/", orderHandler.GetOrders)
		order.POST("/store", orderHandler.CreateOrder)
		order.PUT("/update/:orderId", orderHandler.UpdateOrder)
		order.DELETE("/delete/:orderId", orderHandler.DeleteOrder)
	}

	fmt.Println("Server running on PORT =>", port)
	route.Run(port)
}
