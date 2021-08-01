package endpoint

import (
	"context"
	"github.com/amiranmanesh/imenaria-interview-task/bankcard/proto"
)

type (
	CreateRequest struct {
		BankName   string `json:"bank_name"`
		CardNumber string `json:"card_number"`
		UserID     uint   `json:"user_id"`
	}
	CreateResponse struct {
		Success bool `json:"success"`
		CardID  uint `json:"card_id"`
	}
	UpdateRequest struct {
		CardID     uint   `json:"card_id"`
		BankName   string `json:"bank_name"`
		CardNumber string `json:"card_number"`
	}
	UpdateResponse struct {
		Success bool `json:"success"`
	}
	DeleteRequest struct {
		CardID uint `json:"card_id"`
	}
	DeleteResponse struct {
		Success bool `json:"success"`
	}
	GetRequest struct {
		CardID uint `json:"card_id"`
	}
	GetResponse struct {
		Success    bool   `json:"success"`
		CardID     uint   `json:"card_id"`
		BankName   string `json:"bank_name"`
		CardNumber string `json:"card_number"`
		UserID     uint   `json:"user_id"`
	}
)

func EncodeCreateResponse(ctx context.Context, response interface{}) (interface{}, error) {
	res := response.(CreateResponse)
	return &proto.CreateResponse{
		Success: res.Success,
		CardId:  int32(res.CardID),
	}, nil
}

func DecodeCreateRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*proto.CreateRequest)
	return CreateRequest{
		BankName:   req.BankName,
		CardNumber: req.CardNumber,
		UserID:     uint(req.UserId),
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
		CardID:     uint(req.CardId),
		BankName:   req.BankName,
		CardNumber: req.CardNumber,
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
		CardID: uint(req.CardId),
	}, nil
}

func EncodeGetResponse(ctx context.Context, response interface{}) (interface{}, error) {
	res := response.(GetResponse)
	return &proto.GetResponse{
		Success:    res.Success,
		CardId:     int32(res.CardID),
		BankName:   res.BankName,
		CardNumber: res.CardNumber,
		UserId:     int32(res.UserID),
	}, nil
}

func DecodeGetRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*proto.GetRequest)
	return GetRequest{
		CardID: uint(req.CardId),
	}, nil
}
