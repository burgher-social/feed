package Token

import (
	DB "burgher/Storage/PSQL"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

var secretKey = []byte(string(os.Getenv("JWT_SIGNING")))

func TokenRefresh(refreshToken string) (string, string, error) {
	tokn, err := GetTokenClaims(refreshToken)

	if err != nil {
		return "", "", err
	}
	type User struct {
		gorm.Model
		Id         string
		Name       string
		UserName   string `json:"username"`
		Tag        int
		IsVerified bool
		Email      string  `gorm:"unique;not null"`
		ImageUrl   *string `json:"imageUrl"`
	}
	var user User
	DB.Connect().Where("user_id = ?", tokn.UserId).First(&user)
	fmt.Println("REFRESH TOKEN REQUESTING USER")
	fmt.Println(user)
	if user.Id == "" {
		return "", "", fmt.Errorf("user doesn't exist")
	}
	accessToken, refreshTokenNew := GenerateTokens(tokn.UserId)
	return accessToken, refreshTokenNew, nil
}

func GenerateTokens(userId string) (string, string) {
	now := time.Now()
	claims := TokenClaims{
		Iat:    now.Unix(),
		Exp:    now.Add(time.Hour).Unix(),
		UserId: userId,
	}
	claimsRefresh := TokenClaims{
		Iat:    now.Unix(),
		Exp:    now.Add(time.Hour * 24 * 60).Unix(),
		UserId: userId,
	}
	var myMap map[string]interface{}
	data, _ := json.Marshal(claims)
	json.Unmarshal(data, &myMap)
	accessToken, _ := createToken(myMap)
	data2, _ := json.Marshal(claimsRefresh)
	json.Unmarshal(data2, &myMap)
	refreshToken, _ := createToken(myMap)

	return accessToken, refreshToken
}

func createToken(mp map[string]interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims(mp))

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func getNewToken(refreshToken string) (TokenResponse, error) {
	accessToken, newRefreshToken, err := TokenRefresh(refreshToken)

	var tok TokenResponse
	tok.AccessToken = accessToken
	tok.RefreshToken = newRefreshToken
	return tok, err
}

func GetTokenClaims(tok string) (*TokenClaims, error) {
	now := time.Now().Unix()
	var tokenObj TokenClaims
	token, err := jwt.ParseWithClaims(tok, &tokenObj, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if now > tokenObj.Exp {
		return nil, fmt.Errorf("expired token")
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return &tokenObj, nil
}
