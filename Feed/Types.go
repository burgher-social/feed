package Feed

import (
	DB "burgher/Storage/PSQL"

	"gorm.io/gorm"
)

type UserFeed struct {
	gorm.Model
	Id        string `gorm:"type:string;primaryKey" json:"id"`
	UserId    string `json:"userId"`
	PostId    string `json:"postId"`
	Score     int    `gorm:"type:integer" json:"score"`
	Timestamp int64  `json:"timestamp"`
}

type UserFeedResponse struct {
	PostId    string `json:"postId"`
	Score     int    `json:"score"`
	Timestamp int64  `json:"timestamp"`
}

func init() {
	DB.Connect().AutoMigrate(&UserFeed{})
}
