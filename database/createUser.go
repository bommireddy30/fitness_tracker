package database

import (
	"Connection/dto/requestdto"
	"log"
)

func CreateUser(userDto requestdto.CreateUserReqDto) error {
	_, err := UsersCollection.InsertOne(Ctx, userDto)
	if err != nil {
		log.Println("UsersCollection.InsertOne failed with an error:", err.Error())
		return err
	}
	return nil
}
