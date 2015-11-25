package routers

import (
	"github.com/lufeipeng/lufeipeng/DefenceInfoEditor/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{});
    beego.Router("/view", &controllers.ViewController{});
}
