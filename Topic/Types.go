package Topic

import (
	DB "burgher/Storage/PSQL"

	"gorm.io/gorm"
)

type Topic struct {
	gorm.Model
	Id   string
	Name string
}

func init() {
	DB.Connect().AutoMigrate(&Topic{})
}

type TopicRequest struct {
	Name string `json:"name"`
}

type TopicResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
