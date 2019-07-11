package routers

import (
	"github.com/yeqingGhostCloud/beego_learn/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
