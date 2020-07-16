package main

import (
	"encoding/json"
	"filmTicketSeckill/Lib"
	"filmTicketSeckill/filmTicket"
	"filmTicketSeckill/order"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"sync"
)

var mu sync.Mutex

func main() {
	var errchan = make(chan error)
	go func() {
		err := ReadMessage()
		if err != nil {
			errchan <-err
		}
	}()
	<-errchan

}


//消费消息
func ReadMessage()error  {

	mq:=Lib.NewMQ()
	err :=mq.ConsumeMessage(Lib.QUEUE_TICKET,"read",Seckill)
	if err != nil {
		return fmt.Errorf("consume err:%s",err.Error())
	}
	return nil
}


//秒杀
func Seckill(mods <-chan amqp.Delivery)  {

	go func() {

		res :=&filmTicket.TicketBind{}
		for mod :=range mods {

			json.Unmarshal(mod.Body,res)
			res.ReduceFilmNum()
			order.MakeOrder(res)

			mod.Ack(false)
			log.Printf("用户%d购票成功",res.UserID)

		}
	}()


}