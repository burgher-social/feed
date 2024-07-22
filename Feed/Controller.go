package Feed

import (
	Location "burgher/Location"
	"burgher/Post"
	DB "burgher/Storage/PSQL"
	Utils "burgher/Utils"
	"fmt"
	"strings"
	"time"
)

var minutes_45_in_miliseconds int64 = 0 // 45 * 60 * 1000
var radius int = 10000

func create(userId string, loc Location.Location) {
	// ctx := context.Background()
	// radiusQuery := redis.GeoRadiusQuery{
	// 	Radius: 5, // Radius in kilometers
	// 	Unit:   "km",
	// }
	// if loc, err := Location.Read(userId); err == nil {

	// todo: check if feed generation is already in progress

	var postlocations []Post.PostsLocation
	if err := DB.Connect().Select("*", "ST_AsText(location) as location").Where("ST_DWithin(location, ST_SetSRID(ST_MakePoint(?, ?), 4326)::geography, ?)", loc.Longitude, loc.Latitude, radius).Order("score DESC").Limit(20000).Find(&postlocations).Error; err == nil {
		var userFeeds []UserFeed
		curTimestamp := time.Now().UnixNano() / 1e6
		for i := 0; i < len(postlocations); i++ {
			postId := strings.Split(postlocations[i].Id, ":")[1]
			userFeeds = append(userFeeds, UserFeed{
				Id:        Utils.GenerateId(),
				UserId:    userId,
				PostId:    postId,
				Score:     postlocations[i].Score,
				Timestamp: curTimestamp,
			})
		}
		DB.Connect().Where("user_id = ?", userId).Delete(&UserFeed{})
		DB.Connect().Create(&userFeeds)
	}
	// Redis.GetInstance().GeoRadius(ctx, "users_posts_id", loc.Longitude, loc.Latitude, &radiusQuery).Result()
}

// }

func read(userId string, offset int, limit int) ([]UserFeed, error) {
	var userFeeds []UserFeed

	DB.Connect().Where("user_id = ?", userId).Order("score DESC").Offset(offset).Limit(limit).Find(&userFeeds)
	if len(userFeeds) == 0 {
		if loc, err := Location.Read(userId); err == nil {
			create(userId, loc)
			DB.Connect().Where("user_id = ?", userId).Order("score DESC").Offset(offset).Limit(limit).Find(&userFeeds)
			if len(userFeeds) == 0 {
				return []UserFeed{}, nil
			}
		}
	}
	curTimestamp := time.Now().UnixNano() / 1e6
	lastTimestamp := userFeeds[0].Timestamp

	if curTimestamp-lastTimestamp > minutes_45_in_miliseconds {
		if loc, err := Location.Read(userId); err == nil {
			go create(userId, loc)
		}
	}
	fmt.Println("USER POSTS")
	fmt.Println(userFeeds)
	return userFeeds, nil
}
