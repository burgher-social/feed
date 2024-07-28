package Post

import (
	"fmt"

	Insights "burgher/Insights"
	DB "burgher/Storage/PSQL"

	"gorm.io/gorm"
)

var fields = `posts.*,
		users.user_name as user_user_name, users.image_url as user_image_url,
		post_topics.topic_id as topic_id,
		topics.name as topic_name,
		locations.latitude as posts_locations_latitude, locations.longitude as posts_locations_longitude,
		insights.likes as insights_likes, insights.comments as insights_comments,
		likes_posts.post_id as likes_posts_post_id
		`

func create(post Post, topics []string) (Post, error) {
	if len(post.Content) > 300 {
		return post, fmt.Errorf("content length too long")
	}
	var posttoppics = []PostTopics{}
	for _, x := range topics {
		posttoppics = append(posttoppics, PostTopics{PostId: post.Id, TopicId: x})

	}
	DB.Connect().Create(&post)
	DB.Connect().Create(&posttoppics)
	Insights.InitInsights(post.Id)
	if post.ParentId != nil {
		Insights.CommentCount(1, *post.ParentId)
	}
	return post, nil
}

func joinTables(tx *gorm.DB, userId string) *gorm.DB {
	return tx.
		Joins("LEFT JOIN users on posts.user_id = users.id").
		Joins("LEFT JOIN post_topics on posts.id = post_topics.post_id").
		Joins("LEFT JOIN topics on topics.id = post_topics.topic_id").
		Joins("LEFT JOIN locations on users.id = locations.user_id").
		Joins("LEFT JOIN insights on posts.id = insights.post_id").
		Joins("LEFT JOIN likes_posts on posts.id = likes_posts.post_id AND likes_posts.user_id = '" + userId + "'")
}

func Read(userId string, authUserId string) ([]PostInfo, error) {
	var posts []Post
	res := []PostInfo{}
	joinTables(DB.Connect().Model(&posts).Select(fields).Where("parent_id is NULL and posts.user_id = ?", userId), authUserId).
		Scan(&res)

	return res, nil
}

func ReadOne(postId string, authUserId string) (PostInfo, error) {
	var post []Post
	res := []PostInfo{}
	joinTables(DB.Connect().Model(&post).Select(fields).Where("posts.id = ?", postId), authUserId).
		Scan(&res)

	if len(res) == 0 {
		return PostInfo{}, fmt.Errorf("no post found")
	}

	return res[0], nil
}

func readComments(postId string, authUserId string) ([]PostInfo, error) {
	var post []Post
	res := []PostInfo{}
	joinTables(DB.Connect().Model(&post).Select(fields).Where("posts.parent_id = ?", postId), authUserId).
		Scan(&res)
	return res, nil
}
