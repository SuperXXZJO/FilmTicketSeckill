package Lib

import (
	"github.com/streadway/amqp"
	"log"
)



type MQ struct {
	Channel *amqp.Channel
}

//新建MQ
func NewMQ()*MQ  {
	conn := NewMQConn()
	c,err := conn.Channel()
	if err != nil {
		log.Println("创建channel失败",err)
		return nil
	}
	return &MQ{Channel:c}
}

//声明队列并绑定
func (m *MQ)DeclareQueueAndB (queuename string,exchange string,key string)error  {
	q,err:=m.Channel.QueueDeclare(queuename,false,false,false,false,nil)
	if err != nil {
		return err
	}
	err=m.Channel.QueueBind(q.Name,key,exchange,false,nil)
	if err != nil {
		return err
	}
	return nil
}