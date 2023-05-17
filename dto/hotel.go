package dto

type HotelDto struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	RoomAmount  int    `json:"room_amount"`
	Description string `json:"description"`
}

type HotelsDto []HotelDto
