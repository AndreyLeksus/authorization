package storage

import (
	"authorization/model"
	"encoding/json"
	"fmt"
	"os"
)

func SavedUsers(users []model.Registr) {
	file, err := json.MarshalIndent(users, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.WriteFile("Users.json", file, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Файл успешно записан")
}

func LoadUsers() ([]model.Registr, error) {
	var users []model.Registr

	data, err := os.ReadFile("Users.json")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println(err)
		}
	}

	err = json.Unmarshal(data, &users)
	if err != nil {
		return nil, err
	}

	fmt.Println("Файл успешно прочитан")

	return users, nil
}
