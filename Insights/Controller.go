package Insights

import (
	DB "burgher/Storage/PSQL"
	"burgher/User"
	Utils "burgher/Utils"
	"fmt"

	"gorm.io/gorm/clause"
)

func Like(count int, postId string, authUserId string) error {
	userLike := User.LikesPosts{
		UserId: authUserId,
		PostId: postId,
	}
	DB.Connect().Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "post_id"}},
		DoNothing: true,
	}).Create(&userLike)
	sqlStr := fmt.Sprintf("UPDATE insights SET likes = likes + %d where post_id = '%s'", count, postId)
	DB.Connect().Exec(sqlStr)
	return nil
}

func UnLike(count int, postId string, authUserId string) error {
	var userLike User.LikesPosts = User.LikesPosts{
		UserId: authUserId,
		PostId: postId,
	}
	DB.Connect().Unscoped().Where("user_id = ? AND post_id = ?", userLike.UserId, userLike.PostId).Delete(&userLike)
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
