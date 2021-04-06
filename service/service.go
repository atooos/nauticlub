package service

import "github.com/gin-gonic/gin"

func Init() {
	r := gin.Default()
	r.GET("/users", GetUsers)
	r.POST("/users", CreateUser)
	r.DELETE("/users/:uuid", DeleteUser)
	r.PUT("/users/:uuid", UpdateUser)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
