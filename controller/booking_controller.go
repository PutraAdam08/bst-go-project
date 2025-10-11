package controller

import (
	"fmt"
	"net/http"

	"BSTproject.com/dto"
	"BSTproject.com/model"
	apix "BSTproject.com/utils/api"
	validatorx "BSTproject.com/utils/validator"

	"github.com/gin-gonic/gin"
)

type BookingService interface {
	GetAll() ([]model.Booking, error)
	GetByID(id uint) (*model.Booking, error)
	Create(booking *model.Booking) error
	Update(booking *model.Booking) error
	Delete(id uint) error
	UpdateBookingStatus(userId uint, id uint, status int) error
}

type BookingController struct {
	bookingService BookingService
}

func NewBookingController(bookingService BookingService) *BookingController {
	return &BookingController{
		bookingService: bookingService,
	}
}

// GetAllBookings
// @Summary Get all bookings
// @Description Retrieve all bookings
// @Tags booking
// @Accept json
// @Produce json
// @Security BearerToken
// @Router /bookings [GET]
func (c *BookingController) GetAllBookings(ctx *gin.Context) {
	bookings, err := c.bookingService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, apix.HTTPResponse{
			Message: "failed to fetch bookings",
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, apix.HTTPResponse{
		Message: "successfully fetched bookings",
		Data:    bookings,
	})
}

// GetBookingByID
// @Summary Get booking by ID
// @Description Retrieve a single booking by its ID
// @Tags booking
// @Accept json
// @Produce json
// @Param id path int true "Booking ID"
// @Security BearerToken
// @Router /bookings/{id} [GET]
func (c *BookingController) GetBookingByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	var id uint
	_, err := fmt.Sscan(idParam, &id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "invalid booking ID",
		})
		return
	}

	booking, err := c.bookingService.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, apix.HTTPResponse{
			Message: "booking not found",
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, apix.HTTPResponse{
		Message: "successfully fetched booking",
		Data:    booking,
	})
}

// CreateBooking
// @Summary Create booking
// @Description Create a new booking
// @Tags booking
// @Param Booking body dto.CreateBookingDTO true "Booking Data"
// @Accept json
// @Produce json
// @Security BearerToken
// @Router /bookings [POST]
func (c *BookingController) CreateBooking(ctx *gin.Context) {
	var bookingDTO dto.CreateBookingDTO
	err := ctx.ShouldBindJSON(&bookingDTO)

	userID := ctx.GetInt("user_id")
	var statusPending = 0
	booking := model.Booking{
		Status:    statusPending,
		TimeStart: bookingDTO.TimeStart,
		TimeEnd:   bookingDTO.TimeEnd,
		UserId:    uint(userID),
		RoomId:    bookingDTO.RoomId,
	}

	err = c.bookingService.Create(&booking)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, apix.HTTPResponse{
			Message: "failed to create booking",
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, apix.HTTPResponse{
		Message: "successfully created booking",
		Data:    booking,
	})
}

// UpdateBooking
// @Summary Update booking
// @Description Update an existing booking
// @Tags booking
// @Param id path int true "Booking ID"
// @Param Booking body dto.UpdateBookingDTO true "Updated Booking Data"
// @Accept json
// @Produce json
// @Security BearerToken
// @Router /bookings/{id} [PUT]
func (c *BookingController) UpdateBooking(ctx *gin.Context) {
	idParam := ctx.Param("id")
	var id uint
	_, err := fmt.Sscan(idParam, &id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "invalid booking ID",
		})
		return
	}

	var updateDTO dto.UpdateBookingDTO
	err = ctx.ShouldBindJSON(&updateDTO)
	if err != nil {
		ve, _ := validatorx.ParseValidatorErrors(err)
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "invalid input data",
			Data:    ve,
		})
		return
	}

	booking := model.Booking{
		Id:        id,
		Status:    updateDTO.Status,
		TimeStart: updateDTO.TimeStart,
		TimeEnd:   updateDTO.TimeEnd,
		UserId:    updateDTO.UserId,
		RoomId:    updateDTO.RoomId,
	}

	err = c.bookingService.Update(&booking)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, apix.HTTPResponse{
			Message: "failed to update booking",
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, apix.HTTPResponse{
		Message: "successfully updated booking",
		Data:    booking,
	})
}

// DeleteBooking
// @Summary Delete booking
// @Description Delete a booking by ID
// @Tags booking
// @Param id path int true "Booking ID"
// @Accept json
// @Produce json
// @Security BearerToken
// @Router /bookings/{id} [DELETE]
func (c *BookingController) DeleteBooking(ctx *gin.Context) {
	idParam := ctx.Param("id")
	var id uint
	_, err := fmt.Sscan(idParam, &id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "invalid booking ID",
		})
		return
	}

	err = c.bookingService.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, apix.HTTPResponse{
			Message: "failed to delete booking",
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, apix.HTTPResponse{
		Message: "successfully deleted booking",
	})
}

// UpdateBookingStatus
// @Summary Approve or Reject a booking
// @Description Update a booking’s status (0=pending, 1=approved, 2=rejected, 3=cancelled)
// @Tags booking
// @Param id path int true "Booking ID"
// @Param status query int true "New Status"
// @Accept json
// @Produce json
// @Security BearerToken
// @Router /bookings/{id}/status [PATCH]
func (c *BookingController) UpdateBookingStatus(ctx *gin.Context) {
	idParam := ctx.Param("id")
	var id uint
	_, err := fmt.Sscan(idParam, &id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "invalid booking ID",
		})
		return
	}

	statusParam := ctx.Query("status")
	var status int
	_, err = fmt.Sscan(statusParam, &status)
	if err != nil || status < 0 || status > 3 {
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "invalid status value (allowed: 0-3)",
		})
		return
	}

	userID := ctx.GetInt("user_id")
	err = c.bookingService.UpdateBookingStatus(uint(userID), id, status)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "failed to update booking status",
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, apix.HTTPResponse{
		Message: fmt.Sprintf("booking status updated to %d", status),
	})
}
