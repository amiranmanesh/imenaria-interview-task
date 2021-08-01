package main

import (
	"flag"
	"fmt"
	cardproto "github.com/amiranmanesh/imenaria-interview-task/bankcard/proto"
	"github.com/amiranmanesh/imenaria-interview-task/gateaway/endpoint"
	"github.com/amiranmanesh/imenaria-interview-task/gateaway/service"
	gateaway "github.com/amiranmanesh/imenaria-interview-task/gateaway/transport"
	userproto "github.com/amiranmanesh/imenaria-interview-task/user/proto"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(fmt.Sprintf("Main Service: Failed to load .env %v", err))
	}

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "main",
			"time", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "main service started")
	defer level.Info(logger).Log("msg", "main service ended")

	flag.Parse()

	userGrpcConnection, err := grpc.Dial(
		fmt.Sprintf("%s:%s",
			os.Getenv("GRPC_HOST"), os.Getenv("GRPC_USER_PORT"),
		), grpc.WithInsecure(),
	)
	if err != nil {
		panic(
			fmt.Sprintf("Error connecting to User GRPC Server: %s",
				err.Error(),
			),
		)
	}

	cardGrpcConnection, err := grpc.Dial(
		fmt.Sprintf("%s:%s",
			os.Getenv("GRPC_HOST"), os.Getenv("GRPC_CARD_PORT"),
		), grpc.WithInsecure(),
	)
	if err != nil {
		panic(
			fmt.Sprintf("Error connecting to Card GRPC Server: %s",
				err.Error(),
			),
		)
	}

	var srv endpoint.IService
	{
		srv = service.NewService(
			userproto.NewUserServiceClient(userGrpcConnection),
			cardproto.NewCardServiceClient(cardGrpcConnection),
			logger,
		)
	}

	endpoints := endpoint.MakeEndpoint(srv)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		handler := gateaway.NewHTTPServer(endpoints)
		errs <- http.ListenAndServe(
			fmt.Sprintf(":%s",
				os.Getenv("HTTP_PORT"),
			), handler)
	}()

	level.Error(logger).Log("exit", <-errs)
}
