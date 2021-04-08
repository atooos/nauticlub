package service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/atooos/nauticlub/db"
	"github.com/atooos/nauticlub/model"
)

type ServiceSub struct {
	db     db.Storage
	jwtKey string
}

func (su *ServiceSub) Get(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	u, err := su.db.GetSub(uuid)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusNotFound, nil)
		return
	}
	ctx.JSON(200, u)
}

func (su *ServiceSub) GetAll(ctx *gin.Context) {
	us, err := su.db.GetAllSub()
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(200, us)
}

func (su *ServiceSub) Create(ctx *gin.Context) {
	var u model.Sub
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
	su.db.CreateSub(&u)
	ctx.JSON(200, u)
}

func (su *ServiceSub) Update(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	_, err := su.db.GetSub(uuid)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusNotFound, nil)
		return
	}

	var u model.Sub
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

	err = su.db.UpdateSub(uuid, &u)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ctx.JSON(http.StatusAccepted, u)
}

func (su *ServiceSub) Delete(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	_, err := su.db.GetSub(uuid)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusNotFound, nil)
		return
	}

	err = su.db.DeleteSub(uuid)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ctx.JSON(http.StatusAccepted, nil)
}
