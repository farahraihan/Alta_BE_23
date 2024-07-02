package main

import (
	"os"
	"restAPi/configs"
	"restAPi/internal/controllers/todos"
	"restAPi/internal/controllers/users"
	"restAPi/internal/models"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	cfg := configs.ImportSetting()
	db, _ := configs.ConnectDB(cfg)
	db.AutoMigrate(&models.User{}, &models.Todo{})
	um := models.NewUserModel(db)
	uc := users.NewUserController(um)

	tm := models.NewTodoModel(db)
	tc := todos.NewTodoController(tm)

	// Register
	e.POST("/register", uc.Register)
	e.POST("/login", uc.Login)

	t := e.Group("/todos")
	t.Use(echojwt.WithConfig(
		echojwt.Config{
			SigningKey:    []byte(os.Getenv("JWT_SIGNING_KEY")),
			SigningMethod: jwt.SigningMethodHS256.Name,
		},
	))
	t.GET("", tc.ShowMyTodo())
	t.POST("", tc.CreateTodo())
	// Update and Delete todo
	t.PUT("/:id", tc.UpdateTodo())
	t.DELETE("/:id", tc.DeleteTodo())

	e.GET("/hello", func(c echo.Context) error {
		return c.JSON(200, "hello world")
	}, echojwt.WithConfig(
		echojwt.Config{
			SigningKey:    []byte(os.Getenv("JWT_SIGNING_KEY")),
			SigningMethod: jwt.SigningMethodHS256.Name,
		},
	))
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())

	e.Logger.Error(e.Start(":8000"))
}
