package card

import (
	"time"
)

var _ Card = (*card)(nil)

type Card interface {
	ExpirationYear() int
	ExpirationMonth() time.Month
	Number() string
	SetNumber(number string)
	SetExpirationYear(year int)
	SetExpirationMonth(month time.Month)
}

type card struct {
	number          string
	expirationYear  int
	expirationMonth time.Month
}

func (c *card) ExpirationYear() int {
	return c.expirationYear
}

func (c *card) ExpirationMonth() time.Month {
	return c.expirationMonth
}

func (c *card) Number() string {
	return c.number
}

func (c *card) SetNumber(number string) {
	c.number = number
}

func (c *card) SetExpirationYear(year int) {
	c.expirationYear = year
}

func (c *card) SetExpirationMonth(month time.Month) {
	c.expirationMonth = month
}

func New() Card {
	return &card{}
}
