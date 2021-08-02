package endpoint

import (
	"context"
	"encoding/json"
	"net/http"
)

type (
	UserModel struct {
		UserID    uint   `json:"user_id"`
		Name      string `json:"name"`
		Gender    string `json:"gender"`
		BirthYear int    `json:"birth_year"`
		Avatar    string `json:"avatar"`
	}
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

	CreateUserRequest struct {
		Name      string `json:"name"`
		Gender    string `json:"gender"`
		BirthYear int    `json:"birth_year"`
		Avatar    string `json:"avatar"`
	}
	CreateUserResponse struct {
		Success bool `json:"success"`
		UserID  uint `json:"user_id"`
	}
	UpdateUserRequest struct {
		UserID    uint   `json:"user_id"`
		Name      string `json:"name"`
		Gender    string `json:"gender"`
		BirthYear int    `json:"birth_year"`
		Avatar    string `json:"avatar"`
	}
	UpdateUserResponse struct {
		Success bool `json:"success"`
	}
	DeleteUserRequest struct {
		UserID uint `json:"user_id"`
	}
	DeleteUserResponse struct {
		Success bool `json:"success"`
	}
	GetUserRequest struct {
		UserID uint `json:"user_id"`
	}
	GetUserResponse struct {
		Success   bool            `json:"success"`
		UserInfo  UserModel       `json:"user_info"`
		BankCards []BankCardModel `json:"user_cards"`
	}

	CreateCardRequest struct {
		BankName   string `json:"bank_name"`
		CardNumber string `json:"card_number"`
		UserID     uint   `json:"user_id"`
	}
	CreateCardResponse struct {
		Success bool `json:"success"`
		CardID  uint `json:"card_id"`
	}
	UpdateCardRequest struct {
		CardID     uint   `json:"card_id"`
		BankName   string `json:"bank_name"`
		CardNumber string `json:"card_number"`
	}
	UpdateCardResponse struct {
		Success bool `json:"success"`
	}
	DeleteCardRequest struct {
		CardID uint `json:"card_id"`
	}
	DeleteCardResponse struct {
		Success bool `json:"success"`
	}
	GetCardRequest struct {
		CardID uint `json:"card_id"`
	}
	GetCardResponse struct {
		Success  bool              `json:"success"`
		CardInfo BankCardFullModel `json:"card_info"`
	}

	ErrorResponse struct {
		Success bool   `json:"success"`
		Error   string `json:"error"`
	}
)

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func DecodeCreateUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeUpdateUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req UpdateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeDeleteUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req DeleteUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeGetUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeCreateCardReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateCardRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeUpdateCardReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req UpdateCardRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeDeleteCardReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req DeleteCardRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeGetCardReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetCardRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
