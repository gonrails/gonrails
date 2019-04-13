package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var (
	config *viper.Viper
	MySQL  *mySQL
	Redis  *redis
)

type mySQL struct {
	Host        string
	Port        string
	User        string
	Password    string
	Connections int32
	Idles       int32
}

type redis struct {
	Host string
	Port string
}

func init() {
	config = viper.New()
	config.SetConfigFile("./config/config.yml")

	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}

	loadMySQL()
	loadRedis()

	fmt.Println(MySQL.Host)
	fmt.Println(Redis.Host)
}

func loadMySQL() {

	t := config.GetStringMapString(os.Getenv("GO_ENV") + ".MySQL")
	MySQL = &mySQL{
		Host:     t["host"],
		Port:     t["port"],
		User:     t["user"],
		Password: t["password"],
	}
}

func loadRedis() {
	t := config.GetStringMapString(os.Getenv("GO_ENV") + ".Redis")
	Redis = &redis{
		Host: t["host"],
		Port: t["port"],
	}
}
