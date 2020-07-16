package order

import (
	"filmTicketSeckill/filmTicket"
)

//下订单
func MakeOrder(mod *filmTicket.TicketBind)  {
	m :=&Order{
		UserID: mod.UserID,
		FilmName: mod.FilmName,
	}
	CreateNewOrder(m)
}