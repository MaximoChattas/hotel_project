package model

type Hotel struct {
	Id          int    `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar(300); not null"`
	RoomAmount  int    `gorm:"type:int; not null"`
	Description string `gorm:"type:varchar(30000)"`
}

type Hotels []Hotel
