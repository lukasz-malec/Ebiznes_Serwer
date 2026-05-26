package main

import (
	"net/http"

	"backend/controllers"
	"backend/db"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db.Init()

	e := echo.New()
	setupMiddleware(e)
	setupRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}


// ustawienie CORS
func setupMiddleware(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderContentType},
	}))
}


func setupRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	api := e.Group("/api")
	api.GET("/products", controllers.GetProducts)
	api.POST("/orders", controllers.CreateOrder)
}