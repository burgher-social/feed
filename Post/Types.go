package Post

import (
	"burgher/Insights"
	DB "burgher/Storage/PSQL"
	Topic "burgher/Topic"
	"burgher/User"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Id       string
	Content  string
	UserId   string
	ParentId *string
}

type PostTopics struct {
	gorm.Model
	PostId  string
	TopicId string
}

func init() {
	DB.Connect().AutoMigrate(&Post{})
	DB.Connect().AutoMigrate(&PostTopics{})
	DB.Connect().AutoMigrate(&PostsLocation{})

}

type PostRequest struct {
	Content  string   `json:"content"`
	UserId   string   `json:"userId"`
	ParentId *string  `json:"parentId"`
	Topics   []string `json:"topics"`
}

type PostResponse struct {
	Id       string  `json:"id"`
	Content  string  `json:"content"`
	ParentId *string `json:"parentId"`
	UserId   string  `json:"userId"`
}

type PostInfo struct {
	PostResponse      `gorm:"embedded" json:"post"`
	PostTopics        `gorm:"embedded;embeddedPrefix:post_topics_" json:"postTopic"`
	Topic.Topic       `gorm:"embedded;embeddedPrefix:topic_" json:"topic"`
	User.User         `gorm:"embedded;embeddedPrefix:user_" json:"user"`
	LocationMaker     `gorm:"embedded;embeddedPrefix:posts_locations_" json:"location"`
	Insights.Insights `gorm:"embedded;embeddedPrefix:insights_" json:"insights"`
}

type LocationMaker struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type PostsLocation struct {
	gorm.Model
	Id        string  `gorm:"type:string;primaryKey" json:"id"`
	Timestamp int64   `gorm:"type:bigint" json:"timestamp"`
	Score     int     `gorm:"type:integer" json:"score"`
	Location  *string `gorm:"type:GEOMETRY(Point,4326)" json:"location"`
}
