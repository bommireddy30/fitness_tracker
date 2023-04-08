package responsedto

type LoginResp struct {
	ErrorDescription string `json:"errorDescription"`
	Status           string `json:"status"`
	Token            string `json:"token"`
}
