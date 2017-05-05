package controllers

import (
	"github.com/astaxie/beego"
	models "apis/models"
)

// Operations about user
type PowerController struct {
	beego.Controller
}

// @Title get device
// @Summary get device
// @Param	auth	header	string	"token"
// @Param	hid	header	string	hid
// @Success 200
// @Failure 201
// @router /getDevice [get]
func (this *PowerController) GetDevice() {
	ret_val:= models.Response{Status:true,Rcode:200}
	this.Data["json"]=&ret_val
	ret_val.Message="Success"

	uid:= GetUidByToken(this.Ctx.Request.Header.Get("auth"))
	hid:= this.Ctx.Request.Header.Get("hid")
	print(hid,uid)

	ldevice,err := models.GetDevicePower(uid,hid)

	if err!=nil{
		ret_val.Rcode=201
		ret_val.Message=err.Error()
		this.ServeJSON()
		return
	}
	print(ldevice)

	data:= models.GetDevicePowerOutput{}

	data.Uid=uid
	data.Hid=hid
	data.Devices=ldevice

	ret_val.Data=data

	this.ServeJSON()
	return

}