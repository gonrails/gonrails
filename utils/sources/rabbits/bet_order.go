/*
Author: 曾涛
Desc:   BetOrder 相关的 RabbitMQ 事件 & 抽象
Date:   2019-05-09
Email:  zengtao@risewinter.com
*/

/**
 * ! 这里约定消息通信的协议是 JSON
 */

package rabbits

import (
	"kalista/utils/common"
	"log"
)

const (
	orderExchangeName = "rw-hz-odds-direct" // Rails 那边之前就约定好的 ExchangeName
	orderExchangeType = "direct"
	orderRouteKey     = "datacenter-kalista-order-event"
)

type OrderMessage struct {
	ID    uint
	Event string
}

type Manager struct {
	OrderMessages chan *OrderMessage
}

func (manager *Manager) Exec() {

}

// RunBetOrder -
/**
需要复用连接
*/
func runBetOrder() {
	channel, err := Instance().conn.Channel()
	common.FailOnError(err, "Failed to open a Channel")
	defer channel.Close()

	err = channel.ExchangeDeclare(orderExchangeName, orderExchangeType, false, false, false, false, nil)
	common.FailOnError(err, "Failed to Declare an Exchange")

	queue, err := channel.QueueDeclare("", false, false, true, false, nil)
	common.FailOnError(err, "Failed to Declare a Queue")

	err = channel.QueueBind(queue.Name, orderRouteKey, orderExchangeName, false, nil)
	common.FailOnError(err, "Failed to Bind a Queue to the Exchange")

	msgs, err := channel.Consume(queue.Name, "", true, false, false, false, nil)
	common.FailOnError(err, "Failed to Consume the Queue")

	log.Println("waiting messags...")

	for d := range msgs {
		log.Println(string(d.Body))
	}
}
