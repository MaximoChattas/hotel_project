package service

import (
	"project/dto"
	"time"
)

type service struct{}

type serviceInterface interface {
	GetUserById(id int) (dto.UserDto, error)
	GetUsers() (dto.UsersDto, error)
	InsertUser(userDto dto.UserDto) (dto.UserDto, error)

	GetHotelById(id int) (dto.HotelDto, error)
	GetHotels() (dto.HotelsDto, error)
	InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, error)

	InsertReservation(reservationDto dto.ReservationDto) (dto.ReservationDto, error)
	GetReservations() (dto.ReservationsDto, error)
	GetReservationsByUser(userId int) (dto.ReservationsDto, error)
	GetReservationsByHotel(hotelId int) (dto.ReservationsDto, error)

	CheckAvailability(hotelId int, startDate time.Time, endDate time.Time) bool
}

var (
	Service serviceInterface
)

func init() {
	Service = &service{}
}
