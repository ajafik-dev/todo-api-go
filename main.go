package main

import (
	"strconv"

	"todo/controllers"
	"todo/middlewares"
	"todo/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	router := gin.Default()

	router.Use(middlewares.LoggerMiddleware())

	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})

	if err != nil {
		panic("Unable to connect to the database")
	}

	db.AutoMigrate(&models.Todo{})

	router.GET("/todos", func(c *gin.Context) {
		controllers.GetTodos(c, db)
	})

	router.POST("/todos", func(c *gin.Context) {
		controllers.CreateTodo(c, db)
	})

	router.GET("/todos/:id", func(c *gin.Context) {
		controllers.GetTodo(c, db)
	})

	router.PUT("/todos/:id", func(c *gin.Context) {
		controllers.UpdateTodo(c, db)
	})

	router.DELETE("/todos/:id", func(c *gin.Context) {
		controllers.DeleteTodo(c, db)
	})

	router.POST("/json", func(c *gin.Context) {
		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, user)
	})

	router.POST("/form", func(ctx *gin.Context) {
		var user models.User
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
