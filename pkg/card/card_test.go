package card

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	testTable := []struct {
		name            string
		number          string
		expirationMonth time.Month
		expirationYear  int
	}{
		{"card", "4111111111111111", time.December, 2028},
		{"card1", "4111111111111111", time.January, 2021},
		{"card2", "1111111111111", time.October, 2028},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := New()
			c.SetNumber(testCase.number)
			c.SetExpirationYear(testCase.expirationYear)
			c.SetExpirationMonth(testCase.expirationMonth)

			if c.Number() != testCase.number {
				t.Errorf("card number: expected %s got %s", testCase.number, c.Number())
			}

			if c.ExpirationMonth() != testCase.expirationMonth {
				t.Errorf("expiration month: expected %s got %s", testCase.expirationMonth.String(), c.ExpirationMonth().String())
			}

			if c.ExpirationYear() != testCase.expirationYear {
				t.Errorf("expiration year: expected %d got %d", testCase.expirationYear, c.ExpirationYear())
			}
		})
	}
}
