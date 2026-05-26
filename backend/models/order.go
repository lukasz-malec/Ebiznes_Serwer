package models
import (
	"errors"
	"time"
)


type Customer struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

type OrderItem struct {
	ProductID int     `json:"product_id"`
	Name      string  `json:"name"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type Order struct {
	ID        int         `json:"id"`
	Items     []OrderItem `json:"items"`
	Total     float64     `json:"total"`
	Customer  Customer    `json:"customer"`
	Status    string      `json:"status"`
	CreatedAt time.Time   `json:"created_at"`
}


// walidacja danych

func (i OrderItem) Validate() error {
	if i.Name == "" {
		return errors.New("nazwa pozycji nie może być pusta")
	}
	if i.Price <= 0 {
		return errors.New("cena pozycji musi być większa od zera")
	}
	if i.Quantity <= 0 {
		return errors.New("ilość musi być większa od zera")
	}
	if i.ProductID <= 0 {
		return errors.New("product_id musi być większy od zera")
	}
	return nil
}

func (c Customer) Validate() error {
	if c.Name == "" {
		return errors.New("nazwa klienta nie może być pusta")
	}
	if c.Email == "" {
		return errors.New("email klienta nie może być pusty")
	}
	if c.Address == "" {
		return errors.New("adres klienta nie może być pusty")
	}
	return nil
}

func (o Order) Validate() error {
	if o.Total <= 0 {
		return errors.New("suma zamówienia musi być większa od zera")
	}
	if len(o.Items) == 0 {
		return errors.New("zamówienie musi mieć co najmniej jedną pozycję")
	}
	if err := o.Customer.Validate(); err != nil {
		return err
	}
	return nil
}