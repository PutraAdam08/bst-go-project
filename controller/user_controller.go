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

type UserService interface {
	Register(user *model.User) (*model.User, error)
	Login(user *model.User) (string, error)
	AdminLogin(user *model.User) (string, error)
	GetByID(id uint) (*model.User, error)
	Update(user *model.User) (*model.User, error)
}

type UserController struct {
	userService UserService
}

func NewUserController(userService UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// Register
// @Summary Register
// @Schemes
// @Description Register
// @Tags user
// @Param List body dto.RegisterUserDTO true "Form Register"
// @Accept json
// @Produce json
// @Router /users/register [POST]
func (c *UserController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterUserDTO
	err := ctx.ShouldBindJSON(&registerDTO)
	if err != nil {
		ve, _ := validatorx.ParseValidatorErrors(err)
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "input data invalid",
			Data:    ve,
		})
		return
	}

	user := model.User{
		Email:    registerDTO.Email,
		Password: registerDTO.Password,
		Name:     registerDTO.Name,
	}

	res, err := c.userService.Register(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "failed to create user",
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, apix.HTTPResponse{
		Message: "succesfully created user",
		Data:    res,
	})
}

// AdminLogin
// @Summary AdminLogin
// @Schemes
// @Description AdminLogin
// @Tags admin
// @Param List body dto.LoginUserDTO true "Form AdminLogin"
// @Accept json
// @Produce json
// @Router /admin/login [POST]
func (c *UserController) AdminLogin(ctx *gin.Context) {
	var loginDTO dto.LoginUserDTO
	err := ctx.ShouldBindJSON(&loginDTO)
	if err != nil {
		ve, _ := validatorx.ParseValidatorErrors(err)
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "input data invalid",
			Data:    ve,
		})

		return
	}

	user := model.User{
		Email:    loginDTO.Email,
		Password: loginDTO.Password,
	}

	res, err := c.userService.AdminLogin(&user)
	fmt.Println(err)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "failed to login",
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, apix.HTTPResponse{
		Message: "successfully logged in",
		Data:    res,
	})
}

// Login
// @Summary Login
// @Schemes
// @Description Login
// @Tags user
// @Param List body dto.LoginUserDTO true "Form Login"
// @Accept json
// @Produce json
// @Router /users/login [POST]
func (c *UserController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginUserDTO
	err := ctx.ShouldBindJSON(&loginDTO)
	if err != nil {
		ve, _ := validatorx.ParseValidatorErrors(err)
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "input data invalid",
			Data:    ve,
		})

		return
	}

	user := model.User{
		Email:    loginDTO.Email,
		Password: loginDTO.Password,
	}

	res, err := c.userService.Login(&user)
	fmt.Println(err)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "failed to login",
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, apix.HTTPResponse{
		Message: "successfully logged in",
		Data:    res,
	})
}

// GetUser
// @Summary GetUser
// @Schemes
// @Description GetUser
// @Tags user
// @Accept json
// @Produce json
// @Security BearerToken
// @Router /users [GET]
func (c *UserController) GetUser(ctx *gin.Context) {
	userID := ctx.GetInt("user_id")
	res, err := c.userService.GetByID(uint(userID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "failed to get user",
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, apix.HTTPResponse{
		Message: "successfully get user",
		Data:    res,
	})
}

// UpdateUser
// @Summary UpdateUser
// @Schemes
// @Description UpdateUser
// @Tags user
// @Param List body dto.RegisterUserDTO true "Form UpdateUser"
// @Accept json
// @Produce json
// @Security BearerToken
// @Router /users [PUT]
func (c *UserController) UpdateUser(ctx *gin.Context) {
	userId := ctx.GetInt("user_id")
	var updateDto dto.RegisterUserDTO
	err := ctx.ShouldBindJSON(&updateDto)

	updateModel := model.User{
		Id:       uint(userId),
		Name:     updateDto.Name,
		Email:    updateDto.Email,
		Password: updateDto.Password,
	}

	res, err := c.userService.Update(&updateModel)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "failed to update user",
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, apix.HTTPResponse{
		Message: "successfully update user",
		Data:    res,
	})
}
