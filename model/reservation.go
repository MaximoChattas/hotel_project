package model

type Reservation struct {
	Id        int    `gorm:"primaryKey"`
	StartDate string `gorm:"type:varchar(19); not null"`
	EndDate   string `gorm:"type:varchar(19); not null"`
	UserId    int    `gorm:"foreignkey:UserId"`
	HotelId   int    `gorm:"foreignkey:HotelId"`
}

type Reservations []Reservation
