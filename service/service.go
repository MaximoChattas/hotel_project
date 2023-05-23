package service

import (
	"errors"
	"project/client"
	"project/dto"
	"project/model"
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

func (s *service) InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, error) {
	var hotel model.Hotel

	hotel.Name = hotelDto.Name
	hotel.Description = hotelDto.Description
	hotel.RoomAmount = hotelDto.RoomAmount

	hotel = client.InsertHotel(hotel)

	hotelDto.Id = hotel.Id

	return hotelDto, nil
}

func (s *hotelService) GetHotel() (dto.HotelsDto, error) {

	var hotels model.Hotels = client.GetHotels()
	var hotelsDto dto.HotelsDto

	for _, hotel := range hotels {
		var hotelDto dto.HotelDto
		hotelDto.Id = hotel.Id
		hotelDto.Name = hotel.Name
		hotelDto.RoomAmount = hotel.RoomAmount
		hotelDto.Description = hotel.Description

		hotelDto.StreetName = hotel.StreetName
		hotelDto.StreetNumber = hotel.StreetNumber

		hotelsDto = append(hotelsDto, hotelDto)
	}

	return hotelsDto, nil
}

func (s *hotelService) GetHotelById(id int) (dto.HotelDto, error) {

	var hotel model.Hotel = client.GetHotelById(id)
	var hotelDto dto.HotelDto

	if hotel.Id == 0 {
		return hotelDto, errors.New("hotel not found")
	}
	hotelDto.Id = hotel.Id
	hotelDto.Name = hotel.Name
	hotelDto.RoomAmount = hotel.RoomAmount
	hotelDto.Description = hotel.Description
	hotelDto.StreetName = hotel.StreetName
	hotelDto.StreetNumber = hotel.StreetNumberNumber

	return hotelDto, nil
}
