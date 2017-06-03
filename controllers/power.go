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
// @router /getDevice [post]
func (this *PowerController) PostGetDevice() {
	ret_val:= models.Response{Status:true,Rcode:200}
	this.Data["json"]=&ret_val
	ret_val.Message="Success"

	uid:= GetUidByToken(this.Ctx.Request.Header.Get("auth"))

	// Parse input data
	input_form:=models.PowerInput{}
	if err:= this.ParseForm(&input_form);err!=nil {
		ret_val.Rcode=100
		ret_val.Message=err.Error()
		this.ServeJSON()
		return
	}

	hid:= input_form.Hid

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
// @router /getPowerDeviceLimit [post]
func (this *PowerController) PostGetPowerDeviceLimit() {
	ret_val:= models.Response{Status:true,Rcode:200}
	this.Data["json"]=&ret_val
	ret_val.Message="Success"

	uid:= GetUidByToken(this.Ctx.Request.Header.Get("auth"))

	input_form:=models.PowerInput{}
	if err:= this.ParseForm(&input_form);err!=nil {
		ret_val.Rcode=100
		ret_val.Message=err.Error()
		this.ServeJSON()
		return
	}

	date_limit:=input_form.Date_Limit
	date_skip:=input_form.Date_Skip
	did:=input_form.Did

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
// @router /getPowerDeviceByTime [post]
func (this *PowerController) PostGetPowerDeviceByTime() {
	ret_val:= models.Response{Status:true,Rcode:200}
	this.Data["json"]=&ret_val
	ret_val.Message="Success"

	uid:= GetUidByToken(this.Ctx.Request.Header.Get("auth"))

	input_form:=models.PowerInput{}
	if err:= this.ParseForm(&input_form);err!=nil {
		ret_val.Rcode=100
		ret_val.Message=err.Error()
		this.ServeJSON()
		return
	}

	date_start:=input_form.Date_Start
	date_end:=input_form.Date_End
	did:=input_form.Did

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


// @Title get device
// @Summary get device
// @Param	auth	header	string	"token"
// @Param	hid	header	string	hid
// @Success 200
// @Failure 201
// @router /getHomePowerInYear [post]
func (this *PowerController) PostGetHomePowerInYear() {
	ret_val:= models.Response{Status:true,Rcode:200}
	this.Data["json"]=&ret_val
	ret_val.Message="Success"

	uid:= GetUidByToken(this.Ctx.Request.Header.Get("auth"))

	input_form:=models.PowerInput{}
	if err:= this.ParseForm(&input_form);err!=nil {
		ret_val.Rcode=100
		ret_val.Message=err.Error()
		this.ServeJSON()
		return
	}

	time:=input_form.Time
	hid:=input_form.Hid


	ldevice,err := models.GetHomePowerInYearBYMonth(hid,time)

	if err!=nil{
		ret_val.Rcode=201
		ret_val.Message=err.Error()
		this.ServeJSON()
		return
	}
	print(ldevice)

	data:= models.PowerOutput{}

	data.Uid=uid
	data.Total=len(ldevice)
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
// @router /getHomePowerInYearByMonth [post]
func (this *PowerController) PostGetHomePowerInYearByMonth() {
	ret_val:= models.Response{Status:true,Rcode:200}
	this.Data["json"]=&ret_val
	ret_val.Message="Success"

	uid:= GetUidByToken(this.Ctx.Request.Header.Get("auth"))
	input_form:=models.PowerInput{}
	if err:= this.ParseForm(&input_form);err!=nil {
		ret_val.Rcode=100
		ret_val.Message=err.Error()
		this.ServeJSON()
		return
	}

	hid:= input_form.Hid
	time:=input_form.Time


	ldevice,err := models.GetHomePowerInYearBYMonth(hid,time)

	if err!=nil{
		ret_val.Rcode=201
		ret_val.Message=err.Error()
		this.ServeJSON()
		return
	}
	print(ldevice)

	data:= models.PowerOutput{}

	data.Uid=uid
	data.Total=len(ldevice)
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
// @router /getHomePowerInMonthByDate [post]
func (this *PowerController) PostGetHomePowerInMonthByDate() {
	ret_val:= models.Response{Status:true,Rcode:200}
	this.Data["json"]=&ret_val
	ret_val.Message="Success"

	uid:= GetUidByToken(this.Ctx.Request.Header.Get("auth"))
	input_form:=models.PowerInput{}
	if err:= this.ParseForm(&input_form);err!=nil {
		ret_val.Rcode=100
		ret_val.Message=err.Error()
		this.ServeJSON()
		return
	}

	hid:= input_form.Hid
	time:=input_form.Time

	ldevice,err := models.GetHomePowerInMonthByDate(hid,time)

	if err!=nil{
		ret_val.Rcode=201
		ret_val.Message=err.Error()
		this.ServeJSON()
		return
	}
	print(ldevice)

	data:= models.PowerOutput{}

	data.Uid=uid
	data.Total=len(ldevice)
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
// @router /getDevicePowerInMonthByDate [post]
func (this *PowerController) PostGetDevicePowerInMonthByDate() {
	ret_val:= models.Response{Status:true,Rcode:200}
	this.Data["json"]=&ret_val
	ret_val.Message="Success"

	uid:= GetUidByToken(this.Ctx.Request.Header.Get("auth"))

	input_form:=models.PowerInput{}
	if err:= this.ParseForm(&input_form);err!=nil {
		ret_val.Rcode=100
		ret_val.Message=err.Error()
		this.ServeJSON()
		return
	}

	time:=input_form.Time
	did:=input_form.Did

	ldevice,err := models.GetDevicePowerInMonthByDate(did,time)

	if err!=nil{
		ret_val.Rcode=201
		ret_val.Message=err.Error()
		this.ServeJSON()
		return
	}
	print(ldevice)

	data:= models.PowerOutput{}

	data.Uid=uid
	data.Total=len(ldevice)
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
// @router /getDevicePowerInYearByMonth [post]
func (this *PowerController) PostGetDevicePowerInYearByMonth() {
	ret_val:= models.Response{Status:true,Rcode:200}
	this.Data["json"]=&ret_val
	ret_val.Message="Success"

	uid:= GetUidByToken(this.Ctx.Request.Header.Get("auth"))

	input_form:=models.PowerInput{}
	if err:= this.ParseForm(&input_form);err!=nil {
		ret_val.Rcode=100
		ret_val.Message=err.Error()
		this.ServeJSON()
		return
	}

	time:=input_form.Time
	did:=input_form.Did

	ldevice,err := models.GetDevicePowerInYearBYMonth(did,time)

	if err!=nil{
		ret_val.Rcode=201
		ret_val.Message=err.Error()
		this.ServeJSON()
		return
	}
	print(ldevice)

	data:= models.PowerOutput{}

	data.Uid=uid
	data.Total=len(ldevice)
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
// @router /getRankingDevicePowerInMonth [post]
func (this *PowerController) PostGetRankingDevicePowerInMonth() {
	ret_val:= models.Response{Status:true,Rcode:200}
	this.Data["json"]=&ret_val
	ret_val.Message="Success"

	//uid:= GetUidByToken(this.Ctx.Request.Header.Get("auth"))

	input_form:=models.PowerInput{}
	if err:= this.ParseForm(&input_form);err!=nil {
		ret_val.Rcode=100
		ret_val.Message=err.Error()
		this.ServeJSON()
		return
	}

	time:=input_form.Time
	hid:=input_form.Hid

	ldevice,err := models.GetRankingDevicePowerInMonth(hid,time)

	if err!=nil{
		ret_val.Rcode=201
		ret_val.Message=err.Error()
		this.ServeJSON()
		return
	}
	print(ldevice)

	data:= models.GetRankingDevicePowerOutput{}

	data.Hid=hid
	data.Total=len(ldevice)
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
// @router /getRankingDevicePowerInYear [post]
func (this *PowerController) PostGetRankingDevicePowerInYear() {
	ret_val:= models.Response{Status:true,Rcode:200}
	this.Data["json"]=&ret_val
	ret_val.Message="Success"

	//uid:= GetUidByToken(this.Ctx.Request.Header.Get("auth"))


	input_form:=models.PowerInput{}
	if err:= this.ParseForm(&input_form);err!=nil {
		ret_val.Rcode=100
		ret_val.Message=err.Error()
		this.ServeJSON()
		return
	}

	time:=input_form.Time
	hid:=input_form.Hid

	ldevice,err := models.GetRankingDevicePowerInYear(hid,time)

	if err!=nil{
		ret_val.Rcode=201
		ret_val.Message=err.Error()
		this.ServeJSON()
		return
	}
	print(ldevice)

	data:= models.GetRankingDevicePowerOutput{}

	data.Hid=hid
	data.Total=len(ldevice)
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
// @router /getRankingDevicePowerAll [post]
func (this *PowerController) PostGetRankingDevicePowerAll() {
	ret_val:= models.Response{Status:true,Rcode:200}
	this.Data["json"]=&ret_val
	ret_val.Message="Success"

	//uid:= GetUidByToken(this.Ctx.Request.Header.Get("auth"))

	input_form:=models.PowerInput{}
	if err:= this.ParseForm(&input_form);err!=nil {
		ret_val.Rcode=100
		ret_val.Message=err.Error()
		this.ServeJSON()
		return
	}

	hid:=input_form.Hid

	ldevice,err := models.GetRankingDevicePowerAll(hid)

	if err!=nil{
		ret_val.Rcode=201
		ret_val.Message=err.Error()
		this.ServeJSON()
		return
	}
	print(ldevice)

	data:= models.GetRankingDevicePowerOutput{}

	data.Hid=hid
	data.Total=len(ldevice)
	data.Devices=ldevice

	ret_val.Data=data

	this.ServeJSON()
	return

}