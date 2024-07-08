package Token

import (
	DB "burgher/Storage/PSQL"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(string(os.Getenv("JWT_SIGNING")))

func TokenRefresh(refreshToken string) (error, string, string) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	var tokenObj Token
	DB.Connect().Create(&tokenObj)

	if err != nil {
		return err, "", ""
	}

	if !token.Valid {
		fmt.Println("invalid token")
		return fmt.Errorf("invalid token"), "", ""
	}
	accessToken, refreshTokenNew := GenerateTokens(tokenObj.UserId)
	return nil, accessToken, refreshTokenNew
}

func GenerateTokens(userId string) (string, string) {
	now := time.Now()
	accessToken, _ := createToken(map[string]interface{}{
		"iat": now.Unix(),
		"exp": now.Add(time.Hour).Unix(),
		"id":  userId,
	})
	refreshToken, _ := createToken(map[string]interface{}{
		"iat": now.Unix(),
		"exp": now.Add(time.Hour * 24 * 60).Unix(),
		"id":  userId,
	})

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

func getNewToken(refreshToken string) (error, TokenResponse) {
	err, accessToken, newRefreshToken := TokenRefresh(refreshToken)

	var tok TokenResponse
	tok.AccessToken = accessToken
	tok.RefreshToken = newRefreshToken
	return err, tok
}
