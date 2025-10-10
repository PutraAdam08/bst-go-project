package dto

type CreateRoomDTO struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Capacity    uint   `json:"capacity" binding:"required"`
}

type UpdateRoomDTO struct {
	Name        string `json:"name" binding:"omitempty"`
	Description string `json:"description" binding:"omitempty"`
	Capacity    uint   `json:"capacity" binding:"omitempty"`
}
