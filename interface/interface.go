package _interface

import "github.com/aditya37/test-cc-validator/model"

type CCWriterReader interface {
	InsertCardNumber(value, network string) error
	GetCardNumber(value string) (model.MSTCard, error)
}
