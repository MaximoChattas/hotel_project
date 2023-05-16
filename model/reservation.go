package model

import "time"

type Reservation struct {
	Id        int `gorm:"primaryKey"`
	StartDate time.Time
	EndDate   time.Time
	UserId    int `gorm:"foreignkey:UserId"`
	HotelId   int `gorm:"foreignkey:HotelId"`
}
