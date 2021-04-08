package moke

import (
	"time"

	"github.com/google/uuid"

	"github.com/atooos/nauticlub/model"
)

func (db *DB) CreateSub(s *model.Sub) error {
	s.ID = uuid.NewString()
	s.CreatedAt = time.Now()
	db.ListSub[s.ID] = s
	return nil
}

func (db *DB) DeleteSub(uuid string) error {
	_, ok := db.ListSub[uuid]
	if !ok {
		return ErrNotFound
	}
	delete(db.ListSub, uuid)
	return nil
}

func (db *DB) UpdateSub(uuid string, s *model.Sub) error {
	_, ok := db.ListSub[uuid]
	if !ok {
		return ErrNotFound
	}
	s.UpdateAt = time.Now()
	db.ListSub[uuid] = s
	return nil
}

func (db *DB) GetSub(uuid string) (s *model.Sub, err error) {
	s, ok := db.ListSub[uuid]
	if !ok {
		return nil, ErrNotFound
	}
	return s, nil
}

func (db *DB) GetAllSub() (ss []model.Sub, err error) {
	ss = make([]model.Sub, len(db.ListSub))
	for k := range db.ListSub {
		ss = append(ss, *db.ListSub[k])
	}
	return ss, nil
}
