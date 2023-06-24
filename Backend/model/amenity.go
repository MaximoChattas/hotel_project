package model

type Amenity struct {
	Id   int    `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(300); not null; unique"`
}

type Amenities []Amenity
