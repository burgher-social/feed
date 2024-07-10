package User

import (
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"

	DB "burgher/Storage/PSQL"
	Utils "burgher/Utils"

	Token "burgher/Token"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm/clause"
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

// Used to read user with email. It returns tokens so only use when needs to authenticate.
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

func updateProfilePicture(userId string, file *multipart.File) {
	data, err := io.ReadAll(*file)
	if err != nil {
		log.Println(err)
		return
		// return "", "", err
	}

	// contentType := http.DetectContentType(data)
	imgBase64Str := base64.StdEncoding.EncodeToString(data)
	imageobj := UserImage{
		Id:     Utils.GenerateId(),
		UserId: userId,
		Image:  imgBase64Str,
	}

	DB.Connect().Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "user_id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"image": imgBase64Str,
		}),
	}).Create(&imageobj)
}
