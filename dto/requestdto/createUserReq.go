package requestdto

type CreateUserReqDto struct {
	UserId      string  `bson:"user_Id" json:"userId"`
	UserName    string  `bson:"userName" json:"userName"`
	Password    string  `bson:"password" json:"password"`
	Mail        string  `bson:"mail" json:"mail"`
	PhoneNumber int64   `bson:"phoneNumber" json:"phoneNumber"`
	Age         int     `bson:"age" json:"age"`
	Height      float64 `bson:"height" json:"height"`
	Weight      float64 `bson:"weight" json:"weight"`
}
