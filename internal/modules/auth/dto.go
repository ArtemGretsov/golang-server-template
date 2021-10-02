package auth

type SignupReqDto struct {
	Name     string `validate:"required" json:"name"`
	Login    string `validate:"required" json:"login"`
	Password string `validate:"required" json:"password"`
}

type SignupResDto struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Login string `json:"login"`
	Token string `json:"token"`
}

type SigninReqDto struct {
	Login    string `validate:"required" json:"login"`
	Password string `validate:"required" json:"password"`
}

type SigninResDto struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Login string `json:"login"`
	Token string `json:"token"`
}

type CurrentUserReqDto struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Login string `json:"login"`
}