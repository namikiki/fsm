package req

type UserLogin struct {
	Email    string `json:"email" validate:"required"`
	PassWord string `json:"password" validate:"required"`
}

type UserRegister struct {
	Email    string `json:"email,omitempty" validate:"required,min=10,email"`
	PassWord string `json:"password,omitempty" validate:"required,min=10"`
	UserName string `json:"username,omitempty" validate:"required,min=5"`
}
