package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite" // registers sqlite driver for database/sql
)

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("sqlite", "./echo.db")
	if err != nil {
		log.Fatal("Błąd połączenia z bazą:", err)
	}

	createTables()
	seedProducts()
}

func createTables() {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS products (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT,
			price REAL NOT NULL,
			image_url TEXT
		);

		CREATE TABLE IF NOT EXISTS orders (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			total REAL NOT NULL,
			status TEXT NOT NULL DEFAULT 'pending',
			customer_name TEXT,
			customer_email TEXT,
			customer_address TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);

		CREATE TABLE IF NOT EXISTS order_items (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			order_id INTEGER,
			product_id INTEGER,
			name TEXT,
			quantity INTEGER,
			price REAL,
			FOREIGN KEY(order_id) REFERENCES orders(id)
		);
	`)
	if err != nil {
		log.Fatal("Błąd tworzenia tabel:", err)
	}
}

func seedProducts() {
	var count int
	DB.QueryRow("SELECT COUNT(*) FROM products").Scan(&count)
	if count > 0 {
		return
	}

	products := []struct {
		name        string
		description string
		price       float64
		imageURL    string
	}{
		{"SONY Playstation 5", "Konsola do gry najnowszej generacji", 2878.99, "ps5.jpg"},
		{"APPLE iPhone17e ", "Smartfon", 2999.99, "iphone.jpg"},
		{"Mysz bezprzewodowa", "Ergonomiczna mysz optyczna", 149.99, "mysz.jpg"},
		{"Klawiatura mechaniczna", "Switch Cherry MX Red", 299.99, "klas.jpg"},
		{"Monitor 27\"", "4K IPS 144Hz", 1899.99, "monitor.jpg"},
	}

	for _, p := range products {
		DB.Exec("INSERT INTO products (name, description, price, image_url) VALUES (?, ?, ?, ?)",
			p.name, p.description, p.price, p.imageURL)
	}

	log.Println("Produkty dodane do bazy")
}
