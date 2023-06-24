package dto

type AmenityDto struct {
	Id   int    `json:"id"`
	Name string `json:"name" validate:"required"`
}

type AmenitiesDto []AmenityDto
