package controllers

import (
	"github.com/astaxie/beego"
	models "apis/models"
	"strconv"
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

// @Title get device
// @Summary get device
// @Param	auth	header	string	"token"
// @Param	hid	header	string	hid
// @Success 200
// @Failure 201
// @router /getPowerDeviceLimit [get]
func (this *PowerController) GetPowerDeviceLimit() {
	ret_val:= models.Response{Status:true,Rcode:200}
	this.Data["json"]=&ret_val
	ret_val.Message="Success"

	uid:= GetUidByToken(this.Ctx.Request.Header.Get("auth"))
	did:= this.Ctx.Request.Header.Get("did")
	limit_str:= this.Ctx.Request.Header.Get("date_limit")
	skip_str:= this.Ctx.Request.Header.Get("date_skip")

	date_limit,_:= strconv.ParseInt(limit_str,0,64)
	date_skip,_:= strconv.ParseInt(skip_str,0,64)
	//print(did,uid,limit)

	dev,err1:= models.FindDeviceByID(did)
	if err1!=""{
		ret_val.Rcode=202
		ret_val.Message="No device found"
		this.ServeJSON()
		return
	}

	ldevice,err := models.GetDevicePowerLimit(did,date_skip,date_limit)

	if err!=nil{
		ret_val.Rcode=201
		ret_val.Message=err.Error()
		this.ServeJSON()
		return
	}
	print(ldevice)

	data:= models.GetDevicePowerOneOutput{}

	data.Uid=uid
	data.Hid=dev.Hid
	data.Did=dev.Did
	data.Dname=dev.Dname
	data.Devices=ldevice
	data.Type=dev.Type
	data.Status=dev.Status
	data.Total=len(ldevice)

	ret_val.Data=data

	this.ServeJSON()
	return

}

// @Title get device
// @Summary get device
// @Param	auth	header	string	"token"
// @Param	hid	header	string	hid
// @Success 200
// @Failure 201
// @router /getPowerDeviceByTime [get]
func (this *PowerController) GetPowerDeviceByTime() {
	ret_val:= models.Response{Status:true,Rcode:200}
	this.Data["json"]=&ret_val
	ret_val.Message="Success"

	uid:= GetUidByToken(this.Ctx.Request.Header.Get("auth"))
	did:= this.Ctx.Request.Header.Get("did")

	date_start_str:= this.Ctx.Request.Header.Get("date_start")
	date_end_str:= this.Ctx.Request.Header.Get("date_end")

	date_start,_:= strconv.ParseInt(date_start_str,0,64)
	date_end,_:= strconv.ParseInt(date_end_str,0,64)

	dev,err1:= models.FindDeviceByID(did)
	if err1!=""{
		ret_val.Rcode=202
		ret_val.Message="No device found"
		this.ServeJSON()
		return
	}


	ldevice,err := models.GetDevicePowerByTime(did,date_start,date_end)

	if err!=nil{
		ret_val.Rcode=201
		ret_val.Message=err.Error()
		this.ServeJSON()
		return
	}
	print(ldevice)

	data:= models.GetDevicePowerOneOutput{}

	data.Uid=uid
	data.Hid=dev.Hid
	data.Did=dev.Did
	data.Dname=dev.Dname
	data.Devices=ldevice
	data.Type=dev.Type
	data.Status=dev.Status
	data.Total=len(ldevice)

	ret_val.Data=data

	this.ServeJSON()
	return

}