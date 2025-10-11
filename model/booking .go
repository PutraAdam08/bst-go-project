package model

import (
	"time"
)

type Booking struct {
	Id        uint      `json:"id" gorm:"primary_key:auto_increment"`
	Status    int       `json:"status"`
	TimeStart time.Time `json:"time_start"`
	TimeEnd   time.Time `json:"time_end"`
	UserId    uint      `json:"user_id"`
	RoomId    uint      `json:"room_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Room Room `json:"room,omitempty" gorm:"foreignKey:RoomId;references:Id"`
	User User `json:"user,omitempty" gorm:"foreignKey:UserId;references:Id"`
}
