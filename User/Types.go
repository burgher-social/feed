package User

import (
	DB "burgher/Storage/PSQL"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id         string
	Name       string
	UserName   string
	Tag        int
	IsVerified bool
	Email      string
}

type UserImage struct {
	gorm.Model
	Id     string
	UserId string `gorm:"unique;not null"`
	Image  string
}

func init() {
	DB.Connect().AutoMigrate(&User{})
	DB.Connect().AutoMigrate(&UserImage{})
}

type UserRequest struct {
	Name     string `json:"name"`
	UserName string `json:"username"`
	Tag      int    `json:"tag"`
	Email    string `json:"email"`
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
}
