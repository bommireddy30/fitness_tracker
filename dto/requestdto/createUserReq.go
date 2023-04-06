package requestdto

type CreateUserReqDto struct {
	UserName    string  `json:"userName"`
	Password    string  `json:"password"`
	Mail        string  `json:"mail"`
	PhoneNumber int64   `json:"phoneNumber"`
	Age         int     `json:"age"`
	Height      float64 `json:"height"`
	Weight      float64 `json:"weight"`
}
