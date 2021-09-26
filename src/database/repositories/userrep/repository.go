package userrep

import "github.com/ArtemGretsov/golang-server-template/src/database"

type RepositoryInterface interface {
	GetUserByLogin(login string) (User, error)
	SaveUser(userData User) (User, error)
}

type repository struct{}

var Repository RepositoryInterface = &repository{}

func (repository) GetUserByLogin(login string) (user User, err error) {
	db := database.DB()
	err = db.Get(&user, "select * from users where login=$1", login)
	return
}

func (repository) SaveUser(userData User) (user User, err error) {
	db := database.DB()
	_, err = db.NamedExec("insert into users(name, login, password) values(:name, :login, :password)", &userData)
	return
}