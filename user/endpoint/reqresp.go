package endpoint

import (
	"context"
	"github.com/amiranmanesh/imenaria-interview-task/user/proto"
)

type (
	CreateRequest struct {
		Name      string `json:"name"`
		Gender    string `json:"gender"`
		BirthYear int    `json:"birth_year"`
		Avatar    string `json:"avatar"`
	}
	CreateResponse struct {
		Success bool `json:"success"`
		UserID  uint `json:"user_id"`
	}
	UpdateRequest struct {
		UserID    uint   `json:"user_id"`
		Name      string `json:"name"`
		Gender    string `json:"gender"`
		BirthYear int    `json:"birth_year"`
		Avatar    string `json:"avatar"`
	}
	UpdateResponse struct {
		Success bool `json:"success"`
	}
	DeleteRequest struct {
		UserID uint `json:"user_id"`
	}
	DeleteResponse struct {
		Success bool `json:"success"`
	}
	VerifyRequest struct {
		UserID uint `json:"user_id"`
	}
	VerifyResponse struct {
		Success bool `json:"success"`
	}
)

func EncodeCreateResponse(ctx context.Context, response interface{}) (interface{}, error) {
	res := response.(CreateResponse)
	return &proto.CreateResponse{
		Success: res.Success,
		UserId:  int32(res.UserID),
	}, nil
}

func DecodeCreateRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*proto.CreateRequest)
	return CreateRequest{
		Name:      req.Name,
		Gender:    req.Gender,
		BirthYear: int(req.BirthYear),
		Avatar:    req.Avatar,
	}, nil
}

func EncodeVerifyResponse(ctx context.Context, response interface{}) (interface{}, error) {
	res := response.(VerifyResponse)
	return &proto.VerifyResponse{
		Success: res.Success,
	}, nil
}

func EncodeUpdateResponse(ctx context.Context, response interface{}) (interface{}, error) {
	res := response.(UpdateResponse)
	return &proto.UpdateResponse{
		Success: res.Success,
	}, nil
}

func DecodeUpdateRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*proto.UpdateRequest)
	return UpdateRequest{
		UserID:    uint(req.UserId),
		Name:      req.Name,
		Gender:    req.Gender,
		BirthYear: int(req.BirthYear),
		Avatar:    req.Avatar,
	}, nil
}

func EncodeDeleteResponse(ctx context.Context, response interface{}) (interface{}, error) {
	res := response.(DeleteResponse)
	return &proto.DeleteResponse{
		Success: res.Success,
	}, nil
}

func DecodeDeleteRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*proto.DeleteRequest)
	return DeleteRequest{
		UserID: uint(req.UserId),
	}, nil
}
func DecodeVerifyRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*proto.VerifyRequest)
	return VerifyRequest{
		UserID: uint(req.UserId),
	}, nil
}
