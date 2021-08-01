package service

import (
	"context"
)

type IRepository interface {
	Create(ctx context.Context, name, gender string, birthYear int, avatar string) (uint, error)
	Verify(ctx context.Context, id uint) error
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
//func (s service) SignUp(ctx context.Context, name, email, password string) (string, error) {
//	hashPassword := encrypting.GetHashedPassword(password)
//	return s.repository.SignUp(ctx, name, email, hashPassword)
//}
//
//func (s service) Login(ctx context.Context, email, password string) (string, error) {
//	hashPassword := encrypting.GetHashedPassword(password)
//	return s.repository.Login(ctx, email, hashPassword)
//}
//
//func (s service) Verify(ctx context.Context, token string) (uint, error) {
//	return s.repository.Verify(ctx, token)
//}
