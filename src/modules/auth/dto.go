package auth

type SignupReqDto struct {
	Name     string `validate:"required" json:"name"`
	Login    string `validate:"required" json:"login"`
	Password string `validate:"required" json:"password"`
}

type SignupResDto struct {
	ID    int
	Name  string
	Login string
	Token string
}
