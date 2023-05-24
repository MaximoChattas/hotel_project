package model

type Reservation struct {
	Id        int    `gorm:"primaryKey"`
	StartDate string `gorm:"type:varchar(16); not null"` //Expected time as "DD-MM-YYYY hh:mm"
	EndDate   string `gorm:"type:varchar(16); not null"`
	UserId    int    `gorm:"foreignkey:UserId"`
	HotelId   int    `gorm:"foreignkey:HotelId"`
}

type Reservations []Reservation
