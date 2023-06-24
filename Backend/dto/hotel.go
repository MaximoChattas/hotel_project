package dto

type HotelDto struct {
	Id           int      `json:"id"`
	Name         string   `json:"name" validate:"required"`
	RoomAmount   int      `json:"room_amount" validate:"required"`
	Description  string   `json:"description" validate:"required"`
	StreetName   string   `json:"street_name" validate:"required"`
	StreetNumber int      `json:"street_number" validate:"required"`
	Rate         float64  `json:"rate" validate:"required"`
	Amenities    []string `json:"amenities"`
}

type HotelsDto []HotelDto
