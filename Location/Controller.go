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

var locationUpdateThreshold int64 = 1 * 60

func create(location Location) (Location, error) {
	if loc, err := Read(location.UserId); err == nil {
		now := time.Now().UnixNano() / 1e6

		// Don't update location and posts with that location if time diff is very less
		if loc.Timestamp != 0 && loc.Timestamp-now < locationUpdateThreshold {
			return location, nil
		}
		if posts, perr := Post.Read(loc.UserId, location.UserId); perr == nil {
			for i := 0; i < len(posts); i++ {
				DB.Connect().Unscoped().Delete(&Post.PostsLocation{}, "id = ?", location.UserId+":"+posts[i].PostResponse.Id)
			}
			postLocationTimestamp := now
			postLocationScore := 0
			for i := 0; i < len(posts); i++ {
				postLocationId := location.UserId + ":" + posts[i].PostResponse.Id
				sqlStr := fmt.Sprint(`INSERT INTO "posts_locations" 
					("id","created_at","updated_at","deleted_at","timestamp","score","location")
					VALUES 
					('`, postLocationId, `', 
					now(), 
					now(), 
					null,
					'`, strconv.FormatInt(postLocationTimestamp, 10), `',
					'`, postLocationScore, `',
					ST_GeomFromText('POINT(`, strconv.FormatFloat(location.Longitude, 'f', -1, 64), ` `, strconv.FormatFloat(location.Latitude, 'f', -1, 64), `)', 4326)) ON CONFLICT (id) DO NOTHING
					RETURNING "id";`)
				tx := DB.Connect().Exec(sqlStr)
				if tx.Error != nil {
					fmt.Println(tx.Error)
				}
			}
		}
	}

	DB.Connect().Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "user_id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"latitude":  location.Latitude,
			"longitude": location.Longitude,
			"city":      location.City,
			"timestamp": location.Timestamp,
		}),
	}).Create(&location)
	return location, nil
}

func Read(userId string) (Location, error) {
	var locations Location
	dbRresult := DB.Connect().First(&locations, "user_id = ?", userId)
	if errors.Is(dbRresult.Error, gorm.ErrRecordNotFound) {
		return locations, fmt.Errorf("location doesn't exist")
	}
	return locations, nil
}
