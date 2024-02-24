package models

import (
	"fmt"
	"github.com/maxzhovtyj/card-validator/pkg/card"
	"time"
)

var _ card.Card = (*Card)(nil)

type Card struct {
	Number          string `json:"number"`
	ExpirationMonth string `json:"expirationMonth"`
	ExpirationYear  int    `json:"expirationYear"`
}

var cardMonth = map[string]time.Month{
	"01": time.January,
	"02": time.February,
	"03": time.March,
	"04": time.April,
	"05": time.May,
	"06": time.June,
	"07": time.July,
	"08": time.August,
	"09": time.September,
	"10": time.October,
	"11": time.November,
	"12": time.December,
}

func ParseExpirationMonth(month string) (time.Month, error) {
	m, ok := cardMonth[month]
	if !ok {
		return 0, fmt.Errorf("unknown month")
	}

	return m, nil
}

func (c Card) GetNumber() string {
	return c.Number
}

func (c Card) GetExpirationMonth() time.Month {
	return cardMonth[c.ExpirationMonth]
}

func (c Card) GetExpirationYear() int {
	return c.ExpirationYear
}
