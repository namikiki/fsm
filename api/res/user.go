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
