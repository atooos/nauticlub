package service

import (
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/atooos/nauticlub/db"
	"github.com/atooos/nauticlub/model"
)

type ServiceUser struct {
	db db.Storage
}

func (su *ServiceUser) Get(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	u, err := su.db.GetUser(uuid)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusNotFound, nil)
		return
	}
	ctx.JSON(200, u)
}

func (su *ServiceUser) GetAll(ctx *gin.Context) {
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

func (su *ServiceUser) Login(ctx *gin.Context) {
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

	udb, err := su.db.GetUserByEmail(u.Email)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	if udb.Password != u.Password {
		log.Printf("login attempt with user failed: %v\n", udb.ID)
		ctx.JSON(http.StatusUnauthorized, nil)
		return
	}

	// Create the token
	clains := jwt.MapClaims{
		"id":  udb.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), clains)
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte("mySigningKey"))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"jwt": tokenString})
}
