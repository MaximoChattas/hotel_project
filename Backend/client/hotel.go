package client

import (
	log "github.com/sirupsen/logrus"
	"project/model"
)

func InsertHotel(hotel model.Hotel) model.Hotel {

	result := Db.Create(&hotel)

	if result.Error != nil {
		log.Error("Failed to insert hotel.")
		return hotel
	}

	log.Debug("Hotel created:", hotel.Id)
	return hotel
}

func GetHotelById(id int) model.Hotel {
	var hotel model.Hotel

	Db.Where("id = ?", id).First(&hotel)
	log.Debug("Hotel: ", hotel)

	return hotel
}

func GetHotels() model.Hotels {
	var hotels model.Hotels
	Db.Find(&hotels)

	log.Debug("Hotels: ", hotels)

	return hotels
}
