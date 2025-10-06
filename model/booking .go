package model

import (
	"time"
)

type Booking struct {
	ID        uint      `json:"id"`
	Status    string    `json:"status"`
	Date      string    `json:"date"`
	TimeStart time.Time `json:"time_start"`
	TimeEnd   time.Time `json:"time_end"`
	UserID    uint      `json:"user_id" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:cascade;OnDelete:cascade"`
	RoomID    uint      `json:"room_id"`
	Room      *Room     `json:"room" gorm:"foreignKey:RoomID;references:ID;constraint:OnUpdate:cascade;OnDelete:Cascade;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
