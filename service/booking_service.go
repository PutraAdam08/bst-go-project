package service

import (
	"BSTproject.com/model"
	"errors"
	"time"
)

type BookingRepository interface {
	GetAll() ([]model.Booking, error)
	GetById(id uint) (*model.Booking, error)
	Create(booking *model.Booking) error
	Update(booking *model.Booking) error
	Delete(id uint) error
	GetBookingsByRoomId(roomId uint) ([]model.Booking, error)
}

type bookingService struct {
	jwtService        JWTService
	bookingRepository BookingRepository
}

func NewBookingService(bookingRepository BookingRepository) *bookingService {
	return &bookingService{
		bookingRepository: bookingRepository,
	}
}

func (s *bookingService) GetAll() ([]model.Booking, error) {
	return s.bookingRepository.GetAll()
}

func (s *bookingService) GetByID(id uint) (*model.Booking, error) {
	return s.bookingRepository.GetById(id)
}

func (s *bookingService) Create(booking *model.Booking) error {
	if booking.TimeStart.After(booking.TimeEnd) {
		return errors.New("time_start cannot be after time_end")
	}

	isAvailable, err := s.IsRoomAvailable(booking.RoomId, booking.TimeStart, booking.TimeEnd)
	if err != nil {
		return err
	}
	if !isAvailable {
		return errors.New("room is not available during the selected time")
	}

	return s.bookingRepository.Create(booking)
}

func (s *bookingService) Update(booking *model.Booking) error {
	err := s.bookingRepository.Update(booking)
	if err != nil {
		return err
	}

	return err
}

func (s *bookingService) Delete(id uint) error {
	err := s.bookingRepository.Delete(id)
	if err != nil {
		return err
	}

	return err
}

func (s *bookingService) IsRoomAvailable(roomID uint, start, end time.Time) (bool, error) {
	bookings, err := s.bookingRepository.GetBookingsByRoomId(roomID)
	if err != nil {
		return false, err
	}

	for _, b := range bookings {
		// (start < existing_end) && (end > existing_start)
		if start.Before(b.TimeEnd) && end.After(b.TimeStart) {
			return false, nil
		}
	}

	return true, nil
}

func (s *bookingService) UpdateBookingStatus(id uint, status int) error {
	booking, err := s.bookingRepository.GetById(id)
	if err != nil {
		return err
	}

	// Only pending bookings can be modified
	if booking.Status != 0 {
		return errors.New("only pending bookings can be modified")
	}

	booking.Status = status
	return s.bookingRepository.Update(booking)
}
