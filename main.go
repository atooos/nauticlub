package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"

	"github.com/atooos/nauticlub/db"
	"github.com/atooos/nauticlub/db/moke"
	"github.com/atooos/nauticlub/db/mysql"
	"github.com/atooos/nauticlub/db/sqlite"
	"github.com/atooos/nauticlub/service"
)

type Config struct {
	Port   string
	JWTKey string
	Env    string
	DB     struct {
		Name string
		User string
		Pass string
	}
}

var config Config

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	config.Env = viper.GetString("env")
	config.JWTKey = viper.GetString("jwt_key")
	config.Port = viper.GetString("port")

	// db conn
	config.DB.Name = viper.GetString("db.name")
	config.DB.User = viper.GetString("db.user")
	config.DB.Pass = viper.GetString("db.pass")
}

func main() {
	var db db.Storage
	if config.Env == "local" {
		db = moke.New()
	} else if config.Env == "qa" {
		db = sqlite.New("local.db")
	} else if config.Env == "prod" {
		db = mysql.New(config.DB.Name, config.DB.User, config.DB.Pass)
	}

	srv := service.New(config.Port, db, config.JWTKey)
	gracefullshutdown(srv)
}

func gracefullshutdown(srv *http.Server) {
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
