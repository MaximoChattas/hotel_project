package service

import (
	"github.com/stretchr/testify/assert"
	"project/client"
	"project/dto"
	"project/model"
	"testing"
)

type TestHotel struct{}

func init() {
	client.HotelClient = &TestHotel{}
}

func (t TestHotel) InsertHotel(hotel model.Hotel) model.Hotel {
	hotel.Id = 0

	return hotel
}

func (t TestHotel) GetHotelById(id int) model.Hotel {
	var hotel model.Hotel

	if id > 10 {
		hotel.Id = 0
	} else {
		hotel.Id = id
	}

	return hotel
}

func (t TestHotel) GetHotels() model.Hotels {
	return model.Hotels{}
}

func (t TestHotel) DeleteHotel(hotel model.Hotel) error {
	return nil
}

func TestInsertHotelError(t *testing.T) {

	a := assert.New(t)
	var hotelDto dto.HotelDto

	_, err := HotelService.InsertHotel(hotelDto)

	expectedResponse := "error creating hotel"

	a.NotNil(err)
	a.Equal(expectedResponse, err.Error())

}

func TestGetHotelByIdFound(t *testing.T) {

	a := assert.New(t)

	_, err := HotelService.GetHotelById(1)

	a.Nil(err)
}

func TestGetHotelByIdNotFound(t *testing.T) {

	a := assert.New(t)

	_, err := HotelService.GetHotelById(20)

	expectedResponse := "hotel not found"

	a.NotNil(err)
	a.Equal(expectedResponse, err.Error())
}
