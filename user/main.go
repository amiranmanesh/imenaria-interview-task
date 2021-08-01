package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/amiranmanesh/imenaria-interview-task/user/endpoint"
	"github.com/amiranmanesh/imenaria-interview-task/user/proto"
	"github.com/amiranmanesh/imenaria-interview-task/user/service"
	"github.com/amiranmanesh/imenaria-interview-task/user/transport"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "user",
			"time", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "user service started")
	defer level.Info(logger).Log("msg", "user service ended")

	flag.Parse()
	ctx := context.Background()

	var srv endpoint.IService
	{
		repository := service.NewUserRepository(getDataBaseModel(), logger)
		srv = service.NewService(repository, logger)
	}

	endpoints := endpoint.MakeEndpoint(srv)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", os.Getenv("GRPC_PORT")))
		if err != nil {
			panic(fmt.Sprintf("User Service: Failed to listen %v", err))
		}
		baseServer := grpc.NewServer()
		grpcHandler := transport.NewGRPCServer(ctx, endpoints)
		proto.RegisterUserServiceServer(baseServer, grpcHandler)
		reflection.Register(baseServer)

		errs <- baseServer.Serve(lis)
	}()

	level.Error(logger).Log("exit", <-errs)

}

func getDataBaseModel() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		panic(fmt.Sprintf("User Service: Failed to load .env %v", err))
	}
	dsn := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"),
	)
	connection, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{},
	)
	if err != nil {
		panic(fmt.Sprintf("User Service: Failed to connect to db %v", err))
	}
	return connection
}
