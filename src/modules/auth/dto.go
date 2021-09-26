package auth

type SignupReqDto struct {
	Name     string `validate:"required"`
	Login    string `validate:"required" json:"login"`
	Password string `validate:"required"`
}

type SignupResDto struct {
	ID    int
	Name  string
	Login string
	Token string
}
