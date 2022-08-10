package memory_test

import (
	"testing"

	"github.com/ctirouzh/go-ddd/aggregate"
	"github.com/ctirouzh/go-ddd/domain/customer"
	"github.com/ctirouzh/go-ddd/domain/customer/memory"
	"github.com/google/uuid"
)

func TestMemory_AddCustomer(t *testing.T) {
	type testCase struct {
		name          string
		customer_name string
		expectedErr   error
	}
	testCases := []testCase{
		{
			name:          "add customer",
			customer_name: "Jhon Doe",
			expectedErr:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := memory.New()
			cstm, err := aggregate.NewCustomer(tc.customer_name)
			if err != nil {
				t.Fatal("aggregate->NewCustomer failed.", err)
			}
			err = repo.Add(cstm)
			if err != tc.expectedErr {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}

			found, err := repo.Get(cstm.GetID())
			if err != nil {
				t.Fatal("customer->memory->Add failed", err)
			}
			if found.GetID() != cstm.GetID() {
				t.Errorf("expected %v, got %v", cstm.GetID(), found.GetID())
			}
		})
	}

}

func TestMemory_GetCustomer(t *testing.T) {
	// Create a fake customer to add to repository
	fake, err := aggregate.NewCustomer("Percy")
	if err != nil {
		t.Fatal("aggregate->NewCustomer failed", err)
	}

	repo := memory.New()
	if err = repo.Add(fake); err != nil {
		t.Fatal("memory->Add failed", err)
	}

	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}
	testCases := []testCase{
		{
			name:        "wrong id",
			id:          uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			expectedErr: customer.ErrCustomerNotFound,
		}, {
			name:        "find customer by a valid id",
			id:          fake.GetID(),
			expectedErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := repo.Get(tc.id); err != tc.expectedErr {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}

}

func TestMemory_UpdateCustomer(t *testing.T) {
	names := []string{"Percy", "Jhon Doe"}
	customers := []aggregate.Customer{}
	for _, name := range names {
		cstm, err := aggregate.NewCustomer(name)
		if err != nil {
			t.Fatal("aggregate->NewCustomer failed", err)
		}
		customers = append(customers, cstm)
	}
	// Add first customer to memory repository
	repo := memory.New()
	if err := repo.Add(customers[0]); err != nil {
		t.Fatal("memory->Add failed", err)
	}
	// Set new name for the first customer
	customers[0].SetName("Mercy")

	type testCase struct {
		name        string
		customer    aggregate.Customer
		expectedErr error
	}
	testCases := []testCase{
		{
			name:        "update existing customer name",
			customer:    customers[0],
			expectedErr: nil,
		}, {
			name:        "cannot update an unknown customer",
			customer:    customers[1],
			expectedErr: customer.ErrCustomerNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if err := repo.Update(tc.customer); err != tc.expectedErr {
				t.Errorf("expected err %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
