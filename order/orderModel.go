package order

import (

	"github.com/jinzhu/gorm"
)
type Order struct {
	gorm.Model
	UserID uint
	FilmName string
}


//创建新的订单
func CreateNewOrder(mod *Order) error  {
	if err:=DB.Create(&mod).Error;err !=nil{
		return err
	}
	return nil
}

