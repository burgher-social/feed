package Post

import (
	"fmt"

	DB "burgher/Storage/PSQL"
)

func create(post Post, topics []string) (Post, error) {
	fmt.Println(post)
	var posttoppics = []PostTopics{}
	for _, x := range topics {
		posttoppics = append(posttoppics, PostTopics{PostId: post.Id, TopicId: x})

	}
	DB.Connect().Create(&post)
	DB.Connect().Create(&posttoppics)
	return post, nil
}

func Read(userId string) ([]Post, error) {
	var posts []Post
	DB.Connect().Where("user_id = ?", userId).Find(&posts)
	fields := `posts.*, post_topics.topic_id as topic_id,
		topics.name as topic_name`
	res := []PostInfo{}
	DB.Connect().Model(&posts).Select(fields).Where("user_id = ?", userId).Joins("INNER JOIN post_topics on posts.id = post_topics.post_id").
		Joins("INNER JOIN topics on post_topics.topic_id = topics.id").
		Scan(&res)

	fmt.Println(res)
	fmt.Printf("%+v\n", res)
	return posts, nil
}

func ReadOne(postId string) (PostInfo, error) {
	var post []Post
	fields := `posts.*, post_topics.topic_id as topic_id,
		topics.name as topic_name`
	res := []PostInfo{}
	DB.Connect().Model(&post).Select(fields).Where("id = ?", postId).Joins("INNER JOIN post_topics on posts.id = post_topics.post_id").
		Joins("INNER JOIN topics on post_topics.topic_id = topics.id").
		Scan(&res)
	return res[0], nil
}
