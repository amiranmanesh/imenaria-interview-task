package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type IService interface {
	Create(ctx context.Context, name, gender string, birthYear int, avatar string) (uint, error)
	Verify(ctx context.Context, id uint) error
}

func MakeEndpoint(s IService) Endpoints {
	return Endpoints{
		Create: makeCreateEndpoint(s),
		Verify: makeVerifyEndpoint(s),
	}
}

type Endpoints struct {
	Create endpoint.Endpoint
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
