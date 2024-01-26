package service

import (
	"context"
	"errors"

	_interface "github.com/aditya37/test-cc-validator/interface"
	"github.com/aditya37/test-cc-validator/model"
	"github.com/aditya37/test-cc-validator/utils"
)

type Service struct {
	repo _interface.CCWriterReader
}
type Iservice interface {
	GetCardNumber(ctx context.Context, value string) (model.ResponseValidateCard, error)
	InsertCardNumber(ctx context.Context, request model.RequestAddCard) (model.ResponseAddCard, error)
}

func NewService(repo _interface.CCWriterReader) Iservice {
	return &Service{
		repo: repo,
	}
}

var ErrUnknownCard = errors.New("unknown card")

func (s *Service) GetCardNumber(ctx context.Context, value string) (model.ResponseValidateCard, error) {
	prefix := utils.NewCardValidator().FirstNumber(value)
	if prefix == "4" {
		card, err := s.repo.GetCardNumber(value)
		if err != nil {
			return model.ResponseValidateCard{}, err
		}
		return model.ResponseValidateCard{
			CardNetwork: card.CardNetwork,
			Number:      value,
			IsValid:     true,
		}, nil
	} else if prefix == "5" && utils.NewCardValidator().InRange(value[0:4], "2221", "2720") {
		card, err := s.repo.GetCardNumber(value)
		if err != nil {
			return model.ResponseValidateCard{}, err
		}
		return model.ResponseValidateCard{
			Number:      value,
			IsValid:     true,
			CardNetwork: card.CardNetwork,
		}, nil
	}
	return model.ResponseValidateCard{}, ErrUnknownCard

}

func (s *Service) InsertCardNumber(ctx context.Context, request model.RequestAddCard) (model.ResponseAddCard, error) {
	err := s.repo.InsertCardNumber(request.Number, request.CardNetwork)
	if err != nil {
		return model.ResponseAddCard{}, err
	}
	return model.ResponseAddCard{
		CardNetwork: request.CardNetwork,
	}, nil
}
