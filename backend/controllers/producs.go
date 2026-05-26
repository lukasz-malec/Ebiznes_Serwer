package controllers

import (
	"net/http"

	"backend/db"
	"backend/models"

	"github.com/labstack/echo/v4"
)

func GetProducts(c echo.Context) error {
	rows, err := db.DB.Query("SELECT id, name, description, price, image_url FROM products")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Błąd pobierania produktów"})
	}
	defer rows.Close()

	products := []models.Product{}
	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.ImageURL)
		if err != nil {
			continue
		}
		products = append(products, p)
	}

	return c.JSON(http.StatusOK, products)
}