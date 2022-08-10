package aggregate_test

import (
	"testing"

	"github.com/ctirouzh/go-ddd/aggregate"
)

func TestCustomerFactory(t *testing.T) {
	// Build our needed testcase data struct
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}
	// Create new test cases
	testCases := []testCase{
		{
			test:        "empty name validation",
			name:        "",
			expectedErr: aggregate.ErrInvalidPerson,
		}, {
			test:        "valid name recieved",
			name:        "Percy Bolmer",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		// Run Tests
		t.Run(tc.test, func(t *testing.T) {
			// Create a new customer
			_, err := aggregate.NewCustomer(tc.name)
			// Check if the error matches the expected error
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

		})
	}

}
