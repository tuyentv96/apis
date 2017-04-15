// @APIVersion 1.0.0
// @Title tripleS Home Automation API

package routers

import (
	"apis/controllers"
	"github.com/astaxie/beego"
)

func init() {
	/*
    beego.Router("/", &controllers.MainController{})

	beego.Router("apis/history", &controllers.HistoryController{})
	beego.Router("apis/mgetdevice", &controllers.MGetDeviceController{})

	beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/apis",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
*/
	//beego.Router("/auth", &controllers.AuthController{})
	ns1 := beego.NewNamespace("/v1/",
		beego.NSNamespace("/auth",
			beego.NSInclude(
				&controllers.AuthController{},
			),
		),
		beego.NSBefore(controllers.Auth),
		beego.NSNamespace("/apis/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/apis/admin",
			beego.NSInclude(
				&controllers.AdminController{},
			),
		),
	)

	beego.AddNamespace(ns1)

}
