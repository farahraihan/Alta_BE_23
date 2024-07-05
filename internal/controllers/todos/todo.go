package todos

import (
	"jwt/internal/helper"
	"jwt/internal/models"
	"jwt/internal/utils"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type TodoController struct {
	model *models.TodoModel
}

func NewTodoController(m *models.TodoModel) *TodoController {
	return &TodoController{
		model: m,
	}
}

func (tc *TodoController) CreateTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := utils.DecodeToken(c.Get("user").(*jwt.Token))
		if err != nil {
			return c.JSON(401, helper.ResponseFormat(401, "unauthorized", nil))
		}

		var input TodoRequest

		err = c.Bind(&input)
		if err != nil {
			return c.JSON(400, helper.ResponseFormat(400, "input error", nil))
		}

		err = tc.model.InsertTodo(ToModelTodo(input, userID))

		if err != nil {
			return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
		}

		return c.JSON(201, helper.ResponseFormat(201, "success insert todo", nil))
	}
}

func (tc *TodoController) ShowMyTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := utils.DecodeToken(c.Get("user").(*jwt.Token))
		if err != nil {
			return c.JSON(401, helper.ResponseFormat(401, "unauthorized", nil))
		}

		result, err := tc.model.GetTodoByUser(userID)
		if err != nil {
			return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
		}

		return c.JSON(200, helper.ResponseFormat(200, "success get my todo", result))
	}
}

func (tc *TodoController) UpdateTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		todoIDStr := c.Param("id")
		todoID, err := strconv.ParseUint(todoIDStr, 10, 64)
		if err != nil {
			return c.JSON(400, helper.ResponseFormat(400, "invalid todo ID", nil))
		}

		var updatedTodo models.Todo

		if err := c.Bind(&updatedTodo); err != nil {
			return c.JSON(400, helper.ResponseFormat(400, "input error", nil))
		}

		err = tc.model.UpdateTodoByID(uint(todoID), updatedTodo)

		if err != nil {
			return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
		}

		return c.JSON(200, helper.ResponseFormat(200, "success update todo", nil))
	}
}

func (tc *TodoController) DeleteTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		todoIDStr := c.Param("id")
		todoID, err := strconv.ParseUint(todoIDStr, 10, 64)
		if err != nil {
			return c.JSON(400, helper.ResponseFormat(400, "invalid todo ID", nil))
		}

		err = tc.model.DeleteTodoByID(uint(todoID))

		if err != nil {
			return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
		}

		return c.JSON(200, helper.ResponseFormat(200, "success delete todo", nil))
	}
}
