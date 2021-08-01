package service

import (
	"context"
	"github.com/amiranmanesh/imenaria-interview-task/bankcard/endpoint"
	"github.com/go-kit/kit/log"
)

type IRepository interface {
	Create(ctx context.Context, bankName, bankCardNumber string, userID uint) (uint, error)
	Update(ctx context.Context, cardId uint, bankName, cardNumber string) error
	Delete(ctx context.Context, cardId uint) error
	GetCardByCardID(ctx context.Context, cardId uint) (*BankCardModel, error)
}

type BankCardModel struct {
	BankName   string
	CardNumber string
	UserID     uint
	CardID     uint
}

func NewService(repository IRepository, logger log.Logger) endpoint.IService {
	return &service{repository, log.With(logger, "service")}
}

type service struct {
	repository IRepository
	logger     log.Logger
}

func (s service) Create(ctx context.Context, bankName, bankCardNumber string, userID uint) (uint, error) {
	return s.repository.Create(ctx, bankName, bankCardNumber, userID)
}

func (s service) Update(ctx context.Context, cardId uint, bankName, cardNumber string) error {
	return s.repository.Update(ctx, cardId, bankName, cardNumber)
}

func (s service) Delete(ctx context.Context, cardId uint) error {
	return s.repository.Delete(ctx, cardId)
}

func (s service) GetCardByCardID(ctx context.Context, cardId uint) (*uint, *string, *string, *uint, error) {
	model, err := s.repository.GetCardByCardID(ctx, cardId)
	if err != nil {
		return &model.CardID, &model.BankName, &model.CardNumber, &model.UserID, err
	} else {
		return nil, nil, nil, nil, err
	}
}
