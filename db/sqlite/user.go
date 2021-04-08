package sqlite

import (
	"github.com/google/uuid"

	"github.com/atooos/nauticlub/model"
)

func (db *DB) CreateUser(u *model.User) error {
	u.ID = uuid.NewString()
	return db.conn.Create(u).Error
}

func (db *DB) DeleteUser(uuid string) error {
	return db.conn.Delete(&model.User{ID: uuid}, uuid).Error
}

func (db *DB) UpdateUser(uuid string, u *model.User) error {
	return db.conn.Model(&model.User{}).Updates(model.User{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}).Error
}

func (db *DB) GetUser(uuid string) (*model.User, error) {
	var u model.User = model.User{ID: uuid}
	return &u, db.conn.First(&u).Error
}

func (db *DB) GetAllUser() (us []model.User, err error) {
	// Get all records
	return us, db.conn.Find(&us).Error
}

func (db *DB) GetUserByEmail(email string) (u *model.User, err error) {
	return u, db.conn.Where("email = ?", email).First(u).Error
}
