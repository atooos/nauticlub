package service

import (
	"github.com/gin-gonic/gin"

	"github.com/atooos/nauticlub/db"
)

func Init(port string, db db.Storage) {
	r := gin.Default()
	su := &ServiceUser{
		db: db,
	}
	// Users
	r.GET("/users", su.Get)
	r.POST("/users", su.Create)
	r.DELETE("/users/:uuid", su.Delete)
	r.PUT("/users/:uuid", su.Update)
	// Pdf
	r.POST("/pdf", CreatePDF)
	r.Run(":" + port)
}
