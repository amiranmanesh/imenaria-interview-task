package service

import (
	"context"
	"fmt"
	"github.com/amiranmanesh/imenaria-interview-task/bankcard/endpoint"
	"github.com/go-kit/kit/log"
)

type IRepository interface {
	Create(ctx context.Context, bankName, cardNumber string, userID uint) (uint, error)
	Update(ctx context.Context, cardId uint, bankName, cardNumber string) error
	Delete(ctx context.Context, cardId uint) error
	Get(ctx context.Context, cardId uint) (*BankCardModel, error)
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

func (s service) Create(ctx context.Context, bankName, cardNumber string, userID uint) (uint, error) {
	return s.repository.Create(ctx, bankName, cardNumber, userID)
}

func (s service) Update(ctx context.Context, cardId uint, bankName, cardNumber string) error {
	return s.repository.Update(ctx, cardId, bankName, cardNumber)
}

func (s service) Delete(ctx context.Context, cardId uint) error {
	return s.repository.Delete(ctx, cardId)
}

func (s service) Get(ctx context.Context, cardId uint) (uint, string, string, uint, error) {
	model, err := s.repository.Get(ctx, cardId)
	fmt.Println(1, model)
	if err != nil {
		return 0, "", "", 0, err
	} else {
		return model.CardID, model.BankName, model.CardNumber, model.UserID, err
	}
}
