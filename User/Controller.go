package User

import (
	"fmt"
	"os"
	"time"

	DB "burgher/Storage/PSQL"

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
	now := time.Now()
	accessToken, _ := createToken(map[string]interface{}{
		"iat": now.Unix(),
		"exp": now.Add(time.Hour).Unix(),
		"id":  user.Id,
	})
	refreshToken, _ := createToken(map[string]interface{}{
		"iat": now.Unix(),
		"exp": now.Add(time.Hour * 24 * 60).Unix(),
		"id":  user.Id,
	})
	return user, accessToken, refreshToken, nil
}

func read(username string, tag int) (User, error) {
	var user User
	DB.Connect().First(&user, "user_name = ? and tag = ?", username, tag)

	return user, nil
}
