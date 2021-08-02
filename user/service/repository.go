package service

import (
	"context"
	"github.com/amiranmanesh/imenaria-interview-task/db/sql"
	"github.com/amiranmanesh/imenaria-interview-task/user/proto"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"gorm.io/gorm"
)

func NewUserRepository(db *gorm.DB, logger log.Logger) IRepository {
	if err := db.AutoMigrate(&sql.User{}); err != nil {
		level.Error(logger).Log("Repository auto migration failed", err)
		panic(err)
	}
	return &repository{db, log.With(logger, "Repository")}
}

type repository struct {
	db     *gorm.DB
	logger log.Logger
}

func (r repository) Create(ctx context.Context, userInfo proto.UserInfo) (uint, error) {
	//TODO: handle ctx

	logger := log.With(r.logger, "Create")
	logger.Log("Start creating user object")

	user := &sql.User{}
	user.Name = userInfo.Name
	user.Gender = userInfo.Gender
	user.BirthYear = int(userInfo.BirthYear)
	user.Avatar = userInfo.Avatar

	uid, err := user.Save(r.db)
	if err != nil {
		level.Error(logger).Log("Error is: ", err)
		return uid, err
	}

	logger.Log("User created with id: ", uid)
	return uid, nil
}

func (r repository) Update(ctx context.Context, userId uint, userInfo proto.UserInfo) error {
	//TODO: handle ctx

	logger := log.With(r.logger, "Update")
	logger.Log("Start updating user object with user id", userId)

	user := &sql.User{}
	user.ID = userId

	if err := user.Find(r.db); err != nil {
		level.Error(logger).Log("Error is: ", err)
		return err
	}

	if userInfo.Name != "" {
		user.Name = userInfo.Name
	}
	if userInfo.Gender != "" {
		user.Gender = userInfo.Gender
	}
	if userInfo.BirthYear != 0 {
		user.BirthYear = int(userInfo.BirthYear)
	}
	if userInfo.Avatar != "" {
		user.Avatar = userInfo.Avatar
	}

	if err := user.Update(r.db); err != nil {
		level.Error(logger).Log("Error is: ", err)
		return err
	}

	logger.Log("User updated successfully with user id: ", userId)
	return nil
}

func (r repository) Delete(ctx context.Context, userId uint) error {
	//TODO: handle ctx

	logger := log.With(r.logger, "Delete")
	logger.Log("Start deleting user object with user id", userId)

	user := &sql.User{}
	user.ID = userId

	if err := user.Delete(r.db); err != nil {
		level.Error(logger).Log("Error is: ", err)
		return err
	}

	logger.Log("User deleted successfully with user id: ", userId)
	return nil
}

func (r repository) Get(ctx context.Context, userId uint) (*proto.UserInfo, error) {
	//TODO: handle ctx

	logger := log.With(r.logger, "Get")
	logger.Log("Start getting user object with user id: ", userId)

	userInfo := &proto.UserInfo{}

	user := &sql.User{}
	user.ID = userId

	if err := user.Find(r.db); err != nil {
		level.Error(logger).Log("Error is: ", err)
		return userInfo, err
	}

	userInfo.Name = user.Name
	userInfo.Gender = user.Gender
	userInfo.BirthYear = int32(user.BirthYear)
	userInfo.Avatar = user.Avatar

	logger.Log("User found")
	return userInfo, nil
}

func (r repository) Verify(ctx context.Context, userId uint) error {
	//TODO: handle ctx

	logger := log.With(r.logger, "Verify")
	logger.Log("Start verifying user object with user id: ", userId)

	user := &sql.User{}
	user.ID = userId

	if err := user.Find(r.db); err != nil {
		level.Error(logger).Log("Error is: ", err)
		return err
	}

	logger.Log("User verified")
	return nil
}
