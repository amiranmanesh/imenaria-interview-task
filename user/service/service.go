package service

import (
	"context"
	"fmt"
	"github.com/amiranmanesh/imenaria-interview-task/user/endpoint"
	"github.com/go-kit/kit/log"
)

type IRepository interface {
	Create(ctx context.Context, name, gender string, birthYear int, avatar string) (uint, error)
	Update(ctx context.Context, userId uint, name, gender string, birthYear int, avatar string) error
	Delete(ctx context.Context, userId uint) error
	Verify(ctx context.Context, userId uint) error
}

func NewService(repository IRepository, logger log.Logger) endpoint.IService {
	return &service{repository, log.With(logger, "service")}
}

type service struct {
	repository IRepository
	logger     log.Logger
}

func (s service) Create(ctx context.Context, name, gender string, birthYear int, avatar string) (uint, error) {
	if name == "" {
		return 0, fmt.Errorf("name can not be empty")
	}
	if gender == "" {
		return 0, fmt.Errorf("gender can not be empty")
	}
	if birthYear >= 1000 && birthYear <= 9999 {
		return 0, fmt.Errorf("birth year is out of range")
	}
	return s.repository.Create(ctx, name, gender, birthYear, avatar)
}

func (s service) Update(ctx context.Context, userId uint, name, gender string, birthYear int, avatar string) error {
	return s.repository.Update(ctx, userId, name, gender, birthYear, avatar)
}

func (s service) Delete(ctx context.Context, userId uint) error {
	return s.repository.Delete(ctx, userId)
}

func (s service) Verify(ctx context.Context, userId uint) error {
	return s.repository.Verify(ctx, userId)
}
