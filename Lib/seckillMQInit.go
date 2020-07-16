package Lib

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
)


const (
	QUEUE_TICKET  = "ticket"  //电影票队列
	EXCHANGE_SECKILL = "seckill"  //交换机名称
	ROUTER_KEY_SECKILL = "seckill"  //路由键名称
)


func SeckillInit()error{
	//新建MQ
	MQ:=NewMQ()
	if MQ == nil  {
		return fmt.Errorf("新建MQ失败")
	}
	//申明交换机
	err:=MQ.Channel.ExchangeDeclare(EXCHANGE_SECKILL,"direct",false,false,false,false,nil)
	if err != nil {
		return fmt.Errorf("申明交换机失败：%s",err.Error())
	}
	//声明并绑定队列
	err=MQ.DeclareQueueAndB(QUEUE_TICKET,EXCHANGE_SECKILL,ROUTER_KEY_SECKILL)
	if err != nil {
		return fmt.Errorf("声明队列%s失败：%s",QUEUE_TICKET,err.Error())
	}

	return nil
}

//发送消息
func (m *MQ) SendMessage(key string,exchange string,msg interface{}) error {
	message,_ := json.Marshal(msg)
	return m.Channel.Publish(exchange,key,false,false,amqp.Publishing{
		ContentType:     "application/json",
		Body:            message,
	})
}

//接收消息
func (m *MQ)ConsumeMessage(queuename string,key string,callback func( <-chan amqp.Delivery)) error {
	c,err :=m.Channel.Consume(queuename,key,false,false,false,false,nil)
	if err != nil {
		return fmt.Errorf("consume err：%s",err.Error())
	}
	callback(c)
	return nil
}