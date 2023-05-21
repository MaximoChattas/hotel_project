package client

import (
	"project/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func InsertUser(user model.User) model.User {

	result := Db.Create(&user)

	if result.Error != nil {
		log.Error("Failed to insert user.")
	}

	log.Debug("User created:", user.Id)
	return user
}

func GetUserById(id int) model.User {}

func GetUsers() model.Users {}
