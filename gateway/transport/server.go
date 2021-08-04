// Package Classification of the main API
//
// Documentation for the User and the Bank Card API
//
//     Schemes: http
//     Host: localhost
//     BasePath: /
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Amir Iranmanesh <iranmanesh.ah@gmail.com>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package transport

import (
	"github.com/amiranmanesh/imenaria-interview-task/gateway/endpoint"
	"github.com/gin-gonic/gin"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-openapi/runtime/middleware"
	"net/http"
)

//Transport Layer
func NewHTTPServer(endpoints endpoint.Endpoints) http.Handler {
	router := gin.Default()

	user := router.Group("/user")
	{
		// swagger:route POST /user/create user create_user
		// Create a new user and return the user_id
		// Consumes:
		//     - application/json
		// Produces:
		//     - application/json
		// responses:
		//	200: CreateUserResponse
		//  400: ErrorResponse
		user.POST("/create", gin.WrapH(httptransport.NewServer(
			endpoints.CreateUser,
			endpoint.DecodeCreateUserReq,
			endpoint.EncodeResponse,
		)))
		// swagger:route GET /user/get user get_user
		// return the user information and all user cards info
		// Consumes:
		//     - application/json
		// Produces:
		//     - application/json
		// responses:
		//	200: GetUserResponse
		//  400: ErrorResponse
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

	//swagger
	ops := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(ops, nil)
	router.GET("/docs", gin.WrapH(sh))
	router.GET("/swagger.yaml", gin.WrapH(http.FileServer(http.Dir("./"))))

	return router
}
