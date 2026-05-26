package controllers

import (
	"net/http"
	"database/sql"

	"backend/db"
	"backend/models"

	"github.com/labstack/echo/v4"
)

func CreateOrder(c echo.Context) error {
    order := new(models.Order)
    if err := c.Bind(order); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Nieprawidłowe dane"})
    }

    tx, err := db.DB.Begin()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Błąd transakcji"})
    }
    defer tx.Rollback()

    orderID, err := zapiszZamowienie(tx, order)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Błąd zapisu zamówienia"})
    }

    if err := zapiszPozycje(tx, orderID, order.Items); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Błąd zapisu pozycji"})
    }

    if err := tx.Commit(); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Błąd zatwierdzenia"})
    }

    return c.JSON(http.StatusCreated, map[string]interface{}{
        "order_id": orderID,
        "status":   "confirmed",
    })
}

func zapiszZamowienie(tx *sql.Tx, order *models.Order) (int64, error) {
    result, err := tx.Exec(
        "INSERT INTO orders (total, status, customer_name, customer_email, customer_address) VALUES (?, ?, ?, ?, ?)",
        order.Total, "pending", order.Customer.Name, order.Customer.Email, order.Customer.Address,
    )
    if err != nil {
        return 0, err
    }
    return result.LastInsertId()
}


func zapiszPozycje(tx *sql.Tx, orderID int64, items []models.OrderItem) error {
    for _, item := range items {
        _, err := tx.Exec(
            "INSERT INTO order_items (order_id, product_id, name, quantity, price) VALUES (?, ?, ?, ?, ?)",
            orderID, item.ProductID, item.Name, item.Quantity, item.Price,
        )
        if err != nil {
            return err
        }
    }
    return nil
}