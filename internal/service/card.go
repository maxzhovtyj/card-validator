package service

import (
	"github.com/maxzhovtyj/card-validator/pkg/card"
)

type Card interface {
	Validate(card card.Card) error
}

type service struct {
	cardValidator card.Validator
}

func New(validator card.Validator) Card {
	return &service{
		cardValidator: validator,
	}
}

func (s service) Validate(card card.Card) error {
	return s.cardValidator.Valid(card)
}
