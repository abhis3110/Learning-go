package main

import (
	//"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Open the database connection (replace with your own connection string)
	dsn := "user:password@tcp(127.0.0.1:3306)/yourdb"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to open the database: %v", err)
	}
	defer db.Close()

	// Perform a transaction
	err = createOrderTransaction(db)
	if err != nil {
		log.Printf("transaction failed: %v", err)
	} else {
		log.Println("transaction successful")
	}
}

func createOrderTransaction(db *sql.DB) error {
	// Begin a new transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	// In case of error, ensure the transaction is rolled back
	defer func() {
		if err != nil {
			tx.Rollback()
			log.Println("transaction rolled back")
		}
	}()

	// Insert into the orders table (replace with your own query)
	orderID, err := insertOrder(tx, "Cash", 100, 10, 110, 1)
	if err != nil {
		return fmt.Errorf("failed to insert order: %w", err)
	}

	// Insert associated order items (replace with your own query)
	err = insertOrderItem(tx, orderID, "Product A", 2, "image.jpg", 50, 101)
	if err != nil {
		return fmt.Errorf("failed to insert order item: %w", err)
	}

	err = insertOrderItem(tx, orderID, "Product B", 1, "image2.jpg", 60, 102)
	if err != nil {
		return fmt.Errorf("failed to insert order item: %w", err)
	}

	// Commit the transaction if all operations are successful
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	log.Println("transaction committed")
	return nil
}

func insertOrder(tx *sql.Tx, paymentMethod string, taxPrice, shippingPrice, totalPrice float64, userID int) (int64, error) {
	query := `INSERT INTO orders (payment_method, tax_price, shipping_price, total_price, user_id)
			  VALUES (?, ?, ?, ?, ?)`
	res, err := tx.Exec(query, paymentMethod, taxPrice, shippingPrice, totalPrice, userID)
	if err != nil {
		return 0, fmt.Errorf("failed to insert order: %w", err)
	}

	orderID, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert ID: %w", err)
	}

	return orderID, nil
}

func insertOrderItem(tx *sql.Tx, orderID int64, name string, quantity int, image string, price float64, productID int64) error {
	query := `INSERT INTO order_items (order_id, name, quantity, image, price, product_id)
			  VALUES (?, ?, ?, ?, ?, ?)`
	_, err := tx.Exec(query, orderID, name, quantity, image, price, productID)
	if err != nil {
		return fmt.Errorf("failed to insert order item: %w", err)
	}

	return nil
}
