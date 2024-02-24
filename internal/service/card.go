package service

import (
	"github.com/maxzhovtyj/card-validator/internal/models"
	"github.com/maxzhovtyj/card-validator/pkg/card"
)

type Card interface {
	Validate(card models.Card) error
}

type service struct {
	cardValidator card.Validator
}

func New() Card {
	return &service{}
}

func (s service) Validate(card models.Card) error {
	return s.cardValidator.Valid(card)
}
