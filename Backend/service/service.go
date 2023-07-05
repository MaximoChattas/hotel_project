package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"math"
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
	UserLogin(loginDto dto.UserDto) (dto.UserDto, error)

	GetHotelById(id int) (dto.HotelDto, error)
	GetHotels() (dto.HotelsDto, error)
	InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, error)

	InsertReservation(reservationDto dto.ReservationDto) (dto.ReservationDto, error)
	GetReservationById(id int) (dto.ReservationDto, error)
	GetReservations() (dto.ReservationsDto, error)
	GetReservationsByUser(userId int) (dto.UserReservationsDto, error)
	GetReservationsByUserRange(userId int, startDate string, endDate string) (dto.ReservationsDto, error)
	GetReservationsByHotel(hotelId int) (dto.HotelReservationsDto, error)

	InsertAmenity(amenityDto dto.AmenityDto) (dto.AmenityDto, error)
	GetAmenities() (dto.AmenitiesDto, error)

	CheckAvailability(hotelId int, startDate time.Time, endDate time.Time) bool
	CheckAllAvailability(startDate string, endDate string) (dto.HotelsDto, error)

	InsertImages(imagesDto dto.ImagesDto) (dto.ImagesDto, error)
}

var (
	Service serviceInterface
)

func init() {
	Service = &service{}
}

func (s *service) InsertUser(userDto dto.UserDto) (dto.UserDto, error) {
	var user model.User

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(userDto.Password), bcrypt.DefaultCost)

	if err != nil {
		return userDto, err
	}

	user.Name = userDto.Name
	user.LastName = userDto.LastName
	user.Dni = userDto.Dni
	user.Email = userDto.Email
	user.Password = string(encryptedPassword)
	user.Role = "Customer"

	user = client.InsertUser(user)

	userDto.Id = user.Id
	userDto.Role = user.Role
	userDto.Password = user.Password

	if user.Id == 0 {
		return userDto, errors.New("error creating user")
	}

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
	userDto.Role = user.Role

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
		userDto.Role = user.Role

		usersDto = append(usersDto, userDto)
	}

	return usersDto, nil
}

func (s *service) UserLogin(loginDto dto.UserDto) (dto.UserDto, error) {

	var user = client.GetUserByEmail(loginDto.Email)

	if user.Id == 0 {
		return loginDto, errors.New("user not registered")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDto.Password))
	if err != nil {
		// Passwords don't match
		return loginDto, errors.New("incorrect password")
	}

	var userDto dto.UserDto

	userDto.Id = user.Id
	userDto.Name = user.Name
	userDto.LastName = user.LastName
	userDto.Dni = user.Dni
	userDto.Email = user.Email
	userDto.Role = user.Role
	return userDto, nil
}

func (s *service) InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, error) {
	var hotel model.Hotel

	hotel.Name = hotelDto.Name
	hotel.Description = hotelDto.Description
	hotel.RoomAmount = hotelDto.RoomAmount
	hotel.StreetName = hotelDto.StreetName
	hotel.StreetNumber = hotelDto.StreetNumber
	hotel.Rate = hotelDto.Rate

	for _, amenityName := range hotelDto.Amenities {
		amenity := client.GetAmenityByName(amenityName)

		if amenity.Id == 0 {
			return hotelDto, errors.New("amenity not found")
		}

		hotel.Amenities = append(hotel.Amenities, amenity)

	}

	hotel = client.InsertHotel(hotel)

	hotelDto.Id = hotel.Id

	if hotel.Id == 0 {
		return hotelDto, errors.New("error creating hotel")
	}

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
		hotelDto.Rate = hotel.Rate

		for _, image := range hotel.Images {
			var imageDto dto.ImageDto
			imageDto.Id = image.Id
			imageDto.Path = image.Path
			imageDto.HotelId = image.HotelId

			hotelDto.Images = append(hotelDto.Images, imageDto)
		}

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
	hotelDto.Rate = hotel.Rate

	for _, amenity := range hotel.Amenities {
		hotelDto.Amenities = append(hotelDto.Amenities, amenity.Name)
	}

	for _, image := range hotel.Images {
		var imageDto dto.ImageDto
		imageDto.Id = image.Id
		imageDto.Path = image.Path
		imageDto.HotelId = image.HotelId

		hotelDto.Images = append(hotelDto.Images, imageDto)
	}

	return hotelDto, nil
}

func (s *service) InsertReservation(reservationDto dto.ReservationDto) (dto.ReservationDto, error) {

	userDto := client.GetUserById(reservationDto.UserId)
	hotelDto := client.GetHotelById(reservationDto.HotelId)

	if userDto.Id == 0 {
		return reservationDto, errors.New("user not found")
	}

	if hotelDto.Id == 0 {
		return reservationDto, errors.New("hotel not found")
	}

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

		hoursAmount := timeEnd.Sub(timeStart).Hours()
		nightsAmount := math.Ceil(hoursAmount / 24)
		rate := hotelDto.Rate

		reservation.Amount = rate * nightsAmount

		reservation = client.InsertReservation(reservation)

		reservationDto.Id = reservation.Id
		reservationDto.Amount = reservation.Amount

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
	reservationDto.Amount = reservation.Amount

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
		reservationDto.Amount = reservation.Amount

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
		reservationDto.Amount = reservation.Amount

		reservationsDto = append(reservationsDto, reservationDto)
	}

	userReservationsDto.Reservations = reservationsDto

	return userReservationsDto, nil
}

func (s *service) GetReservationsByUserRange(userId int, startDate string, endDate string) (dto.ReservationsDto, error) {

	var reservationsInRange dto.ReservationsDto

	rangeStart, _ := time.Parse("02-01-2006 15:04", startDate)
	rangeEnd, _ := time.Parse("02-01-2006 15:04", endDate)

	if rangeStart.After(rangeEnd) {
		return reservationsInRange, errors.New("a reservation cant end before it starts")
	}

	reservations := client.GetReservationsByUser(userId)

	for _, reservation := range reservations {

		reservationStart, _ := time.Parse("02-01-2006 15:04", reservation.StartDate)
		reservationEnd, _ := time.Parse("02-01-2006 15:04", reservation.EndDate)

		if reservationStart.After(rangeStart) && reservationEnd.Before(rangeEnd) ||
			reservationStart.Before(rangeStart) && reservationEnd.After(rangeEnd) ||
			reservationStart.Before(rangeEnd) && reservationEnd.After(rangeEnd) ||
			reservationStart.Before(rangeStart) && reservationEnd.After(rangeEnd) ||
			reservationStart.Equal(rangeStart) || reservationEnd.Equal(rangeEnd) {

			var reservationDto dto.ReservationDto

			reservationDto.Id = reservation.Id
			reservationDto.StartDate = reservation.StartDate
			reservationDto.EndDate = reservation.EndDate
			reservationDto.Amount = reservation.Amount
			reservationDto.UserId = reservation.UserId
			reservationDto.HotelId = reservation.HotelId

			reservationsInRange = append(reservationsInRange, reservationDto)
		}
	}

	return reservationsInRange, nil
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
	hotelReservations.HotelRate = hotel.Rate

	for _, reservation := range reservations {
		var reservationDto dto.ReservationDto
		reservationDto.Id = reservation.Id
		reservationDto.StartDate = reservation.StartDate
		reservationDto.EndDate = reservation.EndDate
		reservationDto.HotelId = reservation.HotelId
		reservationDto.UserId = reservation.UserId
		reservationDto.Amount = reservation.Amount

		reservationsDto = append(reservationsDto, reservationDto)
	}

	hotelReservations.Reservations = reservationsDto

	return hotelReservations, nil
}

func (s *service) InsertAmenity(amenityDto dto.AmenityDto) (dto.AmenityDto, error) {
	var amenity model.Amenity

	amenity.Name = amenityDto.Name

	amenity = client.InsertAmenity(amenity)

	if amenity.Id == 0 {
		return amenityDto, errors.New("error creating amenity")
	}

	amenityDto.Id = amenity.Id

	return amenityDto, nil
}

func (s *service) GetAmenities() (dto.AmenitiesDto, error) {
	var amenities model.Amenities = client.GetAmenities()
	var amenitiesDto dto.AmenitiesDto

	for _, amenity := range amenities {
		var amenityDto dto.AmenityDto
		amenityDto.Id = amenity.Id
		amenityDto.Name = amenity.Name

		amenitiesDto = append(amenitiesDto, amenityDto)
	}

	return amenitiesDto, nil
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

func (s *service) CheckAllAvailability(startDate string, endDate string) (dto.HotelsDto, error) {

	var hotelsAvailable dto.HotelsDto

	reservationStart, _ := time.Parse("02-01-2006 15:04", startDate)
	reservationEnd, _ := time.Parse("02-01-2006 15:04", endDate)

	if reservationStart.After(reservationEnd) {
		return hotelsAvailable, errors.New("a reservation cant end before it starts")
	}

	hotels := client.GetHotels()

	for _, hotel := range hotels {
		if s.CheckAvailability(hotel.Id, reservationStart, reservationEnd) {
			var hotelDto dto.HotelDto
			hotelDto.Id = hotel.Id
			hotelDto.Name = hotel.Name
			hotelDto.StreetName = hotel.StreetName
			hotelDto.StreetNumber = hotel.StreetNumber
			hotelDto.RoomAmount = hotel.RoomAmount
			hotelDto.Rate = hotel.Rate
			hotelDto.Description = hotel.Description

			hotelsAvailable = append(hotelsAvailable, hotelDto)
		}
	}

	return hotelsAvailable, nil
}

func (s *service) InsertImages(imagesDto dto.ImagesDto) (dto.ImagesDto, error) {

	var images model.Images

	for _, imageDto := range imagesDto {
		var image model.Image

		image.Path = imageDto.Path
		image.HotelId = imageDto.HotelId

		images = append(images, image)
	}

	images = client.InsertImages(images)

	if len(images) != len(imagesDto) {
		return imagesDto, errors.New("failed to insert images")
	}

	for i, image := range images {
		if image.Id == 0 {
			return imagesDto, errors.New("failed to insert images")
		}

		imagesDto[i].Id = image.Id
	}

	return imagesDto, nil
}
