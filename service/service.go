package service

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/atooos/nauticlub/db"
)

func New(port string, db db.Storage, jwtKey string) *http.Server {
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
	//r.Run(":" + port)
	return &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

}
