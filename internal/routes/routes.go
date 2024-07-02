package routes

import (
	"github.com/labstack/echo/v4"
	"taskday7.4/internal/controllers"
)

func RegisterRoutes(e *echo.Echo) {
	e.POST("/register", controllers.RegisterUser)
	e.POST("/login", controllers.LoginUser)
	e.GET("/todos", controllers.GetTodos)
	e.POST("/todos", controllers.AddTodo)
	e.GET("/todos/:todoId", controllers.GetTodoByID)
	e.PUT("/todos/:todoId", controllers.UpdateTodoByID)
	e.DELETE("/todos/:todoId", controllers.DeleteTodoByID)
}
