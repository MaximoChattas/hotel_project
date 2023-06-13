package dto

type HotelDto struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	RoomAmount   int     `json:"room_amount"`
	Description  string  `json:"description"`
	StreetName   string  `json:"street_name"`
	StreetNumber int     `json:"street_number"`
	Rate         float64 `json:"rate"`
}

type HotelsDto []HotelDto
