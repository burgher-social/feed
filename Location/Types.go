package Location

import (
	DB "burgher/Storage/PSQL"

	"github.com/twpayne/go-geom"
	"gorm.io/gorm"
)

type Location struct {
	gorm.Model
	Id        string
	UserId    string `gorm:"unique;not null"`
	Latitude  float64
	Longitude float64
	City      string
}

func init() {
	DB.Connect().AutoMigrate(&Location{})
	DB.Connect().AutoMigrate(&PostsLocation{})
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

type PostsLocation struct {
	gorm.Model
	Id        string      `gorm:"type:string;primaryKey" json:"id"`
	Timestamp int64       `gorm:"type:bigint" json:"timestamp"`
	Score     int         `gorm:"type:integer" json:"score"`
	Location  *geom.Point `gorm:"type:GEOMETRY(Point,4326)" json:"location"`
}
