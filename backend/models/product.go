package models

import (
	"errors"
)


type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageURL    string  `json:"image_url"`
}


// walidacja danych 
func (p Product) Validate() error {
	if p.Name == "" {
		return errors.New("nazwa produktu nie może być pusta")
	}
	if p.Price <= 0 {
		return errors.New("cena musi być większa od zera")
	}
	if p.ID <= 0 {
		return errors.New("ID musi być większe od zera")
	}
	return nil
}