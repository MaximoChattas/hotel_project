package service

import (
	"project/dto"
)

type userService struct{}

type userServiceInterface interface {
	GetUserById(id int) (dto.UserDto, error)
	GetUsers() (dto.UsersDto, error)
	InsertUser(userDto dto.UserDto) (dto.UserDto, error)
}
