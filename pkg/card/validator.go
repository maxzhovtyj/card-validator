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

func Valid(card Card) error {
	currTime := time.Now()
	currTime.Month()

	if card.ExpirationYear() < currTime.Year() {
		return ErrCardExpired
	} else if card.ExpirationYear() == currTime.Year() && card.ExpirationMonth() > currTime.Month() {
		return ErrCardExpired
	}

	if card.ExpirationMonth() > time.December || card.ExpirationMonth() < time.January {
		return ErrInvalidExpirationMonth
	}

	if len(card.Number()) != 16 {
		return ErrInvalidCardNumber
	}

	for _, d := range card.Number() {
		if !unicode.IsDigit(d) {
			return ErrInvalidCardNumber
		}
	}

	return nil
}
