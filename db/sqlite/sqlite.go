package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/atooos/nauticlub/db"
	"github.com/atooos/nauticlub/model"
)

var _ db.Storage = &DB{}

type DB struct {
	conn *gorm.DB
}

func (db *DB) SetConn(conn *gorm.DB) {
	db.conn = conn
}

func New(fileName string) db.Storage {
	db, err := gorm.Open(sqlite.Open(fileName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Sub{})

	return &DB{
		conn: db,
	}
}
