package filmTicket

import (
	"filmTicketSeckill/Lib"
	"fmt"
	"log"
)


var itemMap = make( map[string]*FilmTicket )


//定时从数据库中获取秒杀电影票列表
func GetFilmMap()  {
	ticketMap,err:=FindAllFilmTicket()
	if err != nil {
		log.Println(err)
		return
	}
	for _,v :=range ticketMap{
		itemMap[v.FilmName] = v
	}
}


//用户购买电影票
func (t *TicketBind) BuyFilmTicket()error  {

	//
	if _,ok :=itemMap[t.FilmName] ;!ok{
		return fmt.Errorf("电影不存在")
	}else if ok {
		mod:=itemMap[t.FilmName]
		if mod.Num <= 0 {
			return fmt.Errorf("电影票已经售完")
		}
	}

	//消息入列
	mq :=Lib.NewMQ()
	err := mq.SendMessage(Lib.ROUTER_KEY_SECKILL,Lib.EXCHANGE_SECKILL,&t)
	if err != nil {
		return fmt.Errorf("发送消息错误：%s",err.Error())
	}

	mod:=itemMap[t.FilmName]
	mod.Num --
	return nil
}


//减少电影票库存数量
func (t *TicketBind) ReduceFilmNum() error  {
	err := ReduceFilmTicketNum(t.FilmName)
	if err != nil {
		return err
	}
	return nil
}

