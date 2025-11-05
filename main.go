package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
}

type User struct {
	ID    int    `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
}

func main() {
	router := gin.Default()

	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})

	if err != nil {
		panic("Unable to connect to the database")
	}

	db.AutoMigrate(&Todo{})

	router.GET("/todos", func(c *gin.Context) {
		var todos []Todo
		db.Find(&todos)
		c.JSON(200, todos)
	})

	router.POST("/todos", func(c *gin.Context) {
		var todo Todo
		if err := c.BindJSON(&todo); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		db.Create(&todo)
		c.JSON(201, todo)
	})

	router.GET("/todos/:id", func(c *gin.Context) {
		var todo Todo
		todoId := c.Param("id")
		if err := db.Where("id = ?", todoId).First(&todo).Error; err != nil {
			c.JSON(404, gin.H{"error": "Todo not found"})
			return
		}
		c.JSON(200, todo)
	})

	router.PUT("/todos/:id", func(c *gin.Context) {
		var todo Todo
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
	})

	router.DELETE("/todos/:id", func(c *gin.Context) {
		var todo Todo
		todoId := c.Param("id")
		if err := db.Where("id = ?", todoId).First(&todo).Error; err != nil {
			c.JSON(404, gin.H{"error": "Todo not found"})
			return
		}
		db.Delete(&todo)
		c.JSON(200, gin.H{"message": "Todo deleted"})
	})


	router.POST("/json", func(c *gin.Context){
		var user User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, user)
	})

	router.POST("/form", func(ctx *gin.Context) {
		var user User
		if err := ctx.Bind(&user); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, user)
	})

	router.GET("/search", func(ctx *gin.Context) {
		query := ctx.DefaultQuery("q", "default value")
		ctx.String(200, "Search query: "+query)
	})

	router.GET("/user/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.String(200, "User ID: "+id)
	})

		router.GET("/divide/:a/:b", func(c *gin.Context) {
		a := c.Param("a")
		b := c.Param("b")

		if b == "0" {
			c.JSON(400, gin.H{"error": "Cannot divide by zero"})
			return
		}
		floatA, _ := strconv.ParseFloat(a, 32)
		floatB, _ := strconv.ParseFloat(b, 32)
		c.JSON(200, gin.H{"result": floatA / floatB})
	})

	router.Run(":8080")

}
