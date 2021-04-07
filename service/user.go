package service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/atooos/nauticlub/db"
	"github.com/atooos/nauticlub/model"
)

type ServiceUser struct {
	db db.Storage
}

func (su *ServiceUser) Get(ctx *gin.Context) {
	us, err := su.db.GetAllUser()
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(200, us)
}

func (su *ServiceUser) Create(ctx *gin.Context) {
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
	su.db.CreateUser(&u)
	ctx.JSON(200, u)
}

func (su *ServiceUser) Update(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	_, err := su.db.GetUser(uuid)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusNotFound, nil)
		return
	}

	var u model.User
	err = ctx.BindJSON(&u)
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

	err = su.db.UpdateUser(uuid, &u)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ctx.JSON(http.StatusAccepted, u)
}

func (su *ServiceUser) Delete(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	_, err := su.db.GetUser(uuid)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusNotFound, nil)
		return
	}

	err = su.db.DeleteUser(uuid)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ctx.JSON(http.StatusAccepted, nil)
}
