package card

import (
	"errors"
	"time"
	"unicode"
)

var (
	ErrInvalidExpirationMonth = errors.New("invalid expiration month")
	ErrInvalidCardNumber      = errors.New("invalid card number")
	ErrCardExpired            = errors.New("card is expired")
)

type Card interface {
	GetNumber() string
	GetExpirationMonth() string
	GetExpirationYear() int64
}

type Validator interface {
	Valid(card Card) error
}

type validator struct {
}

func NewValidator() Validator {
	return &validator{}
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

func (v *validator) Valid(card Card) error {
	currTime := time.Now()
	currTime.Month()

	month, ok := cardMonth[card.GetExpirationMonth()]
	if !ok {
		return ErrInvalidExpirationMonth
	}

	if card.GetExpirationYear() < int64(currTime.Year()) {
		return ErrCardExpired
	} else if card.GetExpirationYear() == int64(currTime.Year()) && month < currTime.Month() {
		return ErrCardExpired
	}

	if len(card.GetNumber()) != 16 {
		return ErrInvalidCardNumber
	}

	for _, d := range card.GetNumber() {
		if !unicode.IsDigit(d) {
			return ErrInvalidCardNumber
		}
	}

	return nil
}
