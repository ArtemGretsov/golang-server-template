package userrep

import (
	"context"

	"github.com/ArtemGretsov/golang-server-template/src/database"
	"github.com/ArtemGretsov/golang-server-template/src/database/_schemagen"
)

type RepositoryInterface interface {
	SaveUser(ctx context.Context, login, name, password string) (*_schemagen.User, error)
}

type repository struct{}

var Repository RepositoryInterface = &repository{}

func (repository) SaveUser(ctx context.Context, login, name, password string) (user *_schemagen.User, err error) {
	DB := database.DB()
	user, err = DB.User.Create().SetLogin(login).SetName(name).SetPassword(password).Save(ctx)
	return
}
