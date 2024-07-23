package User

import (
	"encoding/base64"
	"io"
	"log"
	"mime/multipart"

	DB "burgher/Storage/PSQL"
	Utils "burgher/Utils"

	Token "burgher/Token"

	"gorm.io/gorm/clause"
)

func create(user User, firebaseAuthIdToken string) (User, *string, *string, error) {
	claims, errtok := Utils.VerifyToken(firebaseAuthIdToken, user.Email)
	if errtok != nil {
		return user, nil, nil, errtok
	}
	if (*claims)["picture"] == nil {
		tempUrl := "https://miro.medium.com/v2/resize:fit:720/format:webp/1*EOOeLlRAPdk2k4krTI5HIg.png"
		user.ImageUrl = &tempUrl
	} else {
		tempPicture := (*claims)["picture"].(string)
		user.ImageUrl = &tempPicture
	}

	DB.Connect().Create(&user)
	accessToken, refreshToken := Token.GenerateTokens(user.Id)
	return user, &accessToken, &refreshToken, nil
}

func read(username string, tag int, userId string) (User, error) {
	var user User
	if userId != "" {
		DB.Connect().First(&user, "id = ?", userId)

	} else {
		DB.Connect().First(&user, "user_name = ? and tag = ?", username, tag)
	}

	return user, nil
}

// Used to read user with email. It returns tokens so only use when needs to authenticate.
func readWithEmail(email string, firebaseAuthIdToken string) (User, *string, *string, error) {
	var user User
	_, errtok := Utils.VerifyToken(firebaseAuthIdToken, email)
	if errtok != nil {
		return user, nil, nil, errtok
	}

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
	}
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
