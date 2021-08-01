package sql

import (
	"errors"
	"gorm.io/gorm"
)

var bankCardCreatingFailedError = errors.New("create bank card failed")
var bankCardUpdatingFailedError = errors.New("updating bank card failed")
var bankCardDeletingFailedError = errors.New("deleting bank card failed")
var bankCardNotFoundError = errors.New("bank card not found")
var bankCardGettingAllError = errors.New("getting all bank cards failed")

type BankCard struct {
	gorm.Model
	BankName   string `gorm:"type:varchar(100);not null;" json:"bank_name"`
	CardNumber string `gorm:"type:varchar(20);not null;" json:"card_number"`
	UserID     uint
}

func (b *BankCard) Save(db *gorm.DB) (uint, error) {
	if result := db.Create(&b); result.Error != nil {
		return 0, bankCardCreatingFailedError
	}
	return b.ID, nil
}

func (b *BankCard) Find(db *gorm.DB) error {
	if result := db.First(&b, "id = ?", b.ID); result.Error != nil {
		return bankCardNotFoundError
	}
	return nil
}

func (b *BankCard) Update(db *gorm.DB) error {
	if result := db.Save(&b); result.Error != nil {
		return bankCardUpdatingFailedError
	}
	return nil
}

func (b *BankCard) Delete(db *gorm.DB) error {
	if result := db.Delete(&b); result.Error != nil {
		return bankCardDeletingFailedError
	}
	return nil
}

func (b BankCard) GetAllByUserID(db *gorm.DB) ([]BankCard, error) {
	var bankCards []BankCard
	if result := db.Where("user_id = ?", b.UserID).Find(&bankCards); result.Error != nil {
		return nil, bankCardGettingAllError
	}
	return bankCards, nil
}
