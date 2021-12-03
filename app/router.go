package app

import (
	"github.com/gin-gonic/gin"
	"github.com/nunonano/hacktiv-assignment2/controller"
)

func InitRouter(controller controller.DBConn) *gin.Engine {
	r := gin.Default()

	r.GET("/orders/", controller.GetOrders)
	r.GET("/orders/:id", controller.GetOrder)
	r.POST("/orders", controller.CreateOrders)
	r.PUT("/orders/:id", controller.UpdateOrders)
	r.DELETE("/orders/:id", controller.DeleteOrders)

	return r
}