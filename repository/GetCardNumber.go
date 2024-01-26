package repository

import (
	"errors"

	"github.com/aditya37/test-cc-validator/model"
)

var ErrNotFound = errors.New("number not found")

func (cs *cardStore) GetCardNumber(value string) (model.MSTCard, error) {
	card, ok := cs.db[value]
	if !ok {
		return model.MSTCard{}, ErrNotFound
	}
	return card, nil
}
