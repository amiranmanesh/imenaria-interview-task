package endpoint

import (
	"context"
	"github.com/amiranmanesh/imenaria-interview-task/user/proto"
)

func EncodeCreateResponse(ctx context.Context, response interface{}) (interface{}, error) {
	res := response.(proto.CreateResponse)
	return &proto.CreateResponse{
		Success: res.Success,
		UserId:  res.UserId,
	}, nil
}

func DecodeCreateRequest(ctx context.Context, request interface{}) (interface{}, error) {
	return request.(*proto.CreateRequest), nil
}

func EncodeUpdateResponse(ctx context.Context, response interface{}) (interface{}, error) {
	return response.(*proto.UpdateResponse), nil
}

func DecodeUpdateRequest(ctx context.Context, request interface{}) (interface{}, error) {
	return request.(*proto.UpdateRequest), nil
}

func EncodeDeleteResponse(ctx context.Context, response interface{}) (interface{}, error) {
	return response.(*proto.DeleteResponse), nil
}

func DecodeDeleteRequest(ctx context.Context, request interface{}) (interface{}, error) {
	return request.(*proto.DeleteRequest), nil
}

func EncodeGetResponse(ctx context.Context, response interface{}) (interface{}, error) {
	return response.(*proto.GetResponse), nil
}

func DecodeGetRequest(ctx context.Context, request interface{}) (interface{}, error) {
	return request.(*proto.GetRequest), nil
}

func EncodeVerifyResponse(ctx context.Context, response interface{}) (interface{}, error) {
	return response.(*proto.VerifyResponse), nil
}

func DecodeVerifyRequest(ctx context.Context, request interface{}) (interface{}, error) {
	return request.(*proto.VerifyRequest), nil
}
