package dto

type HotelReservationsDto struct {
	HotelId           int     `json:"hotel_id"`
	HotelName         string  `json:"hotel_name"`
	HotelRoomAmount   int     `json:"hotel_room_amount"`
	HotelDescription  string  `json:"hotel_description"`
	HotelStreetName   string  `json:"hotel_street_name"`
	HotelStreetNumber int     `json:"hotel_street_number"`
	HotelRate         float32 `json:"hotel_rate"`

	Reservations ReservationsDto `json:"reservations"`
}
