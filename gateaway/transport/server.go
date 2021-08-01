package transport

import (
	"github.com/amiranmanesh/imenaria-interview-task/gateaway/endpoint"
	"github.com/gin-gonic/gin"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
)

//Transport Layer
func NewHTTPServer(endpoints endpoint.Endpoints) http.Handler {
	router := gin.Default()

	user := router.Group("/user")
	{
		user.POST("", func(c *gin.Context) {
			gin.WrapH(httptransport.NewServer(
				endpoints.CreateUser,
				endpoint.DecodeCreateUserReq,
				endpoint.EncodeResponse,
			))
		})
		user.GET("/:user_id", func(c *gin.Context) {
			gin.WrapH(httptransport.NewServer(
				endpoints.GetUser,
				endpoint.DecodeGetUserReq,
				endpoint.EncodeResponse,
			))
		})
		user.POST("/:user_id", func(c *gin.Context) {
			gin.WrapH(httptransport.NewServer(
				endpoints.UpdateUser,
				endpoint.DecodeUpdateUserReq,
				endpoint.EncodeResponse,
			))
		})
		user.DELETE("/:user_id", func(c *gin.Context) {
			gin.WrapH(httptransport.NewServer(
				endpoints.DeleteUser,
				endpoint.DecodeDeleteUserReq,
				endpoint.EncodeResponse,
			))
		})
	}

	card := router.Group("/card")
	{
		card.POST("", func(c *gin.Context) {
			gin.WrapH(httptransport.NewServer(
				endpoints.CreateCard,
				endpoint.DecodeCreateCardReq,
				endpoint.EncodeResponse,
			))
		})
		card.GET("/:card_id", func(c *gin.Context) {
			gin.WrapH(httptransport.NewServer(
				endpoints.GetCard,
				endpoint.DecodeGetCardReq,
				endpoint.EncodeResponse,
			))
		})
		card.POST("/:card_id", func(c *gin.Context) {
			gin.WrapH(httptransport.NewServer(
				endpoints.UpdateCard,
				endpoint.DecodeUpdateCardReq,
				endpoint.EncodeResponse,
			))
		})
		card.DELETE("/:card_id", func(c *gin.Context) {
			gin.WrapH(httptransport.NewServer(
				endpoints.DeleteCard,
				endpoint.DecodeDeleteCardReq,
				endpoint.EncodeResponse,
			))
		})
	}

	return router
}
