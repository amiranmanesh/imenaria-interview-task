package service

import (
	"context"
	"github.com/amiranmanesh/imenaria-interview-task/bankcard/proto"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"gorm.io/gorm"

	"github.com/amiranmanesh/imenaria-interview-task/db/sql"
)

func NewCardRepository(db *gorm.DB, logger log.Logger) IRepository {
	if err := db.AutoMigrate(&sql.BankCard{}); err != nil {
		level.Error(logger).Log("Repository auto migration failed", err)
		panic(err)
	}
	return &repository{db, log.With(logger, "Repository")}
}

type repository struct {
	db     *gorm.DB
	logger log.Logger
}

func (r repository) Create(ctx context.Context, bankName, cardNumber string, userID uint) (uint, error) {
	//TODO: handle ctx

	logger := log.With(r.logger, "Create")
	logger.Log("Start creating bank card object for user id %d", userID)

	card := &sql.BankCard{}
	card.BankName = bankName
	card.CardNumber = cardNumber
	card.UserID = userID

	cardId, err := card.Save(r.db)
	if err != nil {
		level.Error(logger).Log("Error is: ", err)
		return cardId, err
	}

	logger.Log("Bank Card created successfully with card id: ", cardId)
	return cardId, nil
}

func (r repository) Update(ctx context.Context, cardInfo *proto.CardInfo) error {
	//TODO: handle ctx

	logger := log.With(r.logger, "Update")
	logger.Log("Start updating bank card object with card id", cardInfo.CardId)

	card := &sql.BankCard{}
	card.ID = uint(cardInfo.CardId)

	if err := card.Find(r.db); err != nil {
		level.Error(logger).Log("Error is: ", err)
		return err
	}

	if cardInfo.BankName != "" {
		card.BankName = cardInfo.BankName
	}
	if cardInfo.CardNumber != "" {
		card.CardNumber = cardInfo.CardNumber
	}

	if err := card.Update(r.db); err != nil {
		level.Error(logger).Log("Error is: ", err)
		return err
	}

	logger.Log("Bank Card updated successfully with card id: ", cardInfo.CardId)
	return nil
}

func (r repository) Delete(ctx context.Context, cardId uint) error {
	//TODO: handle ctx

	logger := log.With(r.logger, "Delete")
	logger.Log("Start deleting bank card object with card id", cardId)

	card := &sql.BankCard{}
	card.ID = cardId

	if err := card.Delete(r.db); err != nil {
		level.Error(logger).Log("Error is: ", err)
		return err
	}

	logger.Log("Bank Card deleted successfully with card id: ", cardId)
	return nil

}

func (r repository) Get(ctx context.Context, cardId uint) (*proto.CardInfoFull, error) {
	//TODO: handle ctx

	logger := log.With(r.logger, "Get Card Info By Card ID")
	logger.Log("Start getting bank card object with card id %d", cardId)

	cardInfo := &proto.CardInfoFull{}

	card := &sql.BankCard{}
	card.ID = cardId

	if err := card.Find(r.db); err != nil {
		level.Error(logger).Log("Error is: ", err)
		return cardInfo, err
	}

	cardInfo.CardId = int32(card.ID)
	cardInfo.UserId = int32(card.UserID)
	cardInfo.BankName = card.BankName
	cardInfo.CardNumber = card.CardNumber

	logger.Log("Card found successfully")
	return cardInfo, nil
}

func (r repository) GetAll(ctx context.Context, userId uint) ([]*proto.CardInfo, error) {
	//TODO: handle ctx

	logger := log.With(r.logger, "Get all cards")
	logger.Log("Start finding bank card objects with user id %d", userId)

	card := sql.BankCard{}
	card.UserID = userId
	cards, err := card.GetAllByUserID(r.db)

	if err != nil {
		level.Error(logger).Log("Error is: ", err)
		return nil, err
	}

	var result []*proto.CardInfo

	for _, card := range cards {
		temp := &proto.CardInfo{}
		temp.CardId = int32(card.ID)
		temp.CardNumber = card.CardNumber
		temp.BankName = card.BankName

		result = append(result, temp)
	}

	logger.Log("Found ", len(result), " Items")
	return result, nil
}
