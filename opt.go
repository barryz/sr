package main

import (
	"flag"
)

var (
	broker    string
	lb        bool
	isRecv    bool
	verbose   bool
	noack     bool
	notAck    bool
	bid       int
	count     int
	qosPcount int
	qosGlobal bool
	ndq       bool
)

func init() {
	flag.StringVar(&broker, "broker", "", "Broker connect string")
	flag.BoolVar(&lb, "l", false, "List all brokers")
	flag.BoolVar(&isRecv, "r", false, "Receive messages")
	flag.BoolVar(&verbose, "v", false, "Verbose log info")
	flag.BoolVar(&noack, "noack", false, "Consume with noack")
	flag.BoolVar(&notAck, "not-ack", false, "Not send ack")
	flag.IntVar(&bid, "b", 0, "Broker ID")
	flag.IntVar(&count, "c", 0, "Message count")
	flag.IntVar(&qosPcount, "qos-pc", 0, "Qos pretch count")
	flag.BoolVar(&qosGlobal, "qos-g", false, "Qos global")
	flag.BoolVar(&ndq, "ndq", false, "None-durable queue")
	flag.Parse()
}
