package Location

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	Post "burgher/Post"
	DB "burgher/Storage/PSQL"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func create(location Location) (Location, error) {
	// ctx := context.Background()
	fmt.Println(location)
	// TODO: don't update location if within certain radius
	if loc, err := Read(location.UserId); err == nil {
		if posts, perr := Post.Read(loc.UserId); perr == nil {
			for i := 0; i < len(posts); i++ {
				DB.Connect().Unscoped().Delete(&Post.PostsLocation{}, "id = ?", location.UserId+":"+posts[i].PostResponse.Id)
			}

			// pt := geom.NewPoint(geom.XY).MustSetCoords(geom.Coord{location.Longitude, location.Latitude}).SetSRID(4326)
			// ewkbData, _ := ewkb.Marshal(pt)
			// fmt.Println(pt)
			a := "0101000020E6100000D02C6409C6675340DB02F85DC7E22940"
			for i := 0; i < len(posts); i++ {
				postLocation := Post.PostsLocation{
					Id:        location.UserId + ":" + posts[i].PostResponse.Id,
					Timestamp: time.Now().UnixNano() / 1e6,
					Location:  &a,
					Score:     0,
				}
				sqlStr := fmt.Sprint(`INSERT INTO "posts_locations" 
					("id","created_at","updated_at","deleted_at","timestamp","score","location")
					VALUES 
					('`, postLocation.Id, `', 
					now(), 
					now(), 
					null,
					'`, strconv.FormatInt(postLocation.Timestamp, 10), `',
					'`, postLocation.Score, `',
					ST_GeomFromText('POINT(`, strconv.FormatFloat(location.Longitude, 'f', -1, 64), ` `, strconv.FormatFloat(location.Latitude, 'f', -1, 64), `)', 4326)) ON CONFLICT (id) DO NOTHING
					RETURNING "id";`)
				// print(sqlStr)
				tx := DB.Connect().Exec(sqlStr)
				if tx.Error != nil {
					fmt.Println(tx.Error)
				}

				// DB.Connect().Create(&postLocation)
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
	dbRresult := DB.Connect().First(&locations, "user_id = ?", userId)
	if errors.Is(dbRresult.Error, gorm.ErrRecordNotFound) {
		// fmt.Println("handling not found error")
		return locations, fmt.Errorf("location doesn't exist")
	}
	return locations, nil
}
