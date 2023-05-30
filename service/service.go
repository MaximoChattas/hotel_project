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
	UserLogin(loginDto dto.UserLoginDto) (bool, error)

	GetHotelById(id int) (dto.HotelDto, error)
	GetHotels() (dto.HotelsDto, error)
	InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, error)

	InsertReservation(reservationDto dto.ReservationDto) (dto.ReservationDto, error)
	GetReservationById(id int) (dto.ReservationDto, error)
	GetReservations() (dto.ReservationsDto, error)
	GetReservationsByUser(userId int) (dto.UserReservationsDto, error)
	GetReservationsByHotel(hotelId int) (dto.HotelReservationsDto, error)

	CheckAvailability(hotelId int, startDate time.Time, endDate time.Time) bool
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
	user.Role = "Customer"

	user = client.InsertUser(user)

	userDto.Id = user.Id

	return userDto, nil
}

func (s *service) GetUserById(id int) (dto.UserDto, error) {

	var user model.User = client.GetUserById(id)

	var userDto dto.UserDto

	if user.Id == 0 {
		return userDto, errors.New("user not found")
	}

	userDto.Id = user.Id
	userDto.Name = user.Name
	userDto.LastName = user.LastName
	userDto.Dni = user.Dni
	userDto.Email = user.Email

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

		usersDto = append(usersDto, userDto)
	}

	return usersDto, nil
}

func (s *service) UserLogin(loginDto dto.UserLoginDto) (bool, error) {

	var user = client.GetUserByEmail(loginDto.Email)

	if user.Id == 0 {
		return false, errors.New("user not found")
	}

	if user.Password != loginDto.Password {
		return false, errors.New("incorrect password")
	}

	return true, nil
}

func (s *service) InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, error) {
	var hotel model.Hotel

	hotel.Name = hotelDto.Name
	hotel.Description = hotelDto.Description
	hotel.RoomAmount = hotelDto.RoomAmount
	hotel.StreetName = hotelDto.StreetName
	hotel.StreetNumber = hotelDto.StreetNumber

	hotel = client.InsertHotel(hotel)

	hotelDto.Id = hotel.Id

	return hotelDto, nil
}

func (s *service) GetHotels() (dto.HotelsDto, error) {

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

func (s *service) GetHotelById(id int) (dto.HotelDto, error) {

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
	hotelDto.StreetNumber = hotel.StreetNumber

	return hotelDto, nil
}

func (s *service) InsertReservation(reservationDto dto.ReservationDto) (dto.ReservationDto, error) {

	timeStart, _ := time.Parse("02-01-2006 15:04", reservationDto.StartDate)
	timeEnd, _ := time.Parse("02-01-2006 15:04", reservationDto.EndDate)

	if timeStart.After(timeEnd) {
		return reservationDto, errors.New("a reservation cant end before it starts")
	}

	if s.CheckAvailability(reservationDto.HotelId, timeStart, timeEnd) {
		var reservation model.Reservation

		reservation.StartDate = reservationDto.StartDate
		reservation.EndDate = reservationDto.EndDate
		reservation.HotelId = reservationDto.HotelId
		reservation.UserId = reservationDto.UserId

		reservation = client.InsertReservation(reservation)

		reservationDto.Id = reservation.Id

		return reservationDto, nil
	}

	return reservationDto, errors.New("there are no rooms available")
}

func (s *service) GetReservationById(id int) (dto.ReservationDto, error) {
	var reservation model.Reservation
	var reservationDto dto.ReservationDto

	reservation = client.GetReservationById(id)

	if reservation.Id == 0 {
		return reservationDto, errors.New("reservation not found")
	}

	reservationDto.Id = reservation.Id
	reservationDto.StartDate = reservation.StartDate
	reservationDto.EndDate = reservation.EndDate
	reservationDto.HotelId = reservation.HotelId
	reservationDto.UserId = reservation.UserId

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

func (s *service) GetReservationsByUser(userId int) (dto.UserReservationsDto, error) {
	var user model.User = client.GetUserById(userId)
	var userReservationsDto dto.UserReservationsDto
	var reservationsDto dto.ReservationsDto

	if user.Id == 0 {
		return userReservationsDto, errors.New("user not found")
	}
	var reservations model.Reservations = client.GetReservationsByUser(userId)

	userReservationsDto.UserId = user.Id
	userReservationsDto.UserName = user.Name
	userReservationsDto.UserLastName = user.LastName
	userReservationsDto.UserDni = user.Dni
	userReservationsDto.UserEmail = user.Email
	userReservationsDto.UserPassword = user.Password

	for _, reservation := range reservations {
		var reservationDto dto.ReservationDto

		reservationDto.Id = reservation.Id
		reservationDto.StartDate = reservation.StartDate
		reservationDto.EndDate = reservation.EndDate
		reservationDto.HotelId = reservation.HotelId
		reservationDto.UserId = reservation.UserId

		reservationsDto = append(reservationsDto, reservationDto)
	}

	userReservationsDto.Reservations = reservationsDto

	return userReservationsDto, nil
}

func (s *service) GetReservationsByHotel(hotelId int) (dto.HotelReservationsDto, error) {
	var hotel model.Hotel = client.GetHotelById(hotelId)
	var hotelReservations dto.HotelReservationsDto
	var reservationsDto dto.ReservationsDto

	if hotel.Id == 0 {
		return hotelReservations, errors.New("hotel not found")
	}

	var reservations model.Reservations = client.GetReservationsByHotel(hotelId)

	hotelReservations.HotelId = hotel.Id
	hotelReservations.HotelName = hotel.Name
	hotelReservations.HotelDescription = hotel.Description
	hotelReservations.HotelRoomAmount = hotel.RoomAmount
	hotelReservations.HotelStreetName = hotel.StreetName
	hotelReservations.HotelStreetNumber = hotel.StreetNumber

	for _, reservation := range reservations {
		var reservationDto dto.ReservationDto
		reservationDto.Id = reservation.Id
		reservationDto.StartDate = reservation.StartDate
		reservationDto.EndDate = reservation.EndDate
		reservationDto.HotelId = reservation.HotelId
		reservationDto.UserId = reservation.UserId

		reservationsDto = append(reservationsDto, reservationDto)
	}

	hotelReservations.Reservations = reservationsDto

	return hotelReservations, nil
}

func (s *service) CheckAvailability(hotelId int, startDate time.Time, endDate time.Time) bool {

	hotel := client.GetHotelById(hotelId)
	reservations := client.GetReservationsByHotel(hotelId)

	roomsAvailable := hotel.RoomAmount

	for _, reservation := range reservations {

		reservationStart, _ := time.Parse("02-01-2006 15:04", reservation.StartDate)
		reservationEnd, _ := time.Parse("02-01-2006 15:04", reservation.EndDate)

		if reservationStart.After(startDate) && reservationEnd.Before(endDate) ||
			reservationStart.Before(startDate) && reservationEnd.After(startDate) ||
			reservationStart.Before(endDate) && reservationEnd.After(endDate) ||
			reservationStart.Before(startDate) && reservationEnd.After(endDate) ||
			reservationStart.Equal(startDate) || reservationEnd.Equal(endDate) {
			roomsAvailable--
		}
		if roomsAvailable == 0 {
			return false
		}
	}

	return true
}
