package requestdto

type UserLogin struct {
	UserName string `bson:"userName" json:"userName" validate:"required"`
	Password string `bson:"password" json:"password" validate:"required"`
}
