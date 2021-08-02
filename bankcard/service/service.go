package service

import (
	"context"
	"fmt"
	"github.com/amiranmanesh/imenaria-interview-task/bankcard/endpoint"
	"github.com/amiranmanesh/imenaria-interview-task/bankcard/proto"
	"github.com/go-kit/kit/log"
)

type IRepository interface {
	Create(ctx context.Context, bankName, cardNumber string, userID uint) (uint, error)
	Update(ctx context.Context, cardInfo *proto.CardInfo) error
	Delete(ctx context.Context, cardId uint) error
	Get(ctx context.Context, cardId uint) (*proto.CardInfoFull, error)
	GetAll(ctx context.Context, userID uint) ([]*proto.CardInfo, error)
}

func NewService(repository IRepository, logger log.Logger) endpoint.IService {
	return &service{repository, log.With(logger, "service")}
}

type service struct {
	repository IRepository
	logger     log.Logger
}

func (s service) Create(ctx context.Context, bankName, cardNumber string, userID uint) (uint, error) {
	if bankName == "" {
		return 0, fmt.Errorf("bank name can not be empty")
	}
	if len(cardNumber) == 16 || len(cardNumber) == 20 {
		return s.repository.Create(ctx, bankName, cardNumber, userID)
	} else {
		return 0, fmt.Errorf("card number must be 16 or 20 characters")
	}
}

func (s service) Update(ctx context.Context, cardInfo *proto.CardInfo) error {
	if cardInfo.CardNumber == "" {
		return s.repository.Update(ctx, cardInfo)
	} else if len(cardInfo.CardNumber) == 16 || len(cardInfo.CardNumber) == 20 {
		return s.repository.Update(ctx, cardInfo)
	} else {
		return fmt.Errorf("card number must be 16 or 20 characters")
	}
}

func (s service) Delete(ctx context.Context, cardId uint) error {
	return s.repository.Delete(ctx, cardId)
}

func (s service) Get(ctx context.Context, cardId uint) (*proto.CardInfoFull, error) {
	return s.repository.Get(ctx, cardId)
}

func (s service) GetAll(ctx context.Context, cardId uint) ([]*proto.CardInfo, error) {
	return s.repository.GetAll(ctx, cardId)
}
