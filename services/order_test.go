package services

import (
	"testing"

	"github.com/ctirouzh/go-ddd/aggregate"
	"github.com/google/uuid"
)

func init_products(t *testing.T) []aggregate.Product {
	beer, err := aggregate.NewProduct("Beer", "Healthy Beverage", 1.99)
	if err != nil {
		t.Error(err)
	}
	peenuts, err := aggregate.NewProduct("Peenuts", "Healthy Snacks", 0.99)
	if err != nil {
		t.Error(err)
	}
	wine, err := aggregate.NewProduct("Wine", "Healthy Snacks", 0.99)
	if err != nil {
		t.Error(err)
	}
	products := []aggregate.Product{
		beer, peenuts, wine,
	}
	return products
}
func TestOrder_NewOrderService(t *testing.T) {
	// Create a few products to insert into in memory repo
	products := init_products(t)

	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)

	if err != nil {
		t.Error(err)
	}

	// Add Customer
	cstm, err := aggregate.NewCustomer("Percy")
	if err != nil {
		t.Error(err)
	}

	err = os.customers.Add(cstm)
	if err != nil {
		t.Error(err)
	}

	// Perform Order for one beer
	order := []uuid.UUID{
		products[0].GetID(),
	}

	_, err = os.CreateOrder(cstm.GetID(), order)

	if err != nil {
		t.Error(err)
	}

}
