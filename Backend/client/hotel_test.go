package client

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"project/model"
	"testing"
)

func init() {
	// DB Connections Paramters
	DBName := ""
	DBUser := "root"
	DBPass := ""
	DBHost := "localhost"

	db, err := gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3306)/"+DBName+"?charset=utf8&parseTime=True")

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	// Add all clients here
	Db = db

	Db.AutoMigrate(&model.Hotel{})

}

func TestGetHotelById(t *testing.T) {
	a := assert.New(t)

	hotel := HotelClient.GetHotelById(1)

	a.Equal(1, hotel.Id)
}

func TestGetHotels(t *testing.T) {
	a := assert.New(t)

	hotels := HotelClient.GetHotels()

	var amount int

	Db.Model(&model.Hotel{}).Count(&amount)

	a.Equal(amount, len(hotels))
}
