package Lib

import (
	"github.com/streadway/amqp"
	"log"
)

var conn *amqp.Connection
func InitMQ()  {
	connmq,err:=amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)

	}
	conn=connmq

}

func NewMQConn()*amqp.Connection  {
	return conn
}

func init() {
	InitMQ()
}