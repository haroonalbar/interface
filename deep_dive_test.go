package main

import (
	"testing"
	"time"
)

// MockDb satisfies ShopModal interface
type MockDb struct{}

func (m *MockDb) CountCustomers(_ time.Time) (int, error) {
	return 1000, nil
}

func (m *MockDb) CountSales(_ time.Time) (int, error) {
	return 333, nil
}

func TestSalesPerCustomer(t *testing.T) {
	m := &MockDb{}

	sr, err := calculateSalesPerCustomer(m)
	if err != nil {
		t.Fatal(err)
	}

	exp := "0.33"
	if sr == exp {
		t.Fatalf("We got %v, expected is %v", sr, exp)
	}
}
