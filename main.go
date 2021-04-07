package main

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/atooos/nauticlub/db"
	"github.com/atooos/nauticlub/db/moke"
	"github.com/atooos/nauticlub/service"
)

type Config struct {
	Port   string
	JWTKey string
	Env    string
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
}

func main() {
	var db db.Storage
	if config.Env == "local" {
		db = moke.New()
	}
	service.Init(config.Port, db, config.JWTKey)
}
