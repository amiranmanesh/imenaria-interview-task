package endpoint

import (
	"context"
	"github.com/amiranmanesh/imenaria-interview-task/bankcard/proto"
	"github.com/go-kit/kit/endpoint"
)

type IService interface {
	Create(ctx context.Context, bankName, cardNumber string, userID uint) (uint, error)
	Update(ctx context.Context, cardInfo *proto.CardInfo) error
	Delete(ctx context.Context, cardId uint) error
	Get(ctx context.Context, cardId uint) (*proto.CardInfoFull, error)
	GetAll(ctx context.Context, userID uint) ([]*proto.CardInfo, error)
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
		req := request.(*proto.CreateRequest)
		cid, err := s.Create(ctx, req.BankName, req.CardNumber, uint(req.UserId))
		if err != nil {
			return &proto.CreateResponse{Success: false}, err
		} else {
			return &proto.CreateResponse{
				Success: true,
				CardId:  int32(cid),
			}, nil
		}
	}
}

func makeUpdateEndpoint(s IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*proto.UpdateRequest)
		err := s.Update(ctx, &proto.CardInfo{
			CardId:     req.CardId,
			BankName:   req.BankName,
			CardNumber: req.CardNumber,
		})
		if err != nil {
			return &proto.UpdateResponse{Success: false}, err
		} else {
			return &proto.UpdateResponse{
				Success: true,
			}, nil
		}
	}
}

func makeDeleteEndpoint(s IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*proto.DeleteRequest)
		err := s.Delete(ctx, uint(req.CardId))
		if err != nil {
			return &proto.DeleteResponse{Success: false}, err
		} else {
			return &proto.DeleteResponse{
				Success: true,
			}, nil
		}
	}
}

func makeGetEndpoint(s IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*proto.GetRequest)
		res, err := s.Get(ctx, uint(req.CardId))
		if err != nil {
			return &proto.GetResponse{Success: false}, err
		} else {
			return &proto.GetResponse{
				Success:  true,
				CardInfo: res,
			}, nil
		}
	}
}

func makeGetAllEndpoint(s IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*proto.GetAllRequest)
		res, err := s.GetAll(ctx, uint(req.UserId))
		if err != nil {
			return &proto.GetAllResponse{Success: false}, err
		} else {
			return &proto.GetAllResponse{
				Success: true,
				UserId:  req.UserId,
				Cards:   res,
			}, nil
		}
	}
}
