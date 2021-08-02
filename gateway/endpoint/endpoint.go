package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"mime/multipart"
)

type IService interface {
	CreateUser(ctx context.Context, name, gender string, birthYear int, avatar string) (uint, error)
	UpdateUser(ctx context.Context, userId uint, name, gender string, birthYear int, avatar string) error
	DeleteUser(ctx context.Context, userId uint) error
	GetUser(ctx context.Context, userId uint) (*UserModel, []BankCardModel, error)
	UploadAvatar(ctx context.Context, file multipart.File, multipartFileHeader *multipart.FileHeader) (string, error)
	CreateCard(ctx context.Context, bankName, cardNumber string, userID uint) (uint, error)
	UpdateCard(ctx context.Context, cardId uint, bankName, cardNumber string) error
	DeleteCard(ctx context.Context, cardId uint) error
	GetCard(ctx context.Context, cardId uint) (*BankCardFullModel, error)
}

func MakeEndpoint(s IService) Endpoints {
	return Endpoints{
		CreateUser:   makeCreateUserEndpoint(s),
		UpdateUser:   makeUpdateUserEndpoint(s),
		DeleteUser:   makeDeleteUserEndpoint(s),
		GetUser:      makeGetUserEndpoint(s),
		UploadAvatar: makeUploadAvatarEndpoint(s),
		CreateCard:   makeCreateCardEndpoint(s),
		UpdateCard:   makeUpdateCardEndpoint(s),
		DeleteCard:   makeDeleteCardEndpoint(s),
		GetCard:      makeGetCardEndpoint(s),
	}
}

type Endpoints struct {
	CreateUser   endpoint.Endpoint
	UpdateUser   endpoint.Endpoint
	DeleteUser   endpoint.Endpoint
	GetUser      endpoint.Endpoint
	UploadAvatar endpoint.Endpoint
	CreateCard   endpoint.Endpoint
	UpdateCard   endpoint.Endpoint
	DeleteCard   endpoint.Endpoint
	GetCard      endpoint.Endpoint
}

func makeCreateUserEndpoint(s IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		uid, err := s.CreateUser(ctx, req.Name, req.Gender, req.BirthYear, req.Avatar)
		if err != nil {
			return ErrorResponse{Success: false, Error: err.Error()}, nil
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
			return ErrorResponse{Success: false, Error: err.Error()}, nil
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
			return ErrorResponse{Success: false, Error: err.Error()}, nil
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
			return ErrorResponse{Success: false, Error: err.Error()}, nil
		} else {
			return GetUserResponse{
				Success:   true,
				UserInfo:  *userInfo,
				BankCards: cards,
			}, nil
		}
	}
}

func makeUploadAvatarEndpoint(s IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UploadAvatarRequest)
		fileCode, err := s.UploadAvatar(ctx, req.File, req.MultipartFileHeader)
		if err != nil {
			return ErrorResponse{Success: false, Error: err.Error()}, nil
		} else {
			return UploadAvatarResponse{
				Success:  true,
				FileCode: fileCode,
			}, nil
		}
	}
}

func makeCreateCardEndpoint(s IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateCardRequest)
		cid, err := s.CreateCard(ctx, req.BankName, req.CardNumber, req.UserID)
		if err != nil {
			return ErrorResponse{Success: false, Error: err.Error()}, nil
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
			return ErrorResponse{Success: false, Error: err.Error()}, nil
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
			return ErrorResponse{Success: false, Error: err.Error()}, nil
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
			return ErrorResponse{Success: false, Error: err.Error()}, nil
		} else {
			return GetCardResponse{
				Success:  true,
				CardInfo: *res,
			}, nil
		}
	}
}
