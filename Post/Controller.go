package Post

import (
	"fmt"

	Insights "burgher/Insights"
	DB "burgher/Storage/PSQL"
)

// var fields = `posts.*, users.user_name as user_user_name, users.image_url as user_image_url, post_topics.topic_id as topic_id,
// topics.name as topic_name, ST_Y(posts_locations.location::geometry) as posts_locations_latitude, ST_X(posts_locations.location::geometry) as posts_locations_longitude`

var fields = `posts.*, users.user_name as user_user_name, users.image_url as user_image_url, post_topics.topic_id as topic_id,
		topics.name as topic_name, locations.latitude as posts_locations_latitude, locations.longitude as posts_locations_longitude,
		insights.likes as insights_likes, insights.comments as insights_comments`

func create(post Post, topics []string) (Post, error) {
	fmt.Println(post)
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

func Read(userId string) ([]PostInfo, error) {
	var posts []Post
	DB.Connect().Where("parent_id is NULL and user_id = ?", userId).Find(&posts)
	// fields := `posts.*, post_topics.topic_id as topic_id,
	// 	topics.name as topic_name`
	res := []PostInfo{}
	// DB.Connect().Model(&posts).Select(fields).Where("user_id = ?", userId).Joins("INNER JOIN post_topics on posts.id = post_topics.post_id").
	// 	Joins("INNER JOIN topics on post_topics.topic_id = topics.id").
	// 	Scan(&res)
	DB.Connect().Model(&posts).Select(fields).Where("parent_id is NULL and posts.user_id = ?", userId).
		Joins("LEFT JOIN users on posts.user_id = users.id").
		Joins("LEFT JOIN post_topics on posts.id = post_topics.post_id").
		Joins("LEFT JOIN topics on topics.id = post_topics.topic_id").
		Joins("LEFT JOIN locations on users.id = locations.user_id").
		Joins("LEFT JOIN insights on posts.id = insights.post_id").
		// Joins("LEFT JOIN posts_locations on CONCAT(users.id, ':', posts.id) = posts_locations.id").
		Scan(&res)

	// fmt.Println(res)
	// fmt.Printf("%+v\n", res)
	return res, nil
}

func ReadOne(postId string) (PostInfo, error) {
	var post []Post
	// fields := `posts.*, users.user_name as user_user_name, users.image_url as user_image_url, post_topics.topic_id as topic_id,
	// 	topics.name as topic_name`
	res := []PostInfo{}
	DB.Connect().Model(&post).Select(fields).Where("posts.id = ?", postId).
		Joins("LEFT JOIN users on posts.user_id = users.id").
		Joins("LEFT JOIN post_topics on posts.id = post_topics.post_id").
		Joins("LEFT JOIN topics on topics.id = post_topics.topic_id").
		Joins("LEFT JOIN locations on users.id = locations.user_id").
		Joins("LEFT JOIN insights on posts.id = insights.post_id").
		// Joins("LEFT JOIN posts_locations on CONCAT(users.id, ':', posts.id) = posts_locations.id").
		Scan(&res)
	fmt.Println(res)
	// fmt.Printf("%+v\n", res)
	if len(res) == 0 {
		return PostInfo{}, fmt.Errorf("no post found")
	}
	return res[0], nil
}

func readComments(postId string) ([]PostInfo, error) {
	var post []Post
	// fields := `posts.*, post_topics.topic_id as topic_id,
	// 	topics.name as topic_name`
	res := []PostInfo{}
	DB.Connect().Model(&post).Select(fields).Where("posts.parent_id = ?", postId).
		Joins("LEFT JOIN users on posts.user_id = users.id").
		Joins("LEFT JOIN post_topics on posts.id = post_topics.post_id").
		Joins("LEFT JOIN topics on topics.id = post_topics.topic_id").
		Joins("LEFT JOIN locations on users.id = locations.user_id").
		Joins("LEFT JOIN insights on posts.id = insights.post_id").
		// Joins("LEFT JOIN posts_locations on CONCAT(users.id, ':', posts.id) = posts_locations.id").
		Scan(&res)
	fmt.Println(res)
	// if len(res) == 0 {
	// 	return []PostInfo{}, fmt.Errorf("no post found")
	// }
	return res, nil
}
