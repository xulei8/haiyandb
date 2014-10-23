package routers

import (
	"github.com/xulei8/haiyandb/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
