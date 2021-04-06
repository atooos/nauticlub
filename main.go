package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	r := gin.Default()
	r.GET("/users", GetUsers)
	r.POST("/users", CreateUser)
	r.DELETE("/users/:uuid", DeleteUser)
	r.PUT("/users/:uuid", UpdateUser)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

var UserList = map[string]*User{}

func GetUsers(ctx *gin.Context) {
	ctx.JSON(200, UserList)
}

func CreateUser(ctx *gin.Context) {
	var u User
	err := ctx.BindJSON(&u)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	err = u.ValidCreatePayload()
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	u.ID = uuid.NewString()
	UserList[u.ID] = &u
	ctx.JSON(200, u)
}

func UpdateUser(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	_, ok := UserList[uuid]
	if !ok {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}

	var u User
	err := ctx.BindJSON(&u)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	err = u.ValidUpdatePayload()
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	UserList[uuid] = &u
	ctx.JSON(http.StatusAccepted, u)
}

func DeleteUser(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	_, ok := UserList[uuid]
	if !ok {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}
	delete(UserList, uuid)
	ctx.JSON(http.StatusAccepted, nil)
}

type User struct {
	ID          string    `json:"uuid"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	BirthDate   time.Time `json:"birth_date"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
}

func (u *User) ValidCreatePayload() error {
	if len(u.FirstName) == 0 {
		return errors.New("empty first name")
	}
	return nil
}

func (u *User) ValidUpdatePayload() error {
	return nil
}
