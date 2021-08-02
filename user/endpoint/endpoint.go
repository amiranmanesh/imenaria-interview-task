package endpoint

import (
	"context"
	"github.com/amiranmanesh/imenaria-interview-task/user/proto"
	"github.com/go-kit/kit/endpoint"
)

type IService interface {
	Create(ctx context.Context, userInfo *proto.UserInfo) (uint, error)
	Update(ctx context.Context, userId uint, userInfo *proto.UserInfo) error
	Delete(ctx context.Context, userId uint) error
	Get(ctx context.Context, userId uint) (*proto.UserInfo, error)
	Verify(ctx context.Context, userId uint) error
}

func MakeEndpoint(s IService) Endpoints {
	return Endpoints{
		Create: makeCreateEndpoint(s),
		Update: makeUpdateEndpoint(s),
		Delete: makeDeleteEndpoint(s),
		Get:    makeGetEndpoint(s),
		Verify: makeVerifyEndpoint(s),
	}
}

type Endpoints struct {
	Create endpoint.Endpoint
	Update endpoint.Endpoint
	Delete endpoint.Endpoint
	Get    endpoint.Endpoint
	Verify endpoint.Endpoint
}

func makeCreateEndpoint(s IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*proto.CreateRequest)
		uid, err := s.Create(ctx, &proto.UserInfo{
			Name:      req.Name,
			Gender:    req.Gender,
			BirthYear: req.BirthYear,
			Avatar:    req.Avatar,
		})
		if err != nil {
			return &proto.CreateResponse{Success: false}, err
		} else {
			return &proto.CreateResponse{
				Success: true,
				UserId:  int32(uid),
			}, nil
		}
	}
}

func makeUpdateEndpoint(s IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*proto.UpdateRequest)
		err := s.Update(ctx, uint(req.UserId), &proto.UserInfo{
			Name:      req.Name,
			Gender:    req.Gender,
			BirthYear: req.BirthYear,
			Avatar:    req.Avatar,
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
		err := s.Delete(ctx, uint(req.UserId))
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
		res, err := s.Get(ctx, uint(req.UserId))
		if err != nil {
			return &proto.GetResponse{Success: false}, err
		} else {
			return &proto.GetResponse{
				Success: true,
				Info:    res,
			}, nil
		}
	}
}

func makeVerifyEndpoint(s IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*proto.VerifyRequest)
		err := s.Verify(ctx, uint(req.UserId))
		if err != nil {
			return &proto.VerifyResponse{Success: false}, err
		} else {
			return &proto.VerifyResponse{
				Success: true,
			}, nil
		}
	}
}
