// Package memory is a in-memory implementation of the customer repository
package memory

import (
	"sync"

	"github.com/ctirouzh/go-ddd/aggregate"
	"github.com/ctirouzh/go-ddd/domain/customer"
	"github.com/google/uuid"
)

// CustemorMemory fulfills the customer repository interface
/*
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
*/
type customerRepository struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

// New is a factory function to generate a new repository of customers
func New() *customerRepository {
	return &customerRepository{
		customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

// Get finds a customer by ID
func (r *customerRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	if customer, ok := r.customers[id]; ok {
		return customer, nil
	}
	return aggregate.Customer{}, customer.ErrCustomerNotFound
}

// Add will add a new customer to the repository
func (r *customerRepository) Add(c aggregate.Customer) error {
	// Make sure Customer isn't already in the repository
	if _, found := r.customers[c.GetID()]; found {
		return customer.ErrCustomerAlreadyExists
	}
	r.Lock()
	defer r.Unlock()
	r.customers[c.GetID()] = c
	return nil
}

// Update will replace an existing customer information with the new customer information
func (r *customerRepository) Update(c aggregate.Customer) error {
	// Make sure Customer is in the repository
	if _, found := r.customers[c.GetID()]; !found {
		return customer.ErrCustomerNotFound
	}
	r.Lock()
	defer r.Unlock()
	r.customers[c.GetID()] = c
	return nil
}
