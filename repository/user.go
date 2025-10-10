package repository

import (
	"errors"
	"fmt"

	"BSTproject.com/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (usr *UserRepository) GetByID(id uint) (*model.User, error) {
	var user model.User
	tx := usr.db.Where("id = ?", id).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}

func (usr *UserRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User
	tx := usr.db.Where("email = ?", email).First(&user)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		fmt.Printf("[GetByEmail] error: %s", tx.Error)
		return nil, tx.Error
	}

	return &user, nil
}

func (usr *UserRepository) Create(user *model.User) error {
	tx := usr.db.Create(user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (usr *UserRepository) Update(user *model.User) error {
	tx := usr.db.Where("id = ?", user.Id).Updates(user)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
