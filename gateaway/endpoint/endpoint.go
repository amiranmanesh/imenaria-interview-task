package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type IService interface {
	CreateUser(ctx context.Context, name, gender string, birthYear int, avatar string) (uint, error)
	UpdateUser(ctx context.Context, userId uint, name, gender string, birthYear int, avatar string) error
	DeleteUser(ctx context.Context, userId uint) error
	GetUser(ctx context.Context, userId uint) (UserModel, []BankCardModel, error)
	CreateCard(ctx context.Context, bankName, cardNumber string, userID uint) (uint, error)
	UpdateCard(ctx context.Context, cardId uint, bankName, cardNumber string) error
	DeleteCard(ctx context.Context, cardId uint) error
	GetCard(ctx context.Context, userId uint) (BankCardFullModel, error)
}

func MakeEndpoint(s IService) Endpoints {
	return Endpoints{
		CreateUser: makeCreateUserEndpoint(s),
		UpdateUser: makeUpdateUserEndpoint(s),
		DeleteUser: makeDeleteUserEndpoint(s),
		GetUser:    makeGetUserEndpoint(s),
		CreateCard: makeCreateCardEndpoint(s),
		UpdateCard: makeUpdateCardEndpoint(s),
		DeleteCard: makeDeleteCardEndpoint(s),
		GetCard:    makeGetCardEndpoint(s),
	}
}

type Endpoints struct {
	CreateUser endpoint.Endpoint
	UpdateUser endpoint.Endpoint
	DeleteUser endpoint.Endpoint
	GetUser    endpoint.Endpoint
	CreateCard endpoint.Endpoint
	UpdateCard endpoint.Endpoint
	DeleteCard endpoint.Endpoint
	GetCard    endpoint.Endpoint
}

func makeCreateUserEndpoint(s IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		uid, err := s.CreateUser(ctx, req.Name, req.Gender, req.BirthYear, req.Avatar)
		if err != nil {
			return CreateUserResponse{Success: false}, err
		} else {
			return CreateUserResponse{
				Success: true,
				UserID:  uid,
			}, nil
		}
	}
}

func makeUpdateUserEndpoint(s IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateUserRequest)
		err := s.UpdateUser(ctx, req.UserID, req.Name, req.Gender, req.BirthYear, req.Avatar)
		if err != nil {
			return UpdateUserResponse{Success: false}, err
		} else {
			return UpdateUserResponse{
				Success: true,
			}, nil
		}
	}
}

func makeDeleteUserEndpoint(s IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteUserRequest)
		err := s.DeleteUser(ctx, req.UserID)
		if err != nil {
			return DeleteUserResponse{Success: false}, err
		} else {
			return DeleteUserResponse{
				Success: true,
			}, nil
		}
	}
}

func makeGetUserEndpoint(s IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserRequest)
		userInfo, cards, err := s.GetUser(ctx, req.UserID)
		if err != nil {
			return GetUserResponse{Success: false}, err
		} else {
			return GetUserResponse{
				Success:   true,
				UserInfo:  userInfo,
				BankCards: cards,
			}, nil
		}
	}
}

func makeCreateCardEndpoint(s IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateCardRequest)
		cid, err := s.CreateCard(ctx, req.BankName, req.CardNumber, req.UserID)
		if err != nil {
			return CreateCardResponse{Success: false}, err
		} else {
			return CreateCardResponse{
				Success: true,
				CardID:  cid,
			}, nil
		}
	}
}

func makeUpdateCardEndpoint(s IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateCardRequest)
		err := s.UpdateCard(ctx, req.CardID, req.BankName, req.CardNumber)
		if err != nil {
			return UpdateCardResponse{Success: false}, err
		} else {
			return UpdateCardResponse{
				Success: true,
			}, nil
		}
	}
}

func makeDeleteCardEndpoint(s IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteCardRequest)
		err := s.DeleteCard(ctx, req.CardID)
		if err != nil {
			return DeleteCardResponse{Success: false}, err
		} else {
			return DeleteCardResponse{
				Success: true,
			}, nil
		}
	}
}

func makeGetCardEndpoint(s IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetCardRequest)
		res, err := s.GetCard(ctx, req.CardID)
		if err != nil {
			return GetCardResponse{Success: false}, err
		} else {
			return GetCardResponse{
				Success:  true,
				CardInfo: res,
			}, nil
		}
	}
}
