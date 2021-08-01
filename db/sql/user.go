package sql

import (
	"errors"
	"gorm.io/gorm"
)

var userCreatingFailedError = errors.New("create user failed")
var userNotFoundError = errors.New("user not found")

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(100); not null;" json:"name"`
	Gender    string `gorm:"type:varchar(20);not null;" json:"gender"`
	BirthYear int    `gorm:"not null;" json:"birth_year"`
	Avatar    string `gorm:"type:varchar(200);" json:"avatar"`
}

func (u *User) Save(db *gorm.DB) (uint, error) {
	result := db.Create(&u)
	if result.Error != nil {
		return -1, userCreatingFailedError
	}

	return u.ID, nil
}

func (u *User) Find(db *gorm.DB) error {
	if result := db.First(&u, "id = ?", u.ID); result.Error != nil {
		return userNotFoundError
	}
	return nil
}