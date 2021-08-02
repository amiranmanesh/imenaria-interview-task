package endpoint

import (
	"context"
	"github.com/amiranmanesh/imenaria-interview-task/bankcard/proto"
)

func EncodeCreateResponse(ctx context.Context, response interface{}) (interface{}, error) {
	return response.(*proto.CreateResponse), nil
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

func EncodeGetAllResponse(ctx context.Context, response interface{}) (interface{}, error) {
	return response.(*proto.GetAllResponse), nil
}

func DecodeGetAllRequest(ctx context.Context, request interface{}) (interface{}, error) {
	return request.(*proto.GetAllRequest), nil
}
