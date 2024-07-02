package controllers

import (
	"net/http"

	"taskday7.4/configs"
	"taskday7.4/internal/models"

	"github.com/labstack/echo/v4"
)

// Handler untuk menambahkan todo
func AddTodo(c echo.Context) error {
	todo := new(models.Todo)
	if err := c.Bind(todo); err != nil {
		return err
	}
	configs.DB.Create(&todo)
	return c.JSON(http.StatusOK, todo)
}

// Handler untuk mendapatkan semua todo
func GetTodos(c echo.Context) error {
	var todos []models.Todo
	configs.DB.Find(&todos)
	return c.JSON(http.StatusOK, todos)
}

// Handler untuk mendapatkan todo berdasarkan ID
func GetTodoByID(c echo.Context) error {
	id := c.Param("todoId")
	var todo models.Todo
	if err := configs.DB.First(&todo, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Todo not found",
		})
	}
	return c.JSON(http.StatusOK, todo)
}

// Handler untuk memperbarui todo berdasarkan ID
func UpdateTodoByID(c echo.Context) error {
	id := c.Param("todoId")
	var todo models.Todo
	if err := configs.DB.First(&todo, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Todo not found",
		})
	}
	if err := c.Bind(&todo); err != nil {
		return err
	}
	configs.DB.Save(&todo)
	return c.JSON(http.StatusOK, todo)
}

// Handler untuk menghapus todo berdasarkan ID
func DeleteTodoByID(c echo.Context) error {
	id := c.Param("todoId")
	if err := configs.DB.Delete(&models.Todo{}, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Todo not found",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Todo successfully deleted",
	})
}
