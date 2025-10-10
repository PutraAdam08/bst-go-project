package service

import (
	"errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"BSTproject.com/model"
	"BSTproject.com/utils/auth"

	"github.com/golang-jwt/jwt/v4"
)

type UserRepository interface {
	Create(user *model.User) error
	Update(user *model.User) error
	GetByID(id uint) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
}

type JWTService interface {
	GenerateToken(userID uint) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
	GetUserByTokenID(token string) (uint, error)
}

type userService struct {
	jwtService     JWTService
	userRepository UserRepository
}

func NewUserService(jwtService JWTService, userRepository UserRepository) *userService {
	return &userService{
		jwtService:     jwtService,
		userRepository: userRepository,
	}
}

func (s *userService) Register(user *model.User) (*model.User, error) {
	existing, err := s.userRepository.GetByEmail(user.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if existing != nil {
		return nil, errors.New("user with this email has already exist")
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Info("[Register, GenerateFromPassword] error: ", err)
		return nil, err
	}

	user.Password = string(passHash)
	user.IsAdmin = false

	err = s.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (s *userService) Login(user *model.User) (string, error) {
	existing, err := s.userRepository.GetByEmail(user.Email)
	if err != nil {
		return "", err
	}

	if existing == nil {
		return "", errors.New("user does not exist")
	}

	res, err := auth.ComparePassword(existing.Password, []byte(user.Password))
	if err != nil {
		return "", err
	}

	if !res {
		return "", errors.New("password does not match")
	}

	token, err := s.jwtService.GenerateToken(existing.Id)
	if err != nil {
		return "", err
	}

	return token, err
}

func (s *userService) AdminLogin(user *model.User) (string, error) {
	existing, err := s.userRepository.GetByEmail(user.Email)
	if err != nil {
		return "", err
	}

	if existing == nil {
		return "", errors.New("user does not exist")
	}

	if !existing.IsAdmin {
		return "", errors.New("user is not admin")
	}

	res, err := auth.ComparePassword(existing.Password, []byte(user.Password))
	if err != nil {
		return "", err
	}

	if !res {
		return "", errors.New("password does not match")
	}

	token, err := s.jwtService.GenerateToken(existing.Id)
	if err != nil {
		return "", err
	}

	return token, err
}

func (s *userService) GetByID(id uint) (*model.User, error) {
	return s.userRepository.GetByID(id)
}

func (s *userService) Update(user *model.User) (*model.User, error) {
	passHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Info("[Update, GenerateFromPassword] error: ", err)
		return nil, err
	}

	user.Password = string(passHash)

	err = s.userRepository.Update(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}
