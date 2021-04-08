package db

import "github.com/atooos/nauticlub/model"

type Storage interface {
	StorageUser
	StorageSub
}

type StorageUser interface {
	CreateUser(u *model.User) error
	DeleteUser(uuid string) error
	UpdateUser(uuid string, u *model.User) error
	GetUser(uuid string) (u *model.User, err error)
	GetAllUser() (us []model.User, err error)
	GetUserByEmail(email string) (u *model.User, err error)
}
type StorageSub interface {
	CreateSub(u *model.Sub) error
	DeleteSub(uuid string) error
	UpdateSub(uuid string, u *model.Sub) error
	GetSub(uuid string) (u *model.Sub, err error)
	GetAllSub() (us []model.Sub, err error)
}
