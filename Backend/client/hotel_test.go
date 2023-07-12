package client

import (
	"project/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestInsertHotel(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to Create Mock Database")
	}
	defer db.Close()

	gormDB, err := gorm.Open("mysql", db)
	if err != nil {
		t.Fatalf("Connection Failed to Open")
	}

	Db = gormDB
	HotelClient = &hotelClient{}

	hotel := model.Hotel{
		Id:           1,
		Name:         "Sample Hotel",
		RoomAmount:   10,
		Description:  "Sample description",
		StreetName:   "Sample Street",
		StreetNumber: 123,
		Rate:         4.5,
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `hotels`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	result := HotelClient.InsertHotel(hotel)

	// Assert the result
	assert.Equal(t, hotel, result)
	assert.Equal(t, 1, hotel.Id)

	// Assert that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %v", err)
	}
}
