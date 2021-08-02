package endpoint

import (
	"context"
	"github.com/amiranmanesh/imenaria-interview-task/bankcard/proto"
)

type (
	BankCardModel struct {
		CardID     uint   `json:"card_id"`
		BankName   string `json:"bank_name"`
		CardNumber string `json:"card_number"`
	}
	BankCardFullModel struct {
		CardID     uint   `json:"card_id"`
		BankName   string `json:"bank_name"`
		CardNumber string `json:"card_number"`
		UserID     uint   `json:"user_id"`
	}

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
		Success  bool              `json:"success"`
		CardInfo BankCardFullModel `json:"card_info"`
	}
	GetAllRequest struct {
		UserID uint `json:"user_id"`
	}
	GetAllResponse struct {
		Success bool            `json:"success"`
		UserID  uint            `json:"user_id"`
		Cards   []BankCardModel `json:"cards"`
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
		CardId:     int32(res.CardInfo.CardID),
		BankName:   res.CardInfo.BankName,
		CardNumber: res.CardInfo.CardNumber,
		UserId:     int32(res.CardInfo.UserID),
	}, nil
}

func DecodeGetRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*proto.GetRequest)
	return GetRequest{
		CardID: uint(req.CardId),
	}, nil
}

func EncodeGetAllResponse(ctx context.Context, response interface{}) (interface{}, error) {
	res := response.(GetAllResponse)
	var items []*proto.CardInfo
	for _, card := range res.Cards {
		items = append(items, &proto.CardInfo{
			CardId:     int32(card.CardID),
			BankName:   card.BankName,
			CardNumber: card.CardNumber,
		})
	}
	return &proto.GetAllResponse{
		Success: res.Success,
		UserId:  int32(res.UserID),
		Cards:   items,
	}, nil
}

func DecodeGetAllRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*proto.GetAllRequest)
	return GetAllRequest{
		UserID: uint(req.UserId),
	}, nil
}
