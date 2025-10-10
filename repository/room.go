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

func (r *RoomRepository) GetAll() ([]model.Room, error) {
	var rooms []model.Room
	tx := r.db.Find(&rooms)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return rooms, nil
}

func (r *RoomRepository) GetById(id uint) (*model.Room, error) {
	var room *model.Room
	tx := r.db.Where("id = ?", id).First(&room)
	if tx.Error != nil {
		return nil, tx.Error
	}
	//TEST UPDATE
	return room, nil
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
	tx := r.db.Where("id = ?", room.Id).Updates(room)
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
