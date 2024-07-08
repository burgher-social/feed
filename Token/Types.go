package Token

import (
	DB "burgher/Storage/PSQL"

	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	UserId       string
	RefreshToken string
}

type TokenResponse struct {
	RefreshToken string
	AccessToken  string
}

func init() {
	DB.Connect().AutoMigrate(&Token{})
}
