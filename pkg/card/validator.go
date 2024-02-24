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
	GetExpirationMonth() time.Month
	GetExpirationYear() int
}

type Validator interface {
	Valid(card Card) error
}

type validator struct {
}

func NewValidator() Validator {
	return &validator{}
}

func (v *validator) Valid(card Card) error {
	currTime := time.Now()
	currTime.Month()

	if card.GetExpirationYear() < currTime.Year() {
		return ErrCardExpired
	} else if card.GetExpirationYear() == currTime.Year() && card.GetExpirationMonth() > currTime.Month() {
		return ErrCardExpired
	}

	if card.GetExpirationMonth() > time.December || card.GetExpirationMonth() < time.January {
		return ErrInvalidExpirationMonth
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
