package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/risinggeek/golang-starter/models"
	"gorm.io/gorm"
)

// GetTodo export
func GetTodo(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Task 1",
	})
}

// AddTodo export
func AddTodo(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)

	var todo models.Todo
	error := ctx.ShouldBindJSON(&todo)
	if error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": error.Error(),
		})
		return
	}

	task := models.Todo{ID: uuid.New(), Message: todo.Message}
	db.Create(&task)

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Todo added successfully",
		"data":    task,
	})
}

// GetTodos export
func GetTodos(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)

	var todos []models.Todo
	db.Find(&todos)
	ctx.JSON(http.StatusOK, gin.H{
		"data": todos,
	})
}

// UpdateTodo export
func UpdateTodo(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)

	type Task struct {
		ID      uuid.UUID
		Message string
	}
	var task Task
	error := ctx.ShouldBindJSON(&task)
	if error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": error.Error(),
		})
		return
	}

	db.Model(&models.Todo{}).Where("ID=?", task.ID).Update("message", task.Message)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Task updated successfully",
	})
}

// DeleteTodo export
func DeleteTodo(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)

	type Task struct {
		ID uuid.UUID
	}

	var task Task
	error := ctx.ShouldBindJSON(&task)
	if error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": error.Error(),
		})
		return
	}

	db.Delete(&models.Todo{}, task.ID)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Task deleted successfully",
	})
}
