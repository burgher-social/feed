package Insights

import (
	DB "burgher/Storage/PSQL"

	"gorm.io/gorm"
)

type Insights struct {
	gorm.Model
	Id       string `json:"id"`
	PostId   string `gorm:"unique;not null" json:"postId"`
	Likes    uint64 `json:"likes"`
	Comments uint64 `json:"comments"`
}

func init() {
	DB.Connect().AutoMigrate(&Insights{})
}
