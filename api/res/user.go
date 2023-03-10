package res

type UserLogin struct {
	Uid string `json:"uid"`
	//Jwt string
}

type UserRegister struct {
	UserName string
	ID       string
	Token    string
}

type Login struct {
	Token  string `json:"token"`
	UserID string `json:"user_id"`
}
