package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/risinggeek/golang-starter/controllers"
)

// TodoRoute export
func TodoRoute(router *gin.Engine) {
	todo := router.Group("/todos")
	{
		todo.GET("/", controller.GetTodos)
		todo.POST("/", controller.AddTodo)
		todo.PUT("/", controller.UpdateTodo)
		todo.DELETE("/", controller.DeleteTodo)
	}
}
