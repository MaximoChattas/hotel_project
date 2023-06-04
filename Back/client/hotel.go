package client

import (
	log "github.com/sirupsen/logrus"
	"project/model"
)

// funcion que permite insertar nuevos hoteles
func InsertHotel(hotel model.Hotel) model.Hotel {

	result := Db.Create(&hotel)

	if result.Error != nil {
		log.Error("Failed to insert hotel.")
	}

	log.Debug("Hotel created:", hotel.Id)
	return hotel
}

// encuentra los hoteles por un ID pasado por parametro
func GetHotelById(id int) model.Hotel {
	var hotel model.Hotel

	Db.Where("id = ?", id).First(&hotel)
	log.Debug("Hotel: ", hotel)

	return hotel
}

// encuentra los hoteles
func GetHotels() model.Hotels {
	var hotels model.Hotels
	Db.Find(&hotels)

	log.Debug("Hotels: ", hotels)

	return hotels
}
