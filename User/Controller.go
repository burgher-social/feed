package User

import (
	"fmt"
	"os"

	DB "burgher/Storage/PSQL"

	Token "burgher/Token"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(string(os.Getenv("JWT_SIGNING")))

// fmt.Println(secretKey)

func createToken(mp map[string]interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims(mp))
	// {
	// 	"username": username,
	// 	"exp":      time.Now().Add(time.Hour * 24).Unix(),
	// })

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func create(user User) (User, string, string, error) {
	fmt.Println(user)

	DB.Connect().Create(&user)
	accessToken, refreshToken := Token.GenerateTokens(user.Id)
	return user, accessToken, refreshToken, nil
}

func read(username string, tag int) (User, error) {
	var user User
	DB.Connect().First(&user, "user_name = ? and tag = ?", username, tag)

	return user, nil
}

func readWithEmail(email string) (User, *string, *string, error) {
	var user User
	err := DB.Connect().First(&user, "email = ?", email)
	var accessToken, refreshToken *string = nil, nil
	if err.Error == nil {
		accessToken3, refreshToken4 := Token.GenerateTokens(user.Id)
		accessToken = &accessToken3
		refreshToken = &refreshToken4
	}
	return user, accessToken, refreshToken, err.Error
}
