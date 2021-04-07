package db

import "github.com/atooos/nauticlub/model"

type Storage interface {
	CreateUser(u *model.User) error
	DeleteUser(uuid string) error
	UpdateUser(uuid string, u *model.User) error
	GetUser(uuid string) (u *model.User, err error)
	GetAllUser() (us []model.User, err error)
}
