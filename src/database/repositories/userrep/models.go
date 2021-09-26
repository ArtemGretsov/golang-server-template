package userrep

type User struct {
	ID       uint   `db:"id"`
	Name     string `db:"name"`
	Login    string `db:"login"`
	Password string `db:"password"`
}

