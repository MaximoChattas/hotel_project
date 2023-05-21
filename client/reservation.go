package client

import (
	log "github.com/sirupsen/logrus"
	"project/model"
)

func InsertReservation(reservation model.Reservation) model.Reservation {

	result := Db.Create(&reservation)

	if result.Error != nil {
		log.Error("Failed to insert reservation.")
	}

	log.Debug("User created:", reservation.Id)
	return reservation
}

func GetReservationById(id int) model.Reservation {}

func GetReservations() model.Reservations {}

func GetUsersReservations(userId int) model.Reservations {}
