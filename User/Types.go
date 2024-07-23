package User

import (
	DB "burgher/Storage/PSQL"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id         string
	Name       string
	UserName   string `json:"username"`
	Tag        int
	IsVerified bool
	Email      string
	ImageUrl   *string `json:"imageUrl"`
}

type UserImage struct {
	gorm.Model
	Id     string
	UserId string `gorm:"unique;not null"`
	Image  string
}

type LikesPosts struct {
	gorm.Model
	UserId string `gorm:"index:u_user_id_post_id,unique" json:"userId"`
	PostId string `gorm:"index:u_user_id_post_id,unique" json:"postId"`
}

func init() {
	DB.Connect().AutoMigrate(&User{})
	DB.Connect().AutoMigrate(&UserImage{})
	DB.Connect().AutoMigrate(&LikesPosts{})
}

type UserRequest struct {
	Name                string `json:"name"`
	UserName            string `json:"username"`
	Tag                 int    `json:"tag"`
	Email               string `json:"email"`
	FirebaseAuthIdToken string `json:"firebaseAuthIdToken"`
}

type UserResponse struct {
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	UserName     string  `json:"username"`
	Tag          int     `json:"tag"`
	IsVerified   bool    `json:"isVerified"`
	Email        string  `json:"email"`
	RefreshToken *string `json:"refreshToken"`
	AccessToken  *string `json:"accessToken"`
	ImageUrl     *string `json:"imageUrl"`
}
