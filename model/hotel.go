package model

type Hotel struct {
	Id           int     `gorm:"primaryKey"`
	Name         string  `gorm:"type:varchar(300); not null"`
	RoomAmount   int     `gorm:"type:int; not null"`
	Description  string  `gorm:"type:varchar(1000)"`
	StreetName   string  `gorm:"type:varchar(100)"`
	StreetNumber int     `gorm:"type:int"`
	Rate         float64 `gorm:"type:decimal(8,2); not null"`
}

type Hotels []Hotel
