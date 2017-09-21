package router

import (
	"github.com/astaxie/beego"
	"github.com/shohi/yclite/controller"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
