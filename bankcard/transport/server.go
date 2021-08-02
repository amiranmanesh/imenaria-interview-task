package transport

import (
	"context"
	"github.com/amiranmanesh/imenaria-interview-task/bankcard/endpoint"
	"github.com/amiranmanesh/imenaria-interview-task/bankcard/proto"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	createHandler grpctransport.Handler
	updateHandler grpctransport.Handler
	deleteHandler grpctransport.Handler
	getHandler    grpctransport.Handler
	getAllHandler grpctransport.Handler
}

func NewGRPCServer(ctx context.Context, endpoints endpoint.Endpoints) proto.CardServiceServer {
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
		getAllHandler: grpctransport.NewServer(
			endpoints.GetAll,
			endpoint.DecodeGetAllRequest,
			endpoint.EncodeGetAllResponse,
		),
	}
}

func (g *grpcServer) Create(ctx context.Context, request *proto.CreateRequest) (*proto.CreateResponse, error) {
	_, response, err := g.createHandler.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.(*proto.CreateResponse), nil
}

func (g *grpcServer) Update(ctx context.Context, request *proto.UpdateRequest) (*proto.UpdateResponse, error) {
	_, response, err := g.updateHandler.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.(*proto.UpdateResponse), nil
}

func (g *grpcServer) Delete(ctx context.Context, request *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	_, response, err := g.deleteHandler.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.(*proto.DeleteResponse), nil
}

func (g *grpcServer) Get(ctx context.Context, request *proto.GetRequest) (*proto.GetResponse, error) {
	_, response, err := g.getHandler.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.(*proto.GetResponse), nil
}

func (g *grpcServer) GetAll(ctx context.Context, request *proto.GetAllRequest) (*proto.GetAllResponse, error) {
	_, response, err := g.getAllHandler.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.(*proto.GetAllResponse), nil
}
