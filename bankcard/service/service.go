package service

import (
	"context"
)

type IRepository interface {
	Create(ctx context.Context, bankName, bankCardNumber string, userID uint) (uint, error)
	GetCardInfoByCardID(ctx context.Context, cardId uint) (*BankCardModel, error)
	GetCardsByUserID(ctx context.Context, userId uint) ([]*BankCardModel, error)
}

type BankCardModel struct {
	BankName   string
	CardNumber string
	UserID     uint
	CardID     uint
}

//
//func NewService(repository IRepository, logger log.Logger) endpoint.IService {
//	return &service{repository, log.With(logger, "service")}
//}
//
//type service struct {
//	repository IRepository
//	logger     log.Logger
//}
//
//func (s service) Create(ctx context.Context, name, gender string, birthYear int, avatar string) (uint, error) {
//	return s.repository.Create(ctx, name, gender, birthYear, avatar)
//}
//
//func (s service) Verify(ctx context.Context, id uint) error {
//	return s.repository.Verify(ctx, id)
//}
