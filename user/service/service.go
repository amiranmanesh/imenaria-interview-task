package service

import (
	"context"
	"github.com/amiranmanesh/imenaria-interview-task/user/endpoint"
	"github.com/go-kit/kit/log"
)

type IRepository interface {
	Create(ctx context.Context, name, gender string, birthYear int, avatar string) (uint, error)
	Verify(ctx context.Context, id uint) error
}

func NewService(repository IRepository, logger log.Logger) endpoint.IService {
	return &service{repository, log.With(logger, "service")}
}

type service struct {
	repository IRepository
	logger     log.Logger
}

func (s service) Create(ctx context.Context, name, gender string, birthYear int, avatar string) (uint, error) {
	return s.repository.Create(ctx, name, gender, birthYear, avatar)
}

func (s service) Verify(ctx context.Context, id uint) error {
	return s.repository.Verify(ctx, id)
}
