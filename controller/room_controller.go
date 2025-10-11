package controller

import (
	"net/http"
	"strconv"

	"BSTproject.com/dto"
	"BSTproject.com/model"
	apix "BSTproject.com/utils/api"
	"github.com/gin-gonic/gin"
)

type RoomService interface {
	GetAll() ([]model.Room, error)
	GetByID(id uint) (*model.Room, error)
	Create(room *model.Room) error
	Update(room *model.Room) error
	Delete(id uint) error
}

type RoomController struct {
	roomService RoomService
}

func NewRoomController(roomService RoomService) *RoomController {
	return &RoomController{
		roomService: roomService,
	}
}

// GetAll
// @Summary Get all rooms
// @Description Retrieve all room data
// @Tags room
// @Produce json
// @Security BearerToken
// @Router /rooms [GET]
func (c *RoomController) GetAll(ctx *gin.Context) {
	rooms, err := c.roomService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "failed to get rooms",
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, apix.HTTPResponse{
		Message: "successfully retrieved rooms",
		Data:    rooms,
	})
}

// GetByID
// @Summary Get room by ID
// @Description Retrieve a specific room by ID
// @Tags room
// @Produce json
// @Security BearerToken
// @Param id path int true "Room ID"
// @Router /rooms/{id} [GET]
func (c *RoomController) GetByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "invalid room ID",
		})
		return
	}

	room, err := c.roomService.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "failed to get room",
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, apix.HTTPResponse{
		Message: "successfully retrieved room",
		Data:    room,
	})
}

// Create new room
// @Summary Create room
// @Description Add a new room
// @Tags room
// @Accept json
// @Produce json
// @Security BearerToken
// @Param data body dto.CreateRoomDTO true "Room data"
// @Router /rooms [POST]
func (c *RoomController) Create(ctx *gin.Context) {
	var roomDTO dto.CreateRoomDTO
	if err := ctx.ShouldBindJSON(&roomDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "invalid input",
			Data:    err.Error(),
		})
		return
	}

	room := model.Room{
		Name:        roomDTO.Name,
		Description: roomDTO.Description,
		Capacity:    roomDTO.Capacity,
	}

	if err := c.roomService.Create(&room); err != nil {
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "failed to create room",
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, apix.HTTPResponse{
		Message: "successfully created room",
		Data:    room,
	})
}

// Update existing room
// @Summary Update room
// @Description Update a room by ID
// @Tags room
// @Accept json
// @Produce json
// @Security BearerToken
// @Param id path int true "Room ID"
// @Param data body dto.UpdateRoomDTO true "Updated Room"
// @Router /rooms/{id} [PUT]
func (c *RoomController) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "invalid room ID",
		})
		return
	}

	var updateDTO dto.UpdateRoomDTO
	if err := ctx.ShouldBindJSON(&updateDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "invalid input",
			Data:    err.Error(),
		})
		return
	}

	room := model.Room{
		Id:          uint(id),
		Name:        updateDTO.Name,
		Description: updateDTO.Description,
		Capacity:    updateDTO.Capacity,
	}

	if err := c.roomService.Update(&room); err != nil {
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "failed to update room",
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, apix.HTTPResponse{
		Message: "successfully updated room",
		Data:    room,
	})
}

// Delete room
// @Summary Delete room
// @Description Delete a room by ID
// @Tags room
// @Produce json
// @Security BearerToken
// @Param id path int true "Room ID"
// @Router /rooms/{id} [DELETE]
func (c *RoomController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "invalid room ID",
		})
		return
	}

	if err := c.roomService.Delete(uint(id)); err != nil {
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "failed to delete room",
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, apix.HTTPResponse{
		Message: "successfully deleted room",
	})
}
