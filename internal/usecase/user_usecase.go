package usecase

import (
	"errors"

	dtos "github.com/savanyv/zenith-pay/internal/dto"
	"github.com/savanyv/zenith-pay/internal/model"
	"github.com/savanyv/zenith-pay/internal/repository"
	"github.com/savanyv/zenith-pay/internal/utils/helpers"
)

type UserUsecase interface {
	Register(req *dtos.CreateUserRequest) (*dtos.CreateUserResponse, error)
	Login(req *dtos.LoginRequest) (*dtos.LoginResponse, error)
}

type userUsecase struct {
	userRepo repository.UserRespository
	bcrypt helpers.BcryptHelper
	jwt helpers.JWTService
}

func NewUserUsecase(userRepo repository.UserRespository, jwt helpers.JWTService) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
		jwt: jwt,
		bcrypt: helpers.NewBcryptHelper(),
	}
}

func (u *userUsecase) Register(req *dtos.CreateUserRequest) (*dtos.CreateUserResponse, error) {
	userExists, err := u.userRepo.GetByUsername(req.Username)
	if err == nil && userExists != nil {
		return nil, errors.New("username already taken")
	}

	hashPassword, err := u.bcrypt.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	user := &model.User{
		Username: req.Username,
		Password: hashPassword,
		FullName: req.FullName,
		Email: req.Email,
		Role: req.Role,
		IsActive: true,
	}

	if err := u.userRepo.Create(user); err != nil {
		return nil, errors.New("failed to create user")
	}

	res := &dtos.CreateUserResponse{
		ID: user.ID.String(),
		Username: user.Username,
		FullName: user.FullName,
		Email: user.Email,
		Role: user.Role,
		IsActive: user.IsActive,
	}

	return res, nil
}

func (u *userUsecase) Login(req *dtos.LoginRequest) (*dtos.LoginResponse, error) {
	user, err := u.userRepo.GetByUsername(req.Username)
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	if err := u.bcrypt.ComparePassword(user.Password, req.Password); err != nil {
		return nil, errors.New("invalid username or password")
	}

	token, err := u.jwt.GenerateToken(user.ID.String(), user.Username)
	if err != nil {
		return nil, errors.New("failed to generate access token")
	}

	res := &dtos.LoginResponse{
		AccessToken: token,
	}

	return res, nil
}
