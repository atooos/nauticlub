package service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/atooos/nauticlub/model"
)

var UserList = map[string]*model.User{}

func GetUsers(ctx *gin.Context) {
	ctx.JSON(200, UserList)
}

func CreateUser(ctx *gin.Context) {
	var u model.User
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

	var u model.User
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
