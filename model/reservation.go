package model

import "time"

type Reservation struct {
	Id        int       `gorm:"primaryKey"`
	StartDate time.Time `gorm:"column:StartDate"`
	EndDate   time.Time `gorm:"column:EndDate"`
	UserId    int       `gorm:"foreignkey:UserId"`
	HotelId   int       `gorm:"foreignkey:HotelId"`
}

type Reservations []Reservation
