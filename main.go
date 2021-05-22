package main

import (
	"PROJECT-INDIVIDU-GO-REACT/config"
	"PROJECT-INDIVIDU-GO-REACT/entity"
	"fmt"
)

var DB = config.Connect()

func main() {
	var Users []entity.User

	if err := DB.Find(&Users).Error; err != nil {
		panic(err.Error())
	}

	fmt.Println(Users)
}
