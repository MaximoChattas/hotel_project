package db

import (
	"project/client"
	"project/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	Db  *gorm.DB
	err error
)

func init() {
	// DB Connections Paramters
	DBName := "miranda"
	DBUser := "root"
	DBPass := "pass"
	DBHost := "database"

	Db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3307)/"+DBName+"?charset=utf8&parseTime=True")

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	// Add all clients here
	client.Db = Db

}

func StartDbEngine() {
	// Migrate all model classes
	Db.AutoMigrate(&model.Hotel{})
	Db.AutoMigrate(&model.Reservation{})
	Db.AutoMigrate(&model.User{})
	Db.AutoMigrate(&model.Amenity{})
	Db.AutoMigrate(&model.Image{})

	log.Info("Finishing Migration Database Tables")
}
