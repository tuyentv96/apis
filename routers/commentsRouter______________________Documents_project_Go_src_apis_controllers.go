package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["apis/controllers:AuthController"] = append(beego.GlobalControllerRouter["apis/controllers:AuthController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apis/controllers:AuthController"] = append(beego.GlobalControllerRouter["apis/controllers:AuthController"],
		beego.ControllerComments{
			Method: "PostTokenVerify",
			Router: `/tokenVerify`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apis/controllers:MGetDeviceController"] = append(beego.GlobalControllerRouter["apis/controllers:MGetDeviceController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/getAllDevice`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
