/*
Author: 曾涛
Desc:   RabbitMQ & 以及其连接的信息 - 复用连接的 Channel 放在具体的业务逻辑里面 黄粱一梦二十年
Date:   2019-05-08
Email:  zengtao@risewinter.com
*/

package rabbits

import (
	"fmt"
	"kalista/config"

	"github.com/streadway/amqp"
)

// 这里使用单例模式

var (
	instance *rabbitMQSource
)

const (
	exchangeName = ""
	exchangeType = "direct"
	routeKey     = ""
)

type rabbitMQSource struct {
	conn *amqp.Connection
}

// Instance returns the singleton instance
func Instance() *rabbitMQSource {
	if instance == nil {
		instance = newInstance()
	}
	return instance
}

func RunRabbit() {
	runBetOrder()
}

func newInstance() *rabbitMQSource {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", config.Rabbit.User, config.Rabbit.Password, config.Rabbit.Host, config.Rabbit.Port))

	if err != nil {
		panic(err)
	}

	return &rabbitMQSource{
		conn: conn,
	}
}
