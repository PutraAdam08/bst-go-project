package service

import "BSTproject.com/model"

type RoomRepository interface {
	GetAll() ([]model.Room, error)
	GetById(id uint) (*model.Room, error)
	Create(room *model.Room) error
	Update(room *model.Room) error
	Delete(id uint) error
}

type roomService struct {
	roomRepository RoomRepository
}

func NewRoomService(roomRepository RoomRepository) *roomService {
	return &roomService{
		roomRepository: roomRepository,
	}
}

func (s *roomService) GetAll() ([]model.Room, error) {
	return s.roomRepository.GetAll()
}

func (s *roomService) GetByID(id uint) (*model.Room, error) {
	return s.roomRepository.GetById(id)
}

func (s *roomService) Create(room *model.Room) error {
	return s.roomRepository.Create(room)
}

func (s *roomService) Update(room *model.Room) error {
	return s.roomRepository.Update(room)
}

func (s *roomService) Delete(id uint) error {
	return s.roomRepository.Delete(id)
}
