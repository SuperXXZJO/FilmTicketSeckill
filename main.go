package main

import (
	"filmTicketSeckill/Lib"
	"filmTicketSeckill/filmTicket"
	"fmt"
	"github.com/gin-gonic/gin"


)


func main() {

	filmTicket.GetFilmMap()

	router :=gin.Default()
	router.POST("/buyticket",filmTicket.BuyTicket)

	errchan :=make(chan error)
	go func() {
		err:=Lib.SeckillInit()
		if err != nil {
			fmt.Println(err)
			errchan <-err
		}
	}()
	go func() {
		err:=router.Run()
		if err != nil {
			errchan <- err
		}
	}()
	<-errchan

}
