package client

import (
	log "github.com/sirupsen/logrus"
	"project/model"
)


func InsertReservation(reservation model.Reservation) model.Reservation {

	result := Db.Create(&reservation)

	if result.Error != nil {
		log.Error("Failed to insert reservation.")
		return reservation
	}

	log.Debug("Reservation created:", reservation.Id)
	return reservation
}

func GetReservationById(id int) model.Reservation {
	var reservation model.Reservation

	Db.Where("id = ?", id).First(&reservation)
	log.Debug("Reservation: ", reservation)

	return reservation
}

func GetReservations() model.Reservations {
	var reservations model.Reservations
	Db.Find(&reservations)

	log.Debug("Reservations: ", reservations)

	return reservations
}

func GetReservationsByUser(userId int) model.Reservations {
	var reservations model.Reservations

	Db.Where("user_id = ?", userId).Find(&reservations)
	log.Debug("Reservations: ", reservations)

	return reservations
}

func GetReservationsByHotel(hotelId int) model.Reservations {
	var reservations model.Reservations

	Db.Where("hotel_id = ?", hotelId).Find(&reservations)
	log.Debug("Reservations: ", reservations)

	return reservations
}