package moke

import (
	"errors"

	dbStorage "github.com/atooos/nauticlub/db"
	"github.com/atooos/nauticlub/model"
	"github.com/google/uuid"
)

var _ dbStorage.Storage = &DB{}

func New() dbStorage.Storage {
	return &DB{
		ListUser: map[string]*model.User{},
	}
}

type DB struct {
	ListUser map[string]*model.User
}

func (db *DB) CreateUser(u *model.User) error {
	u.ID = uuid.NewString()
	db.ListUser[u.ID] = u
	return nil
}

func (db *DB) DeleteUser(uuid string) error {
	_, ok := db.ListUser[uuid]
	if !ok {
		return errors.New("not found")
	}
	delete(db.ListUser, uuid)
	return nil
}

func (db *DB) UpdateUser(uuid string, u *model.User) error {
	_, ok := db.ListUser[uuid]
	if !ok {
		return errors.New("not found")
	}
	db.ListUser[uuid] = u
	return nil
}

func (db *DB) GetUser(uuid string) (u *model.User, err error) {
	u, ok := db.ListUser[uuid]
	if !ok {
		return nil, errors.New("not found")
	}
	return u, nil
}

func (db *DB) GetAllUser() (us []model.User, err error) {
	us = make([]model.User, len(db.ListUser))
	for k := range db.ListUser {
		us = append(us, *db.ListUser[k])
	}
	return us, nil
}
