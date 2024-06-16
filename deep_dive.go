package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	// postgres driver
	_ "github.com/lib/pq"
)

type ShopModal interface {
	CountCustomers(time.Time) (int, error)
	CountSales(time.Time) (int, error)
}

type ShopDB struct {
	*sql.DB
}

func (sdb *ShopDB) CountCustomers(since time.Time) (int, error) {
	var count int
	err := sdb.QueryRow("SELECT count(*) FROM customers WHERE timestamp > $1", since).Scan(&count)
  fmt.Printf("Customers : %v \n",count)
	return count, err
}

func (sdb *ShopDB) CountSales(since time.Time) (int, error) {
	var count int
	err := sdb.QueryRow("SELECT count(*) FROM sales WHERE timestamp > $1", since).Scan(&count)
  fmt.Printf("Sales : %v \n",count)
	return count, err
}

func Shop() {
	// connect to postgres on local db deep no need for sslmode
	db, err := sql.Open("postgres", "postgres://haroonalbar:@localhost:5432/deep?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	s := &ShopDB{db}

	sr, err := calculateSalesPerCustomer(s)
	if err != nil {
		log.Fatal(err)
	}

  fmt.Printf("Sales per customer per day: %v \n",sr)

}

func calculateSalesPerCustomer(sm ShopModal) (string, error) {
	// time 1 day ago
	t := time.Now().Add(-24 * time.Hour)

	customers, err := sm.CountCustomers(t)
	if err != nil {
		return "", err
	}

	sales, err := sm.CountSales(t)
	if err != nil {
		return "", err
	}

	final := float64(sales) / float64(customers)
	return fmt.Sprintf("%.2f", final), nil

}
