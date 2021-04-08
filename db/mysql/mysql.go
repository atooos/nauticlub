package mysql

import (
	"fmt"

	sql "gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/atooos/nauticlub/db"
	"github.com/atooos/nauticlub/db/sqlite"
	"github.com/atooos/nauticlub/model"
)

type DB = sqlite.DB

func New(dbname, user, pass string) db.Storage {
	dsn := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, pass, dbname)
	conn, err := gorm.Open(sql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	conn.AutoMigrate(&model.User{})
	conn.AutoMigrate(&model.Sub{})

	var db DB
	db.SetConn(conn)
	return &db
}
