package transport

import (
	"context"
	"github.com/amiranmanesh/imenaria-interview-task/user/endpoint"
	"github.com/amiranmanesh/imenaria-interview-task/user/proto"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	createHandler grpctransport.Handler
	updateHandler grpctransport.Handler
	deleteHandler grpctransport.Handler
	getHandler    grpctransport.Handler
	verifyHandler grpctransport.Handler
}

func NewGRPCServer(ctx context.Context, endpoints endpoint.Endpoints) proto.UserServiceServer {
	return &grpcServer{
		createHandler: grpctransport.NewServer(
			endpoints.Create,
			endpoint.DecodeCreateRequest,
			endpoint.EncodeCreateResponse,
		),
		updateHandler: grpctransport.NewServer(
			endpoints.Update,
			endpoint.DecodeUpdateRequest,
			endpoint.EncodeUpdateResponse,
		),
		deleteHandler: grpctransport.NewServer(
			endpoints.Delete,
			endpoint.DecodeDeleteRequest,
			endpoint.EncodeDeleteResponse,
		),
		getHandler: grpctransport.NewServer(
			endpoints.Get,
			endpoint.DecodeGetRequest,
			endpoint.EncodeGetResponse,
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

func (s *grpcServer) Update(ctx context.Context, request *proto.UpdateRequest) (*proto.UpdateResponse, error) {
	_, response, err := s.updateHandler.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.(*proto.UpdateResponse), nil
}

func (s *grpcServer) Delete(ctx context.Context, request *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	_, response, err := s.deleteHandler.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.(*proto.DeleteResponse), nil
}

func (s *grpcServer) Get(ctx context.Context, request *proto.GetRequest) (*proto.GetResponse, error) {
	_, response, err := s.getHandler.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.(*proto.GetResponse), nil
}

func (s *grpcServer) Verify(ctx context.Context, request *proto.VerifyRequest) (*proto.VerifyResponse, error) {
	_, response, err := s.verifyHandler.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.(*proto.VerifyResponse), nil
}
