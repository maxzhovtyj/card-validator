package card

import (
	"testing"
)

var _ Card = (*card)(nil)

type card struct {
	number          string
	expirationMonth string
	expirationYear  int64
}

func (c *card) GetNumber() string {
	return c.number
}
func (c *card) GetExpirationMonth() string {
	return c.expirationMonth
}
func (c *card) GetExpirationYear() int64 {
	return c.expirationYear
}

func TestValid(t *testing.T) {
	testTable := []struct {
		name            string
		number          string
		expirationMonth string
		expirationYear  int64
		valid           bool
	}{
		{"valid", "4111111111111111", "12", 2028, true},
		{"valid", "4111111111111111", "12", 2024, true},
		{"invalid", "4111111111111111", "12", 2021, false},
		{"invalid1", "1111111111111", "12", 2028, false},
		{"invalid2", "411111ABC1111111", "12", 2028, false},
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
