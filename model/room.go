package model

import "time"

type Room struct {
	Id          uint      `json:"id" gorm:"primary_key:auto_increment"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Capacity    uint      `json:"capacity"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
