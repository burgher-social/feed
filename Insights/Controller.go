package Insights

import (
	DB "burgher/Storage/PSQL"
	Utils "burgher/Utils"
	"fmt"
)

func Like(count int, postId string) error {
	sqlStr := fmt.Sprintf("UPDATE insights SET likes = likes + %d where post_id = '%s'", count, postId)
	DB.Connect().Exec(sqlStr)
	return nil
}

func CommentCount(count int, postId string) error {
	sqlStr := fmt.Sprintf("UPDATE insights SET comments = comments + %d where post_id = '%s'", count, postId)
	DB.Connect().Exec(sqlStr)
	return nil
}

func InitInsights(postId string) {
	insights := Insights{
		Id:       Utils.GenerateId(),
		PostId:   postId,
		Likes:    0,
		Comments: 0,
	}
	DB.Connect().Create(&insights)
}
