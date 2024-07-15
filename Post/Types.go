package Post

import (
	DB "burgher/Storage/PSQL"
	Topic "burgher/Topic"

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
	PostResponse `gorm:"embedded" json:"post"`
	PostTopics   `gorm:"embedded;embeddedPrefix:post_topics_" json:"postTopic"`
	Topic.Topic  `gorm:"embedded;embeddedPrefix:topic_" json:"topic"`
}
