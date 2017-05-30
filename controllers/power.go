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

	date_start_str:= this.Ctx.Request.Header.Get("datestart")
	date_end_str:= this.Ctx.Request.Header.Get("dateend")

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


// @Title get device
// @Summary get device
// @Param	auth	header	string	"token"
// @Param	hid	header	string	hid
// @Success 200
// @Failure 201
// @router /getHomePowerInYear [get]
func (this *PowerController) GetHomePowerInYear() {
	ret_val:= models.Response{Status:true,Rcode:200}
	this.Data["json"]=&ret_val
	ret_val.Message="Success"

	uid:= GetUidByToken(this.Ctx.Request.Header.Get("auth"))
	hid:= this.Ctx.Request.Header.Get("hid")

	time_str:= this.Ctx.Request.Header.Get("time")
	//date_end_str:= this.Ctx.Request.Header.Get("dateend")

	time_in_year,_:= strconv.ParseInt(time_str,0,64)
	//date_end,_:= strconv.ParseInt(date_end_str,0,64)


	ldevice,err := models.GetHomePowerInYearBYMonth(hid,time_in_year)

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
// @router /getHomePowerInYear [get]
func (this *PowerController) GetHomePowerInYearBy() {
	ret_val:= models.Response{Status:true,Rcode:200}
	this.Data["json"]=&ret_val
	ret_val.Message="Success"

	uid:= GetUidByToken(this.Ctx.Request.Header.Get("auth"))
	hid:= this.Ctx.Request.Header.Get("hid")

	time_str:= this.Ctx.Request.Header.Get("time")
	//date_end_str:= this.Ctx.Request.Header.Get("dateend")

	time_in_year,_:= strconv.ParseInt(time_str,0,64)
	//date_end,_:= strconv.ParseInt(date_end_str,0,64)


	ldevice,err := models.GetHomePowerInYearBYMonth(hid,time_in_year)

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
// @router /getHomePowerInMonthByDate [get]
func (this *PowerController) GetHomePowerInMonthByDate() {
	ret_val:= models.Response{Status:true,Rcode:200}
	this.Data["json"]=&ret_val
	ret_val.Message="Success"

	uid:= GetUidByToken(this.Ctx.Request.Header.Get("auth"))
	hid:= this.Ctx.Request.Header.Get("hid")

	time_str:= this.Ctx.Request.Header.Get("time")
	//date_end_str:= this.Ctx.Request.Header.Get("dateend")

	time_in_year,_:= strconv.ParseInt(time_str,0,64)
	//date_end,_:= strconv.ParseInt(date_end_str,0,64)


	ldevice,err := models.GetHomePowerInMonthByDate(hid,time_in_year)

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
// @router /getDevicePowerInMonthByDate [get]
func (this *PowerController) GetDevicePowerInMonthByDate() {
	ret_val:= models.Response{Status:true,Rcode:200}
	this.Data["json"]=&ret_val
	ret_val.Message="Success"

	uid:= GetUidByToken(this.Ctx.Request.Header.Get("auth"))
	did:= this.Ctx.Request.Header.Get("did")

	time_str:= this.Ctx.Request.Header.Get("time")
	//date_end_str:= this.Ctx.Request.Header.Get("dateend")

	time_in_year,_:= strconv.ParseInt(time_str,0,64)
	//date_end,_:= strconv.ParseInt(date_end_str,0,64)

	println(did,time_in_year)

	ldevice,err := models.GetDevicePowerInMonthByDate(did,time_in_year)

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
// @router /getDevicePowerInYearByMonth [get]
func (this *PowerController) GetDevicePowerInYearByMonth() {
	ret_val:= models.Response{Status:true,Rcode:200}
	this.Data["json"]=&ret_val
	ret_val.Message="Success"

	uid:= GetUidByToken(this.Ctx.Request.Header.Get("auth"))
	did:= this.Ctx.Request.Header.Get("did")

	time_str:= this.Ctx.Request.Header.Get("time")
	//date_end_str:= this.Ctx.Request.Header.Get("dateend")

	time_in_year,_:= strconv.ParseInt(time_str,0,64)
	//date_end,_:= strconv.ParseInt(date_end_str,0,64)

	println(did,time_in_year)

	ldevice,err := models.GetDevicePowerInYearBYMonth(did,time_in_year)

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
// @router /getRankingDevicePowerInMonth [get]
func (this *PowerController) GetRankingDevicePowerInMonth() {
	ret_val:= models.Response{Status:true,Rcode:200}
	this.Data["json"]=&ret_val
	ret_val.Message="Success"

	//uid:= GetUidByToken(this.Ctx.Request.Header.Get("auth"))

	time_str:= this.Ctx.Request.Header.Get("time")
	hid:= this.Ctx.Request.Header.Get("hid")

	time_in_year,_:= strconv.ParseInt(time_str,0,64)

	println(time_in_year)

	ldevice,err := models.GetRankingDevicePowerInMonth(hid,time_in_year)

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
// @router /getRankingDevicePowerInYear [get]
func (this *PowerController) GetRankingDevicePowerInYear() {
	ret_val:= models.Response{Status:true,Rcode:200}
	this.Data["json"]=&ret_val
	ret_val.Message="Success"

	//uid:= GetUidByToken(this.Ctx.Request.Header.Get("auth"))

	time_str:= this.Ctx.Request.Header.Get("time")
	hid:= this.Ctx.Request.Header.Get("hid")

	time_in_year,_:= strconv.ParseInt(time_str,0,64)

	println(time_in_year)

	ldevice,err := models.GetRankingDevicePowerInYear(hid,time_in_year)

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
// @router /getRankingDevicePowerAll [get]
func (this *PowerController) GetRankingDevicePowerAll() {
	ret_val:= models.Response{Status:true,Rcode:200}
	this.Data["json"]=&ret_val
	ret_val.Message="Success"

	//uid:= GetUidByToken(this.Ctx.Request.Header.Get("auth"))

	//time_str:= this.Ctx.Request.Header.Get("time")
	hid:= this.Ctx.Request.Header.Get("hid")

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