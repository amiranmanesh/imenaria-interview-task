package sql

import (
	"errors"
	"gorm.io/gorm"
)

var bankCardCreatingFailedError = errors.New("create bank card failed")
var bankCardNotFoundError = errors.New("bank card not found")

type BankCard struct {
	gorm.Model
	BankName   string `gorm:"type:varchar(100);not null;" json:"bank_name"`
	CardNumber string `gorm:"type:varchar(20);not null;" json:"card_number"`
	UserID     uint
	User       User
}

func (b *BankCard) Save(db *gorm.DB) error {

	if result := db.Create(&b); result.Error != nil {
		return bankCardCreatingFailedError
	}

	return nil
}

func (b *BankCard) Find(db *gorm.DB) error {
	if result := db.First(&b, "id = ?", b.ID); result.Error != nil {
		return bankCardNotFoundError
	}
	return nil
}
