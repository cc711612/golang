package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"golang_user/config"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null"`
	Username string `gorm:"not null"`
	Password string `gorm:"not null"`
}

func GetList() ([]*User, error) {
	db, err := config.ConnectDatabase()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var users []*User
	if err := db.Where("deleted_at is null").Find(&users).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.New("User not found")
		}
		return nil, err
	}

	return users, nil
}

func GetUserByID(id uint) (*User, error) {
	db, err := config.ConnectDatabase()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var user User
	if err := db.Where("id = ?", id).Where("deleted_at is null").First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.New("User not found")
		}
		return nil, err
	}

	return &user, nil
}

func (u *User) Create() error {
	db, err := config.ConnectDatabase()
	if err != nil {
		return err
	}
	defer db.Close()
	if err := db.Create(&u).Error; err != nil {
		return err
	}

	return nil
}
func (u *User) Update() error {
	db, err := config.ConnectDatabase()
	if err != nil {
		return err
	}
	defer db.Close()
	if err := db.Save(&u).Error; err != nil {
		return err
	}

	return nil
}
func DeleteUserByID(id uint) error {
	db, err := config.ConnectDatabase()
	if err != nil {
		return err
	}
	defer db.Close()
	if err := db.Delete(&User{}, id).Error; err != nil {
		return err
	}

	return nil
}
