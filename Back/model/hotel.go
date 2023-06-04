package model

type Hotel struct {
	Id           int     `gorm:"primaryKey"`
	Name         string  `gorm:"type:varchar(300); not null"`
	RoomAmount   int     `gorm:"type:int; not null"`
	Description  string  `gorm:"type:varchar(1000)"`
	StreetName   string  `gorm:"type:varchar(100)"`
	StreetNumber int     `gorm:"type:int"`
	Rate         float32 `gorm:"type:decimal(7,6)"`
}

type Hotels []Hotel
