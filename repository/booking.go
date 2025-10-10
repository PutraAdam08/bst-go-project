package repository

import (
	"BSTproject.com/model"
	"gorm.io/gorm"
)

type BookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) *BookingRepository {
	return &BookingRepository{
		db: db,
	}
}

func (r *BookingRepository) GetAll() ([]model.Booking, error) {
	var bookings []model.Booking
	tx := r.db.Preload("Room").Preload("User").Find(&bookings)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return bookings, nil
}

func (r *BookingRepository) GetById(id uint) (*model.Booking, error) {
	var booking *model.Booking
	tx := r.db.Preload("Room").Preload("User").Where("id = ?", id).First(&booking)
	if tx.Error != nil {
		return nil, tx.Error
	}
	//TEST UPDATE
	return booking, nil
}

func (r *BookingRepository) Create(booking *model.Booking) error {
	tx := r.db.Create(booking)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r *BookingRepository) Update(booking *model.Booking) error {
	tx := r.db.Where("id = ?", booking.Id).Updates(booking)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r *BookingRepository) Delete(id uint) error {
	tx := r.db.Where("id = ?", id).Delete(model.Booking{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r *BookingRepository) GetBookingsByRoomId(roomId uint) ([]model.Booking, error) {
	var bookings []model.Booking
	tx := r.db.Where("room_id = ?", roomId).Find(&bookings)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return bookings, nil
}
