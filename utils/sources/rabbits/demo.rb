#!/usr/bin/env ruby

=begin
  Author: 曾涛
  Desc:   BetOrder 的 RabbitMQ 消息 Demo
  Date:   2019-05-09
  Email:  zengtao@risewinter.com
=end

require "bunny"
require "pry"

conn = Bunny.new
conn.start

channel = conn.create_channel
exchange = channel.direct("rw-hz-odds-direct")

def messages
  {
    id:    10,
    event: "Created"
  }.to_json
end

exchange.publish(messages, routing_key: "datacenter-kalista-order-event")
conn.close