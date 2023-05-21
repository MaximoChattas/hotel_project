package service

import (
	"project/dto"
)

type service struct{}

type serviceInterface interface {
	GetUserById(id int) (dto.UserDto, error)
	GetUsers() (dto.UsersDto, error)
	InsertUser(userDto dto.UserDto) (dto.UserDto, error)
}

var (
	Service serviceInterface
)

func init() {
	Service = &service{}
}
