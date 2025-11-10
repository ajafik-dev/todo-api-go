package controllers

import (
	"todo/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetTodos(c *gin.Context, db *gorm.DB) {
	var todos []models.Todo
	db.Find(&todos)
	c.JSON(200, todos)
}

func CreateTodo(c *gin.Context, db *gorm.DB) {
	var todo models.Todo
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	db.Create(&todo)
	c.JSON(201, todo)
}

func GetTodo(c *gin.Context, db *gorm.DB) {
	var todo models.Todo
	todoId := c.Param("id")
	if err := db.Where("id = ?", todoId).First(&todo).Error; err != nil {
		c.JSON(404, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(200, todo)
}

func UpdateTodo(c *gin.Context, db *gorm.DB) {
	var todo models.Todo
	todoId := c.Param("id")
	if err := db.Where("id = ?", todoId).First(&todo).Error; err != nil {
		c.JSON(404, gin.H{"error": "Todo not found"})
		return
	}
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	db.Save(&todo)
	c.JSON(200, todo)
}

func DeleteTodo(c *gin.Context, db *gorm.DB) {
	var todo models.Todo
	todoId := c.Param("id")
	if err := db.Where("id = ?", todoId).First(&todo).Error; err != nil {
		c.JSON(404, gin.H{"error": "Todo not found"})
		return
	}
	db.Delete(&todo)
	c.JSON(200, gin.H{"message": "Todo deleted"})
}
