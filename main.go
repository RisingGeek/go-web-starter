package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/risinggeek/golang-starter/db"
	"github.com/risinggeek/golang-starter/routes"
	"gorm.io/gorm"
)

func main() {
	router := gin.Default()
	db := db.Setup()

	router.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
		ctx.Next()
	})

	router.GET("/test", func(ctx *gin.Context) {
		fmt.Println(ctx.MustGet("db").(*gorm.DB))
		ctx.JSON(200, gin.H{
			"message": "OK",
		})
	})

	routes.TodoRoute(router)

	router.Run(":8080")
}
