package Topic

import (
	"fmt"

	DB "burgher/Storage/PSQL"
)

func create(topic Topic) (Topic, error) {
	fmt.Println(topic)

	DB.Connect().Create(&topic)
	return topic, nil
}

func read(topic string) ([]Topic, error) {
	var topics []Topic
	fmt.Println(topic)
	DB.Connect().Where("name = ?", topic).Find(&topics)
	return topics, nil
}

func readAll() ([]Topic, error) {
	var topics []Topic
	DB.Connect().Find(&topics)
	return topics, nil
}
