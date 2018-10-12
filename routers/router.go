package routers

import (
	"github.com/shanghaiyiqi/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
