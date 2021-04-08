package moke

import (
	"errors"

	dbStorage "github.com/atooos/nauticlub/db"
	"github.com/atooos/nauticlub/model"
)

var ErrNotFound = errors.New("not found")

var _ dbStorage.Storage = &DB{}

func New() dbStorage.Storage {
	return &DB{
		ListUser: map[string]*model.User{},
		ListSub:  map[string]*model.Sub{},
	}
}

type DB struct {
	ListUser map[string]*model.User
	ListSub  map[string]*model.Sub
}
