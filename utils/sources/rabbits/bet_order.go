// /*
// Author: 曾涛
// Desc:   BetOrder 相关的 RabbitMQ 事件 & 抽象
// Date:   2019-05-09
// Email:  zengtao@risewinter.com
// */
//
package rabbits

//
// import (
// 	"encoding/json"
//
// 	"github.com/gonrails/gonrails/utils/common"
// )
//
// const (
// 	orderExchangeName = "rw-hz-odds-direct" // Rails 那边之前就约定好的 ExchangeName
// 	orderExchangeType = "direct"
// 	orderRouteKey     = "datacenter-kalista-order-event"
// )
//
// var (
// 	orderChan = make(chan *OrderMessage)
// )
//
// // OrderMessage 抽象通信内容
// /**
//  * ! 这里约定消息通信的协议是 JSON
//  */
// type OrderMessage struct {
// 	ID    uint
// 	Event string
// }
//
// // RunBetOrder -
// /**
// 需要复用连接
// */
// func runBetOrder() {
// 	channel, err := theInstance().conn.Channel()
// 	common.FailOnError(err, "Failed to open a Channel")
// 	defer channel.Close()
//
// 	err = channel.ExchangeDeclare(orderExchangeName, orderExchangeType, false, false, false, false, nil)
// 	common.FailOnError(err, "Failed to Declare an Exchange")
//
// 	queue, err := channel.QueueDeclare("", false, false, true, false, nil)
// 	common.FailOnError(err, "Failed to Declare a Queue")
//
// 	err = channel.QueueBind(queue.Name, orderRouteKey, orderExchangeName, false, nil)
// 	common.FailOnError(err, "Failed to Bind a Queue to the Exchange")
//
// 	msgs, err := channel.Consume(queue.Name, "", true, false, false, false, nil)
// 	common.FailOnError(err, "Failed to Consume the Queue")
//
// 	var message OrderMessage
// 	for d := range msgs {
// 		json.Unmarshal(d.Body, &message)
// 		orderChan <- &message
// 	}
// }
