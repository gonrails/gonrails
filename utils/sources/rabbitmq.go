package sources

import (
	"fmt"
	"kalista/config"
	"kalista/utils/common"

	"github.com/streadway/amqp"
)

// 这里的 RabbitMQ 用来消费队列里面的消息

var (
	rabbitInstance *rabbitMQSource
)

const (
	exchangeName = ""
	exchangeType = "direct"
	routeKey     = ""
)

type rabbitMQSource struct {
	conn *amqp.Connection
}

func RabbitInstance() *rabbitMQSource {
	if rabbitInstance == nil {
		rabbitInstance = newRabbitInstance()
	}
	return rabbitInstance
}

func runRabbit() {
	channel, err := RabbitInstance().conn.Channel()
	common.FailOnError(err, "Failed to open a Channel")
	defer channel.Close()

	// durable RabbitMQ 重启之后、Queue 会被自动重新 Declare
	// auto-delete 如果没有 Consumer 连接, Queue 会被自动删除
	// Exclusive 专用的（对于 Connection 来说）
	_, err = channel.QueueDeclare("", true, false, true, false, nil)
	common.FailOnError(err, "Failed to Declare a Queue")

	//err = channel.QueueBind(queue.Name, routeKey, exchangeName, false, nil)
	//common.FailOnError(err, "Failed to bind to a queue")

}

func newRabbitInstance() *rabbitMQSource {

	conn, err := amqp.DialConfig(
		fmt.Sprintf("amqp://%s:%s@%s:%s/", config.Rabbit.User, config.Rabbit.Password, config.Rabbit.Host, config.Rabbit.Port),
		amqp.Config{},
	)

	if err != nil {
		panic(err)
	}

	return &rabbitMQSource{
		conn: conn,
	}
}
