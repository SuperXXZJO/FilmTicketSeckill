package order

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var(
	DB *gorm.DB

)

func InitDB(){
	db,err:=gorm.Open("mysql","root:root@/films?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)

	}
	db.AutoMigrate(&Order{})

	DB = db

}




func init() {
	InitDB()
	
}
