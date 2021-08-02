package transport

import (
	"github.com/amiranmanesh/imenaria-interview-task/gateway/endpoint"
	"github.com/gin-gonic/gin"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
)

//Transport Layer
func NewHTTPServer(endpoints endpoint.Endpoints) http.Handler {
	router := gin.Default()

	user := router.Group("/user")
	{
		user.POST("/create", gin.WrapH(httptransport.NewServer(
			endpoints.CreateUser,
			endpoint.DecodeCreateUserReq,
			endpoint.EncodeResponse,
		)))
		user.GET("/get", gin.WrapH(httptransport.NewServer(
			endpoints.GetUser,
			endpoint.DecodeGetUserReq,
			endpoint.EncodeResponse,
		)))
		user.POST("/update", gin.WrapH(httptransport.NewServer(
			endpoints.UpdateUser,
			endpoint.DecodeUpdateUserReq,
			endpoint.EncodeResponse,
		)))
		user.DELETE("/delete", gin.WrapH(httptransport.NewServer(
			endpoints.DeleteUser,
			endpoint.DecodeDeleteUserReq,
			endpoint.EncodeResponse,
		)))
		user.POST("/upload", gin.WrapH(httptransport.NewServer(
			endpoints.UploadAvatar,
			endpoint.DecodeUploadAvatarReq,
			endpoint.EncodeResponse,
		)))
	}

	card := router.Group("/card")
	{
		card.POST("/create", gin.WrapH(httptransport.NewServer(
			endpoints.CreateCard,
			endpoint.DecodeCreateCardReq,
			endpoint.EncodeResponse,
		)))
		card.GET("/get", gin.WrapH(httptransport.NewServer(
			endpoints.GetCard,
			endpoint.DecodeGetCardReq,
			endpoint.EncodeResponse,
		)))
		card.POST("/update", gin.WrapH(httptransport.NewServer(
			endpoints.UpdateCard,
			endpoint.DecodeUpdateCardReq,
			endpoint.EncodeResponse,
		)))
		card.DELETE("/delete", gin.WrapH(httptransport.NewServer(
			endpoints.DeleteCard,
			endpoint.DecodeDeleteCardReq,
			endpoint.EncodeResponse,
		)))
	}

	return router
}
