package Location

import (
	"fmt"
	"time"

	"burgher/Post"
	DB "burgher/Storage/PSQL"

	"github.com/twpayne/go-geom"
	"gorm.io/gorm/clause"
)

func create(location Location) (Location, error) {
	// ctx := context.Background()
	fmt.Println(location)

	if loc, err := Read(location.UserId); err == nil {
		if posts, perr := Post.Read(loc.UserId); perr == nil {
			for i := 0; i < len(posts); i++ {
				DB.Connect().Delete(&PostsLocation{}, location.UserId+":"+posts[i].Id)
			}

			for i := 0; i < len(posts); i++ {
				postLocation := PostsLocation{
					Id:        location.UserId + ":" + posts[i].Id,
					Timestamp: time.Now().UnixNano() / 1e6,
					Location:  *geom.NewPoint(geom.XY).MustSetCoords(geom.Coord{location.Longitude, location.Latitude}),
					Score:     0,
				}
				DB.Connect().Create(&postLocation)
			}

			// REDIS
			// // Find location changed much or not
			// if posts, perr := Post.Read(loc.UserId); perr == nil {
			// 	for i := 0; i < len(posts); i++ {
			// 		Redis.GetInstance().ZRem(ctx, "users_posts_id", location.UserId+":"+posts[i].Id).Err()
			// 	}

			// 	for i := 0; i < len(posts); i++ {
			// 		Redis.GetInstance().GeoAdd(ctx, "users_posts_id", &redis.GeoLocation{
			// 			Longitude: loc.Longitude,
			// 			Latitude:  loc.Latitude,
			// 			Name:      location.UserId + ":" + posts[i].Id,
			// 		}).Result()
			// 	}
			// }
		}
	}

	// DB.Connect().(&location)
	DB.Connect().Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "user_id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"latitude":  location.Latitude,
			"longitude": location.Longitude,
			"city":      location.City,
		}),
	}).Create(&location)
	return location, nil
}

func Read(userId string) (Location, error) {
	var locations Location
	DB.Connect().First(&locations, "user_id = ?", userId)
	return locations, nil
}
