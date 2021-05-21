package main

import (
	"PROJECT-INDIVIDU-GO-REACT/entity"
	// "PROJECT-INDIVIDU-GO-REACT/handler"
	// "PROJECT-INDIVIDU-GO-REACT/user"
	"PROJECT-INDIVIDU-GO-REACT/config"
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
