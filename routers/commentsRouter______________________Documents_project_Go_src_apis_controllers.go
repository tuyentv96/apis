package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["apis/controllers:AdminController"] = append(beego.GlobalControllerRouter["apis/controllers:AdminController"],
		beego.ControllerComments{
			Method: "PostAddDevice",
			Router: `/addDevice`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

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

	beego.GlobalControllerRouter["apis/controllers:AuthController"] = append(beego.GlobalControllerRouter["apis/controllers:AuthController"],
		beego.ControllerComments{
			Method: "PostGetCode",
			Router: `/forgetPassword/getCode`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apis/controllers:AuthController"] = append(beego.GlobalControllerRouter["apis/controllers:AuthController"],
		beego.ControllerComments{
			Method: "PostVerifyCode",
			Router: `/forgetPassword/verifyCode`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apis/controllers:AuthController"] = append(beego.GlobalControllerRouter["apis/controllers:AuthController"],
		beego.ControllerComments{
			Method: "PostLogout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apis/controllers:PowerController"] = append(beego.GlobalControllerRouter["apis/controllers:PowerController"],
		beego.ControllerComments{
			Method: "GetDevice",
			Router: `/getDevice`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apis/controllers:PowerController"] = append(beego.GlobalControllerRouter["apis/controllers:PowerController"],
		beego.ControllerComments{
			Method: "GetPowerDeviceLimit",
			Router: `/getPowerDeviceLimit`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apis/controllers:PowerController"] = append(beego.GlobalControllerRouter["apis/controllers:PowerController"],
		beego.ControllerComments{
			Method: "GetPowerDeviceByTime",
			Router: `/getPowerDeviceByTime`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apis/controllers:PowerController"] = append(beego.GlobalControllerRouter["apis/controllers:PowerController"],
		beego.ControllerComments{
			Method: "GetHomePowerInYear",
			Router: `/getHomePowerInYear`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apis/controllers:PowerController"] = append(beego.GlobalControllerRouter["apis/controllers:PowerController"],
		beego.ControllerComments{
			Method: "GetHomePowerInYearBy",
			Router: `/getHomePowerInYear`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apis/controllers:PowerController"] = append(beego.GlobalControllerRouter["apis/controllers:PowerController"],
		beego.ControllerComments{
			Method: "GetHomePowerInMonthByDate",
			Router: `/getHomePowerInMonthByDate`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apis/controllers:PowerController"] = append(beego.GlobalControllerRouter["apis/controllers:PowerController"],
		beego.ControllerComments{
			Method: "GetDevicePowerInMonthByDate",
			Router: `/getDevicePowerInMonthByDate`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apis/controllers:PowerController"] = append(beego.GlobalControllerRouter["apis/controllers:PowerController"],
		beego.ControllerComments{
			Method: "GetDevicePowerInYearByMonth",
			Router: `/getDevicePowerInYearByMonth`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apis/controllers:PowerController"] = append(beego.GlobalControllerRouter["apis/controllers:PowerController"],
		beego.ControllerComments{
			Method: "GetRankingDevicePowerInMonth",
			Router: `/getRankingDevicePowerInMonth`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apis/controllers:PowerController"] = append(beego.GlobalControllerRouter["apis/controllers:PowerController"],
		beego.ControllerComments{
			Method: "GetRankingDevicePowerInYear",
			Router: `/getRankingDevicePowerInYear`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apis/controllers:PowerController"] = append(beego.GlobalControllerRouter["apis/controllers:PowerController"],
		beego.ControllerComments{
			Method: "GetRankingDevicePowerAll",
			Router: `/getRankingDevicePowerAll`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apis/controllers:TimerController"] = append(beego.GlobalControllerRouter["apis/controllers:TimerController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/setTimer`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apis/controllers:TimerController"] = append(beego.GlobalControllerRouter["apis/controllers:TimerController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/getTimerByUid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apis/controllers:TimerController"] = append(beego.GlobalControllerRouter["apis/controllers:TimerController"],
		beego.ControllerComments{
			Method: "GetTimer",
			Router: `/getTimerByDid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apis/controllers:TimerController"] = append(beego.GlobalControllerRouter["apis/controllers:TimerController"],
		beego.ControllerComments{
			Method: "PostDeleteTimer",
			Router: `/deleteTimer`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apis/controllers:TimerController"] = append(beego.GlobalControllerRouter["apis/controllers:TimerController"],
		beego.ControllerComments{
			Method: "PostCreateCronTimer",
			Router: `/setCronTimer`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apis/controllers:TimerController"] = append(beego.GlobalControllerRouter["apis/controllers:TimerController"],
		beego.ControllerComments{
			Method: "PostDeleteCronTimer",
			Router: `/deleteCronTimer`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apis/controllers:TimerController"] = append(beego.GlobalControllerRouter["apis/controllers:TimerController"],
		beego.ControllerComments{
			Method: "GetCron",
			Router: `/getCronTimer`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apis/controllers:TimerController"] = append(beego.GlobalControllerRouter["apis/controllers:TimerController"],
		beego.ControllerComments{
			Method: "PostStopCronTimer",
			Router: `/stopCronTimer`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apis/controllers:TimerController"] = append(beego.GlobalControllerRouter["apis/controllers:TimerController"],
		beego.ControllerComments{
			Method: "PostStartCronTimer",
			Router: `/startCronTimer`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apis/controllers:UserController"] = append(beego.GlobalControllerRouter["apis/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAllHome",
			Router: `/getAllHome`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apis/controllers:UserController"] = append(beego.GlobalControllerRouter["apis/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAllDevice",
			Router: `/getAllDevice`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apis/controllers:UserController"] = append(beego.GlobalControllerRouter["apis/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAllDeviceByHid",
			Router: `/getAllDeviceByHid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apis/controllers:UserController"] = append(beego.GlobalControllerRouter["apis/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetHisotory",
			Router: `/getHistoryHome`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apis/controllers:UserController"] = append(beego.GlobalControllerRouter["apis/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetHisotoryByTime",
			Router: `/getHistoryHomeByTime`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apis/controllers:UserController"] = append(beego.GlobalControllerRouter["apis/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetHistoryByDid",
			Router: `/getHistoryDevice`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apis/controllers:UserController"] = append(beego.GlobalControllerRouter["apis/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetHistoryByDidByTime",
			Router: `/getHistoryDeviceByTime`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apis/controllers:UserController"] = append(beego.GlobalControllerRouter["apis/controllers:UserController"],
		beego.ControllerComments{
			Method: "PostUpdatePassword",
			Router: `/updatePassword`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apis/controllers:UserController"] = append(beego.GlobalControllerRouter["apis/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetProfile",
			Router: `/getProfile`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apis/controllers:UserController"] = append(beego.GlobalControllerRouter["apis/controllers:UserController"],
		beego.ControllerComments{
			Method: "PostUpdateProfile",
			Router: `/updateProfile`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

}
