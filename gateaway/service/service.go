package service

import (
	"context"
	"fmt"
	cardProto "github.com/amiranmanesh/imenaria-interview-task/bankcard/proto"
	"github.com/amiranmanesh/imenaria-interview-task/gateaway/endpoint"
	userProto "github.com/amiranmanesh/imenaria-interview-task/user/proto"
	"github.com/go-kit/kit/log"
)

func NewService(userServiceClient userProto.UserServiceClient, cardServiceClient cardProto.CardServiceClient, logger log.Logger) endpoint.IService {
	return &service{userServiceClient, cardServiceClient, log.With(logger, "service")}
}

type service struct {
	userServiceClient userProto.UserServiceClient
	cardServiceClient cardProto.CardServiceClient
	logger            log.Logger
}

func (s service) CreateUser(ctx context.Context, name, gender string, birthYear int, avatar string) (uint, error) {
	//TODO: handle uploading avatar

	req := &userProto.CreateRequest{
		Name:      name,
		Gender:    gender,
		BirthYear: int32(birthYear),
		Avatar:    avatar,
	}
	res, err := s.userServiceClient.Create(ctx, req)
	if err != nil {
		//todo handle error msg
		return 0, err
	}

	if res.Success {
		return uint(res.UserId), nil
	} else {
		return 0, fmt.Errorf("error in save user")
	}
}

func (s service) UpdateUser(ctx context.Context, userId uint, name, gender string, birthYear int, avatar string) error {
	//TODO: handle uploading avatar if avatar is not empty

	req := &userProto.UpdateRequest{
		UserId:    int32(userId),
		Name:      name,
		Gender:    gender,
		BirthYear: int32(birthYear),
		Avatar:    avatar,
	}
	res, err := s.userServiceClient.Update(ctx, req)
	if err != nil {
		//todo handle error msg
		return err
	}

	if res.Success {
		return nil
	} else {
		return fmt.Errorf("error in update user")
	}
}

func (s service) DeleteUser(ctx context.Context, userId uint) error {
	req := &userProto.DeleteRequest{
		UserId: int32(userId),
	}
	res, err := s.userServiceClient.Delete(ctx, req)
	if err != nil {
		//todo handle error msg
		return err
	}

	if res.Success {
		return nil
	} else {
		return fmt.Errorf("error in delete user")
	}
}

func (s service) GetUser(ctx context.Context, userId uint) (endpoint.UserModel, []endpoint.BankCardModel, error) {
	//TODO: implement getting user
	panic("implement me")
}

func (s service) CreateCard(ctx context.Context, bankName, cardNumber string, userID uint) (uint, error) {
	req := &cardProto.CreateRequest{
		BankName:   bankName,
		CardNumber: cardNumber,
		UserId:     int32(userID),
	}
	res, err := s.cardServiceClient.Create(ctx, req)
	if err != nil {
		//todo handle error msg
		return 0, err
	}

	//TODO: assign card to user

	if res.Success {
		return uint(res.CardId), nil
	} else {
		return 0, fmt.Errorf("error in create card")
	}
}

func (s service) UpdateCard(ctx context.Context, cardId uint, bankName, cardNumber string) error {
	req := &cardProto.UpdateRequest{
		CardId:     int32(cardId),
		BankName:   bankName,
		CardNumber: cardNumber,
	}
	res, err := s.cardServiceClient.Update(ctx, req)
	if err != nil {
		//todo handle error msg
		return err
	}

	if res.Success {
		return nil
	} else {
		return fmt.Errorf("error in update card")
	}
}

func (s service) DeleteCard(ctx context.Context, cardId uint) error {
	req := &cardProto.DeleteRequest{
		CardId: int32(cardId),
	}
	res, err := s.cardServiceClient.Delete(ctx, req)
	if err != nil {
		//todo handle error msg
		return err
	}

	if res.Success {
		return nil
	} else {
		return fmt.Errorf("error in delete card")
	}
}

func (s service) GetCard(ctx context.Context, userId uint) (endpoint.BankCardFullModel, error) {
	//TODO: implement getting card
	panic("implement me")
}
