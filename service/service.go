package service

import (
	"project/dto"
)

type service struct{}

type serviceInterface interface {
	GetUserById(id int) (dto.UserDto, error)
	GetUsers() (dto.UsersDto, error)
	InsertUser(userDto dto.UserDto) (dto.UserDto, error)
	GetHotelById(id int) (dto.HotelDto, error)
	GetHotels() (dto.HotelsDto, error)
	InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, error)
}

var (
	Service serviceInterface
)

func init() {
	Service = &service{}
}
