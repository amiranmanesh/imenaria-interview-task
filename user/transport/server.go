package transport

import (
	"context"
	"github.com/amiranmanesh/imenaria-interview-task/user/endpoint"
	"github.com/amiranmanesh/imenaria-interview-task/user/proto"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	createHandler grpctransport.Handler
	verifyHandler grpctransport.Handler
}

func NewGRPCServer(ctx context.Context, endpoints endpoint.Endpoints) proto.UserServiceServer {
	return &grpcServer{
		createHandler: grpctransport.NewServer(
			endpoints.Create,
			endpoint.DecodeCreateRequest,
			endpoint.EncodeCreateResponse,
		),
		verifyHandler: grpctransport.NewServer(
			endpoints.Verify,
			endpoint.DecodeVerifyRequest,
			endpoint.EncodeVerifyResponse,
		),
	}
}

func (s *grpcServer) Create(ctx context.Context, request *proto.CreateRequest) (*proto.CreateResponse, error) {
	_, response, err := s.createHandler.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.(*proto.CreateResponse), nil
}

func (s *grpcServer) Verify(ctx context.Context, request *proto.VerifyRequest) (*proto.VerifyResponse, error) {
	_, response, err := s.verifyHandler.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.(*proto.VerifyResponse), nil
}
