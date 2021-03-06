package endpoint

import (
	"context"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"strconv"
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

	UploadAvatarRequest struct {
		File                multipart.File        `json:"_"`
		MultipartFileHeader *multipart.FileHeader `json:"_"`
	}
	UploadAvatarResponse struct {
		Success  bool   `json:"success"`
		FileCode string `json:"file_code"`
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
	switch response.(type) {
	case ErrorResponse:
		w.WriteHeader(400)
	}
	w.Header().Add("Content-Type", "application/json")
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
	userId, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		return nil, err
	}
	var req GetUserRequest
	req.UserID = uint(userId)
	return req, nil
}

func DecodeUploadAvatarReq(ctx context.Context, r *http.Request) (interface{}, error) {
	f, fh, err := r.FormFile("avatar")
	if err != nil {
		return nil, err
	}
	req := UploadAvatarRequest{
		File:                f,
		MultipartFileHeader: fh,
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
	userId, err := strconv.Atoi(r.URL.Query().Get("card_id"))
	if err != nil {
		return nil, err
	}
	var req GetCardRequest
	req.CardID = uint(userId)
	return req, nil
}
