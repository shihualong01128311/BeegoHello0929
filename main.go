package main

import (
	"BeegoHello0929/db_mysql"
	_ "BeegoHello0929/routers"
	"github.com/astaxie/beego"
)

func main() {
	db_mysql.Connect()

	beego.Run()
}
