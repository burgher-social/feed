package User

import (
	"fmt"

	DB "burgher/Storage/PSQL"
)

func create(user User) (User, error) {
	fmt.Println(user)

	DB.Connect().Create(&user)
	return user, nil
}

func read(username string, tag int) (User, error) {
	var user User
	DB.Connect().First(&user, "user_name = ? and tag = ?", username, tag)
	return user, nil
}
