package Location

import (
	DB "burgher/Storage/PSQL"

	"gorm.io/gorm"
)

type Location struct {
	gorm.Model
	Id        string
	UserId    string `gorm:"unique;not null"`
	Latitude  float64
	Longitude float64
	City      string
	Timestamp int64
}

func init() {
	DB.Connect().AutoMigrate(&Location{})
}

type LocationRequest struct {
	UserId    string  `json:"userId"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	City      string  `json:"city"`
}

type LocationResponse struct {
	Id        string  `json:"id"`
	UserId    string  `json:"userId"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	City      string  `json:"city"`
}
