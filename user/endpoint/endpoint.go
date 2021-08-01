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
		//req := request.(LoginRequest)
		//token, err := s.Login(ctx, req.Email, req.Password)
		//if err != nil {
		//	return LoginResponse{Success: false}, err
		//} else {
		//	return LoginResponse{
		//		Success: true,
		//		Token:   token,
		//	}, nil
		//}
	}
}

func makeVerifyEndpoint(s IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//req := request.(VerifyTokenRequest)
		//userID, err := s.Verify(ctx, req.Token)
		//if err != nil {
		//	return VerifyTokenResponse{Success: false}, err
		//} else {
		//	return VerifyTokenResponse{
		//		Success: true,
		//		UserID:  int32(userID),
		//	}, nil
		//}
	}
}
