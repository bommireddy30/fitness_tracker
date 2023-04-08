package database

import (
	"Connection/dto/requestdto"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(userDto requestdto.CreateUserReqDto) error {
	_, err := UsersCollection.InsertOne(Ctx, userDto)
	if err != nil {
		log.Println("UsersCollection.InsertOne failed with an error:", err.Error())
		return err
	}
	return nil
}

func UserLogin(userName string) (requestdto.CreateUserReqDto, error) {

	userDto := requestdto.CreateUserReqDto{}
	filter := bson.M{}
	filter["userName"] = userName

	err := UsersCollection.FindOne(Ctx, filter).Decode(&userDto)
	if err != nil {
		log.Println("UserLogin - UsersCollection.FindOne failed with an error:", err.Error())
		return userDto, err
	}

	return userDto, nil
}
