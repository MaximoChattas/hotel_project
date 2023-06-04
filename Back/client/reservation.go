package client

import (
	log "github.com/sirupsen/logrus"
	"project/model"
)

// Inserts a reservation into db
func InsertReservation(reservation model.Reservation) model.Reservation {

	result := Db.Create(&reservation)

	if result.Error != nil {
		log.Error("Failed to insert reservation.")
	}

	log.Debug("User created:", reservation.Id)
	return reservation
}

// Returns a single reservation by its id
func GetReservationById(id int) model.Reservation {
	var reservation model.Reservation

	Db.Where("id = ?", id).First(&reservation)
	log.Debug("Reservation: ", reservation)

	return reservation
}

// Returns all reservations
func GetReservations() model.Reservations {
	var reservations model.Reservations
	Db.Find(&reservations)

	log.Debug("Reservations: ", reservations)

	return reservations
}

// Return all reservations from a specific user
func GetReservationsByUser(userId int) model.Reservations {
	var reservations model.Reservations

	Db.Where("user_id = ?", userId).Find(&reservations)
	log.Debug("Reservations: ", reservations)

	return reservations
}

// Return all reservations in a specific hotel
func GetReservationsByHotel(hotelId int) model.Reservations {
	var reservations model.Reservations

	Db.Where("hotel_id = ?", hotelId).Find(&reservations)
	log.Debug("Reservations: ", reservations)

	return reservations
}
