package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type IService interface {
	Create(ctx context.Context, name, gender string, birthYear int, avatar string) (uint, error)
	Update(ctx context.Context, userId uint, name, gender string, birthYear int, avatar string) error
	Delete(ctx context.Context, userId uint) error
	Verify(ctx context.Context, userId uint) error
}

func MakeEndpoint(s IService) Endpoints {
	return Endpoints{
		Create: makeCreateEndpoint(s),
		Update: makeCreateEndpoint(s),
		Delete: makeCreateEndpoint(s),
		Verify: makeVerifyEndpoint(s),
	}
}

type Endpoints struct {
	Create endpoint.Endpoint
	Update endpoint.Endpoint
	Delete endpoint.Endpoint
	Verify endpoint.Endpoint
}

func makeCreateEndpoint(s IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		uid, err := s.Create(ctx, req.Name, req.Gender, req.BirthYear, req.Avatar)
		if err != nil {
			return CreateResponse{Success: false}, err
		} else {
			return CreateResponse{
				Success: true,
				UserID:  uid,
			}, nil
		}
	}
}

func makeUpdateEndpoint(s IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateRequest)
		err := s.Update(ctx, req.UserID, req.Name, req.Gender, req.BirthYear, req.Avatar)
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
		err := s.Delete(ctx, req.UserID)
		if err != nil {
			return DeleteResponse{Success: false}, err
		} else {
			return DeleteResponse{
				Success: true,
			}, nil
		}
	}
}

func makeVerifyEndpoint(s IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(VerifyRequest)
		err := s.Verify(ctx, req.UserID)
		if err != nil {
			return VerifyResponse{Success: false}, err
		} else {
			return VerifyResponse{
				Success: true,
			}, nil
		}
	}
}
