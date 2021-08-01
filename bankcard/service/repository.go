package service

import (
	"context"
	"github.com/amiranmanesh/imenaria-interview-task/db/sql"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"gorm.io/gorm"
)

func NewCardRepository(db *gorm.DB, logger log.Logger) IRepository {
	if err := db.AutoMigrate(sql.BankCard{}); err != nil {
		level.Error(logger).Log("Repository auto migration failed", err)
		panic(err)
	}
	return &repository{db, log.With(logger, "Repository")}
}

type repository struct {
	db     *gorm.DB
	logger log.Logger
}

func (r repository) Create(ctx context.Context, bankName, bankCardNumber string, userID uint) (uint, error) {
	//TODO: handle ctx

	logger := log.With(r.logger, "Create")
	logger.Log("Start creating bank card object for user id %d", userID)

	card := &sql.BankCard{}
	card.BankName = bankName
	card.CardNumber = bankCardNumber
	card.UserID = userID

	cid, err := card.Save(r.db)
	if err != nil {
		level.Error(logger).Log("Error is: ", err)
		return cid, err
	}

	logger.Log("Bank Card created with id: ", cid)
	return cid, nil
}

func (r repository) GetCardInfoByCardID(ctx context.Context, cardId uint) (*BankCardModel, error) {
	//TODO: handle ctx

	logger := log.With(r.logger, "Get Card Info By Card ID")
	logger.Log("Start getting bank card object with card id %d", cardId)

	card := &sql.BankCard{}
	card.ID = cardId

	if err := card.FindByCardID(r.db); err != nil {
		level.Error(logger).Log("Error is: ", err)
		return nil, err
	}

	cardInfo := &BankCardModel{}
	cardInfo.CardID = card.ID
	cardInfo.UserID = card.UserID
	cardInfo.BankName = card.BankName
	cardInfo.CardNumber = card.CardNumber

	logger.Log("Card found")
	return cardInfo, nil
}

func (r repository) GetCardsByUserID(ctx context.Context, userId uint) ([]*BankCardModel, error) {
	//TODO: handle ctx

	logger := log.With(r.logger, "Get Cards Info By User ID")
	logger.Log("Start finding bank card objects with user id %d", userId)

	card := sql.BankCard{}
	cards, err := card.GetAllByUserID(r.db)

	if err != nil {
		level.Error(logger).Log("Error is: ", err)
		return nil, err
	}

	var result []*BankCardModel

	for _, card := range cards {
		temp := &BankCardModel{}
		temp.UserID = userId
		temp.CardID = card.ID
		temp.CardNumber = card.CardNumber
		temp.BankName = card.BankName

		result = append(result, temp)
	}

	logger.Log("Found ", len(result), " Items")
	return result, nil
}
