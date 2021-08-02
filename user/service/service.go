package service

import (
	"context"
	"fmt"
	"github.com/amiranmanesh/imenaria-interview-task/user/endpoint"
	"github.com/amiranmanesh/imenaria-interview-task/user/proto"
	"github.com/go-kit/kit/log"
)

type IRepository interface {
	Create(ctx context.Context, userInfo *proto.UserInfo) (uint, error)
	Update(ctx context.Context, userId uint, userInfo *proto.UserInfo) error
	Delete(ctx context.Context, userId uint) error
	Get(ctx context.Context, userId uint) (*proto.UserInfo, error)
	Verify(ctx context.Context, userId uint) error
}

func NewService(repository IRepository, logger log.Logger) endpoint.IService {
	return &service{repository, log.With(logger, "service")}
}

type service struct {
	repository IRepository
	logger     log.Logger
}

func (s service) Create(ctx context.Context, userInfo *proto.UserInfo) (uint, error) {
	if userInfo.Name == "" {
		return 0, fmt.Errorf("name can not be empty")
	}
	if userInfo.Gender == "" {
		return 0, fmt.Errorf("gender can not be empty")
	}
	if userInfo.BirthYear <= 1000 && userInfo.BirthYear >= 9999 {
		return 0, fmt.Errorf("birth year is out of range")
	}
	return s.repository.Create(ctx, userInfo)
}

func (s service) Update(ctx context.Context, userId uint, userInfo *proto.UserInfo) error {
	return s.repository.Update(ctx, userId, userInfo)
}

func (s service) Delete(ctx context.Context, userId uint) error {
	return s.repository.Delete(ctx, userId)
}

func (s service) Get(ctx context.Context, userId uint) (*proto.UserInfo, error) {
	return s.repository.Get(ctx, userId)
}

func (s service) Verify(ctx context.Context, userId uint) error {
	return s.repository.Verify(ctx, userId)
}
