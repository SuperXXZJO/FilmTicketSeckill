package filmTicket

import (
	"github.com/gin-gonic/gin"
)

type TicketBind struct {
	FilmName  string	`json:"film_name" binding:"required"`
	Location  string	`json:"location" binding:"required"`
	Time      string  `json:"time" binding:"required"`
	UserID    uint		`json:"user_id" binding:"required"`

}

func NewTicketItem()*TicketBind{
	return &TicketBind{}
}

//购票
func BuyTicket(c *gin.Context)  {

	item := NewTicketItem()
	if err:=c.BindJSON(item);err!=nil{
		c.JSON(401,gin.H{
			"message":err.Error(),
		})
		return
	}

	if err:=item.BuyFilmTicket();err!=nil{
		c.JSON(500,gin.H{
			"message":err.Error(),
		})
		return
	}
	c.JSON(200,gin.H{
		"message": "购票成功",
	})

}

