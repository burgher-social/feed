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
}

func init() {
	DB.Connect().AutoMigrate(&User{})
}

type UserRequest struct {
	Name     string `json:"name"`
	UserName string `json:"username"`
	Tag      int    `json:"tag"`
}

type UserResponse struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	UserName   string `json:"username"`
	Tag        int    `json:"tag"`
	IsVerified bool   `json:"isVerified"`
}
