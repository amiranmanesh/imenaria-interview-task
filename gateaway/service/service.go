package service

import (
	"context"
	"fmt"
	cardProto "github.com/amiranmanesh/imenaria-interview-task/bankcard/proto"
	"github.com/amiranmanesh/imenaria-interview-task/gateaway/endpoint"
	userProto "github.com/amiranmanesh/imenaria-interview-task/user/proto"
	"github.com/amiranmanesh/imenaria-interview-task/utils/files"
	"github.com/go-kit/kit/log"
	"mime/multipart"
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
		return err
	}

	if res.Success {
		return nil
	} else {
		return fmt.Errorf("error in delete user")
	}
}

func (s service) GetUser(ctx context.Context, userId uint) (*endpoint.UserModel, []endpoint.BankCardModel, error) {
	reqUser := &userProto.GetRequest{
		UserId: int32(userId),
	}
	resUser, err := s.userServiceClient.Get(ctx, reqUser)
	if err != nil {
		return nil, nil, err
	}

	userInfo := &endpoint.UserModel{
		UserID:    userId,
		Name:      resUser.Info.Name,
		Gender:    resUser.Info.Gender,
		BirthYear: int(resUser.Info.BirthYear),
		Avatar:    resUser.Info.Avatar,
	}

	reqCards := &cardProto.GetAllRequest{
		UserId: int32(userId),
	}
	resCards, err := s.cardServiceClient.GetAll(ctx, reqCards)
	if err != nil {
		return nil, nil, err
	}

	var cardsResult []endpoint.BankCardModel
	for _, card := range resCards.Cards {
		temp := endpoint.BankCardModel{}
		temp.CardID = uint(card.CardId)
		temp.CardNumber = card.CardNumber
		temp.BankName = card.BankName

		cardsResult = append(cardsResult, temp)
	}

	return userInfo, cardsResult, nil
}

func (s service) UploadAvatar(ctx context.Context, file multipart.File, multipartFileHeader *multipart.FileHeader) (string, error) {
	return files.FilesHandler.Save(file, multipartFileHeader)
}

func (s service) CreateCard(ctx context.Context, bankName, cardNumber string, userID uint) (uint, error) {
	reqUser := &userProto.VerifyRequest{
		UserId: int32(userID),
	}
	resUser, err := s.userServiceClient.Verify(ctx, reqUser)
	if err != nil {
		return 0, err
	}
	if !resUser.Success {
		return 0, fmt.Errorf("user does not exist")
	}

	req := &cardProto.CreateRequest{
		BankName:   bankName,
		CardNumber: cardNumber,
		UserId:     int32(userID),
	}
	res, err := s.cardServiceClient.Create(ctx, req)
	if err != nil {
		return 0, err
	}

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
		return err
	}

	if res.Success {
		return nil
	} else {
		return fmt.Errorf("error in delete card")
	}
}

func (s service) GetCard(ctx context.Context, cardId uint) (*endpoint.BankCardFullModel, error) {
	req := &cardProto.GetRequest{
		CardId: int32(cardId),
	}
	res, err := s.cardServiceClient.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	resModel := &endpoint.BankCardFullModel{
		CardID:     uint(res.CardId),
		BankName:   res.BankName,
		CardNumber: res.CardNumber,
		UserID:     uint(res.UserId),
	}

	if res.Success {
		return resModel, nil
	} else {
		return nil, fmt.Errorf("error in getting card")
	}
}
