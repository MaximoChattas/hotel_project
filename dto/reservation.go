package dto

import "time"

type ReservationDto struct {
	Id        int       `json:"id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	UserId    int       `json:"user_id"`
	HotelId   int       `json:"hotel_id"`
}

type ReservationsDto []ReservationDto
