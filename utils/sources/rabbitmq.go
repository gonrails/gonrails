package sources

import "github.com/streadway/amqp"

var (
	rabbitInstance *rabbitMQSource
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

}

func newRabbitInstance() *rabbitMQSource {

	conn, err := amqp.DialConfig(
		"",
		amqp.Config{},
	)

	if err != nil {
		panic(err)
	}

	return &rabbitMQSource{
		conn: conn,
	}
}
