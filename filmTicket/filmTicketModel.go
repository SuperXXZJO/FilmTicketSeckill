package filmTicket

import (

	"github.com/jinzhu/gorm"
)

type FilmTicket struct {
	gorm.Model
	FilmName string
	Loc string
	Num int
}

//查询所有电影票
func FindAllFilmTicket() (Tickets []*FilmTicket,err error) {
	if err=DB.Find(&Tickets).Error;err!=nil{
		return nil,err
	}
	return Tickets,nil
}



//通过电影名查询电影票
func FindFilmTicketByFilmName(filmname string)(res *FilmTicket,err error)  {

	if err = DB.Where("film_name = ? ",filmname).First(&res).Error;err!=nil{
		return &FilmTicket{},err
	}
	return res,nil

}

//减少电影票数量
func ReduceFilmTicketNum(filmname string)error{
	err:=DB.Model(&FilmTicket{}).Where("film_name = ?",filmname).UpdateColumn("num",gorm.Expr("num - ?",1)).Error
	if err != nil {
		return err
	}
	return nil
}