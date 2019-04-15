package config

import (
	"os"
	"strconv"

	"github.com/spf13/viper"
)

var (
	config *viper.Viper
	MySQL  *mySQL
	Redis  *redis
	Rabbit *rabbit
)

type mySQL struct {
	Host        string
	Port        string
	Name        string
	Username    string
	Password    string
	Connections int
	Idles       int
}

type redis struct {
	Host string
	Port string
}

type rabbit struct {
	Host     string
	Port     string
	User     string
	Password string
}

func init() {
	config = viper.New()
	config.SetConfigFile("./config/config.yml")

	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}

	loadMySQL()
	loadRedis()
	loadRabbit()
}

func loadMySQL() {

	t := config.GetStringMapString(os.Getenv("GO_ENV") + ".MySQL")
	connections, _ := strconv.Atoi(t["connections"])
	idles, _ := strconv.Atoi(t["idles"])

	MySQL = &mySQL{
		Host:        t["host"],
		Port:        t["port"],
		Name:        t["name"],
		Username:    t["username"],
		Password:    t["password"],
		Connections: connections,
		Idles:       idles,
	}
}

func loadRedis() {
	t := config.GetStringMapString(os.Getenv("GO_ENV") + ".Redis")
	Redis = &redis{
		Host: t["host"],
		Port: t["port"],
	}
}

func loadRabbit() {
	t := config.GetStringMapString(os.Getenv("GO_ENV") + ".RabbitMQ")
	Rabbit = &rabbit{
		Host:     t["host"],
		Port:     t["port"],
		User:     t["user"],
		Password: t["password"],
	}
}
