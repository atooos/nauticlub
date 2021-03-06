package service

import (
	"log"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWTMiddlware(secretKey string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.Request.Header.Get("Authorization")
		if len(auth) == 0 || !strings.ContainsAny(auth, "Bearer ") {
			log.Println("auth empty or not valid")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, nil)
			return
		}

		auth = strings.ReplaceAll(auth, "Bearer ", "")
		var kf jwt.Keyfunc = func(*jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		}

		token, err := jwt.Parse(auth, kf)
		if err != nil || !token.Valid {
			log.Println("token jwt not valid", err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, nil)
			return
		}

		userUUID := token.Claims.(jwt.MapClaims)["id"]
		log.Println("claims uuid user", userUUID)
		ctx.Set("user_uuid", userUUID)
	}
}
