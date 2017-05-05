// @APIVersion 1.0.0
// @Title tripleS Home Automation API

package routers

import (
	"apis/controllers"
	"github.com/astaxie/beego"
)

func init() {
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
		beego.NSNamespace("/apis/timer",
			beego.NSInclude(
				&controllers.TimerController{},
			),
		),
		beego.NSNamespace("/apis/power",
			beego.NSInclude(
				&controllers.PowerController{},
			),
		),
	)

	beego.AddNamespace(ns1)

}
