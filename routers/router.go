package routers

import (
	"BeegoHello0929/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/register", &controllers.RegisterController{})
}
