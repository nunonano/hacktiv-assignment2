package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nunonano/hacktiv-assignment2/model"
	"gorm.io/gorm"
)

var err error

func (conn *DBConn)  DeleteOrders(c *gin.Context) {
	id := c.Params.ByName("id")
	var orders model.Orders
	err := conn.DB.First(&orders, id).Error
	if err != nil{
		result := gin.H{
			"result":"data not found",
		}
		c.AbortWithStatusJSON(404, result)
	}
	err = conn.DB.Delete(&orders).Error
	if err != nil{
		result := gin.H{
			"result":"delete failed",
		}
		c.AbortWithStatusJSON(404, result)
	}
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

func (conn *DBConn) UpdateOrders(c *gin.Context) {

	var orders model.Orders
	id := c.Params.ByName("id")

	if err := conn.DB.Where("id = ?", id).First(&orders).Error; err != nil {
		result := gin.H{
			"result":"data not found",
		}
		c.AbortWithStatusJSON(404, result)
		fmt.Println(err)
		return

	}

	c.BindJSON(&orders)
	conn.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&orders)
	c.JSON(200, orders)

}

func (conn *DBConn) CreateOrders(c *gin.Context) {

	var orders model.Orders
	c.BindJSON(&orders)

	conn.DB.Create(&orders)
	c.JSON(200, orders)
}

func (conn *DBConn) GetOrder(c *gin.Context) {
	id := c.Params.ByName("id")
	var orders model.Orders
	if err := conn.DB.Preload("Items").Where("id = ?", id).First(&orders).Error; err != nil {
		result := gin.H{
			"result":"data not found",
		}
		c.AbortWithStatusJSON(404, result)
		fmt.Println(err)
	} else {
		c.JSON(200, orders)
	}
}

func (conn *DBConn) GetOrders(c *gin.Context) {
	var orders []model.Orders
	if err := conn.DB.Preload("Items").Find(&orders).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, orders)
	}
}