package service

import (
	"github.com/gin-gonic/gin"

	"github.com/atooos/nauticlub/db"
)

func Init(port string, db db.Storage, jwtKey string) {
	r := gin.Default()
	su := &ServiceUser{
		db:     db,
		jwtKey: jwtKey,
	}
	// Users
	r.GET("/users", su.GetAll)
	r.GET("/users/:uuid", su.Get)
	r.POST("/users", su.Create)
	r.DELETE("/users/:uuid", JWTMiddlware(jwtKey), su.Delete)
	r.PUT("/users/:uuid", su.Update)
	r.POST("/login", su.Login)
	// Pdf
	r.POST("/pdf", CreatePDF)
	r.Run(":" + port)
}
