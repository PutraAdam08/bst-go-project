package repository

import (
	"fmt"

	"BSTproject.com/model"
	"gorm.io/gorm"
)

type RoomRepository struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) *RoomRepository {
	return &RoomRepository{
		db: db,
	}
}

func (r *RoomRepository) Create(room *model.Room) error {
	fmt.Println(room)
	tx := r.db.Create(room)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r *RoomRepository) Update(room *model.Room) error {
	tx := r.db.Where("id = ?", room.ID).Updates(room)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r *RoomRepository) Delete(id uint) error {
	tx := r.db.Where("id = ?", id).Delete(model.Room{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
