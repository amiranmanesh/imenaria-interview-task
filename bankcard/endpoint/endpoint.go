package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type IService interface {
	Create(ctx context.Context, bankName, cardNumber string, userID uint) (uint, error)
	Update(ctx context.Context, cardId uint, bankName, cardNumber string) error
	Delete(ctx context.Context, cardId uint) error
	Get(ctx context.Context, cardId uint) (BankCardFullModel, error)
	GetAll(ctx context.Context, userID uint) ([]BankCardModel, error)
}

func MakeEndpoint(s IService) Endpoints {
	return Endpoints{
		Create: makeCreateEndpoint(s),
		Update: makeUpdateEndpoint(s),
		Delete: makeDeleteEndpoint(s),
		Get:    makeGetEndpoint(s),
		GetAll: makeGetAllEndpoint(s),
	}
}

type Endpoints struct {
	Create endpoint.Endpoint
	Update endpoint.Endpoint
	Delete endpoint.Endpoint
	Get    endpoint.Endpoint
	GetAll endpoint.Endpoint
}

func makeCreateEndpoint(s IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		cid, err := s.Create(ctx, req.BankName, req.CardNumber, req.UserID)
		if err != nil {
			return CreateResponse{Success: false}, err
		} else {
			return CreateResponse{
				Success: true,
				CardID:  cid,
			}, nil
		}
	}
}

func makeUpdateEndpoint(s IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateRequest)
		err := s.Update(ctx, req.CardID, req.BankName, req.CardNumber)
		if err != nil {
			return UpdateResponse{Success: false}, err
		} else {
			return UpdateResponse{
				Success: true,
			}, nil
		}
	}
}

func makeDeleteEndpoint(s IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequest)
		err := s.Delete(ctx, req.CardID)
		if err != nil {
			return DeleteResponse{Success: false}, err
		} else {
			return DeleteResponse{
				Success: true,
			}, nil
		}
	}
}

func makeGetEndpoint(s IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		res, err := s.Get(ctx, req.CardID)
		if err != nil {
			return GetResponse{Success: false}, err
		} else {
			return GetResponse{
				Success:  true,
				CardInfo: res,
			}, nil
		}
	}
}

func makeGetAllEndpoint(s IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllRequest)
		res, err := s.GetAll(ctx, req.UserID)
		if err != nil {
			return GetAllResponse{Success: false}, err
		} else {
			return GetAllResponse{
				Success: true,
				UserID:  req.UserID,
				Cards:   res,
			}, nil
		}
	}
}
