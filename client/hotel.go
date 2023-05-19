package client

import (
	log "github.com/sirupsen/logrus"
	"project/model"
)

func InsertHotel(hotel model.Hotel) model.Hotel {

	result := Db.Create(&hotel)

	if result.Error != nil {
		log.Error("Failed to insert hotel.")
	}

	log.Debug("User created:", hotel.Id)
	return hotel
}
