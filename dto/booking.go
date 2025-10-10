package dto

import "time"

type CreateBookingDTO struct {
	Status    int       `json:"status" binding:"required"`
	TimeStart time.Time `json:"time_start" binding:"required"`
	TimeEnd   time.Time `json:"time_end" binding:"required"`
	RoomId    uint      `json:"room_id" binding:"required"`
}

type UpdateBookingDTO struct {
	Status    int       `json:"status"`
	TimeStart time.Time `json:"time_start"`
	TimeEnd   time.Time `json:"time_end"`
	UserId    uint      `json:"user_id"`
	RoomId    uint      `json:"room_id"`
}
