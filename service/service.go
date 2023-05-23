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
	InsertUser(userDto dto.UserDto) (dto.UserDto, error)
	GetUserById(id int) (dto.UserDto, error)
	GetUsers() (dto.UsersDto, error)

	GetHotelById(id int) (dto.HotelDto, error)
	GetHotels() (dto.HotelsDto, error)
	InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, error)

	InsertReservation(reservationDto dto.ReservationDto) (dto.ReservationDto, error)
	GetReservations() (dto.ReservationsDto, error)
	GetReservationsByUser(userId int) (dto.ReservationsDto, error)   //To do
	GetReservationsByHotel(hotelId int) (dto.ReservationsDto, error) //To do

	CheckAvailability(hotelId int, startDate time.Time, endDate time.Time) bool //To do
}

var (
	Service serviceInterface
)

func init() {
	Service = &service{}
}

func (s *service) InsertUser(userDto dto.UserDto) (dto.UserDto, error) {
	var user model.User

	user.Name = userDto.Name
	user.LastName = userDto.LastName
	user.Dni = userDto.Dni
	user.Email = userDto.Email
	user.Password = userDto.Password

	user = client.InsertUser(user)

	userDto.Id = user.Id

	return userDto, nil
}

func (s *service) GetUserById(id int) (dto.UserDto, error) {

	var user model.User = client.GetUserById(id)

	var userDto dto.UserDto

	if user.Id == 0 {
		return userDto, errors.New("User not found")
	}

	userDto.Id = user.Id
	userDto.Name = user.Name
	userDto.LastName = user.LastName
	userDto.Dni = user.Dni
	userDto.Email = user.Email
	userDto.Password = user.Password

	return userDto, nil
}

func (s *service) GetUsers() (dto.UsersDto, error) {
	var users model.Users = client.GetUsers()
	var usersDto dto.UsersDto

	for _, user := range users {
		var userDto dto.UserDto
		userDto.Id = user.Id
		userDto.Name = user.Name
		userDto.LastName = user.LastName
		userDto.Dni = user.Dni
		userDto.Email = user.Email
		userDto.Password = user.Password

		usersDto = append(usersDto, userDto)
	}

	return usersDto, nil
}

func (s *service) InsertReservation(reservationDto dto.ReservationDto) (dto.ReservationDto, error) {
	var reservation model.Reservation

	reservation.StartDate = reservationDto.StartDate
	reservation.EndDate = reservationDto.EndDate
	reservation.HotelId = reservationDto.HotelId
	reservation.UserId = reservationDto.UserId

	reservation = client.InsertReservation(reservation)

	reservationDto.Id = reservation.Id

	return reservationDto, nil
}

func (s *service) GetReservations() (dto.ReservationsDto, error) {

	var reservations model.Reservations = client.GetReservations()
	var reservationsDto dto.ReservationsDto

	for _, reservation := range reservations {
		var reservationDto dto.ReservationDto

		reservationDto.Id = reservation.Id
		reservationDto.StartDate = reservation.StartDate
		reservationDto.EndDate = reservation.EndDate
		reservationDto.HotelId = reservation.HotelId
		reservationDto.UserId = reservation.UserId

		reservationsDto = append(reservationsDto, reservationDto)
	}

	return reservationsDto, nil
}