package card

import (
	"testing"
	"time"
)

var _ Card = (*card)(nil)

type card struct {
	number          string
	expirationMonth time.Month
	expirationYear  int
}

func (c *card) GetNumber() string {
	return c.number
}
func (c *card) GetExpirationMonth() time.Month {
	return c.expirationMonth
}
func (c *card) GetExpirationYear() int {
	return c.expirationYear
}

func TestValid(t *testing.T) {
	testTable := []struct {
		name            string
		number          string
		expirationMonth time.Month
		expirationYear  int
		valid           bool
	}{
		{"valid", "4111111111111111", time.December, 2028, true},
		{"invalid", "4111111111111111", time.January, 2021, false},
		{"invalid1", "1111111111111", time.October, 2028, false},
		{"invalid2", "411111ABC1111111", time.October, 2028, false},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := &card{
				number:          testCase.number,
				expirationMonth: testCase.expirationMonth,
				expirationYear:  testCase.expirationYear,
			}

			v := NewValidator()

			err := v.Valid(c)
			if testCase.valid != (err == nil) {
				t.Errorf("validation, expected %v got %v", testCase.valid, err)
			}
		})
	}
}
