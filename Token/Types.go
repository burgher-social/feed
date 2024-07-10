package Token

import (
	DB "burgher/Storage/PSQL"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	UserId       string `json:"userId"`
	RefreshToken string `json:"refreshToken"`
}

type TokenClaims struct {
	jwt.RegisteredClaims
	UserId       string `json:"userId"`
	RefreshToken string `json:"refreshToken"`
	Iat          int64  `json:"iat"`
	Exp          int64  `json:"exp"`
}

type TokenResponse struct {
	RefreshToken string
	AccessToken  string
}

func init() {
	DB.Connect().AutoMigrate(&Token{})
}
