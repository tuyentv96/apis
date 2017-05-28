package controllers

import (
	"github.com/astaxie/beego"
	models "apis/models"
	service "apis/services"
	"fmt"
	"encoding/json"
)

// Operations about user
type TimerController struct {
	beego.Controller
}

// @Title setTimerByDate
// @Summary setTimerByDate
// @Param	auth	header	string	"token"
// @Param	timer_name	formData	string timer_name
// @Param	devices	formData	string devices
// @Param	status	formData	int status
// @Param	time	formData	int time
// @Success 200
// @Failure 201
// @router /setTimer [post]
func (this *TimerController) Post() {
	ret_val:= models.Response{Status:true,Rcode:200}
	this.Data["json"]=&ret_val

	uid:= GetUidByToken(this.Ctx.Request.Header.Get("auth"))

	set_timer_input:=models.SetTimerInput{}
	if err:= this.ParseForm(&set_timer_input);err!=nil {
		ret_val.Rcode=100
		ret_val.Message="Wrong format"
		this.ServeJSON()
		return

	}
	set_timer_input.Uid=uid

	if set_timer_input.Timer_Name=="" || set_timer_input.Uid==""{
		ret_val.Rcode=103
		ret_val.Message="Please fill in all fields"
		this.ServeJSON()
		return
	}

	if existed:= models.CheckTimerExisted(set_timer_input.Uid,set_timer_input.Timer_Name,int(set_timer_input.Time));existed==true{
		ret_val.Rcode=312
		ret_val.Message="Timer existed"
		this.ServeJSON()
		return
	}

	json.Unmarshal([]byte(this.GetString("devices")),&set_timer_input.Devices)
	fmt.Printf("%+v",set_timer_input,this.GetString("devices"))



	list_dev:= models.FindListDeviceByID(set_timer_input.Devices)
	ret:=models.SaveTimerDevice(list_dev,set_timer_input.Timer_Name,"H1",set_timer_input.Uid,set_timer_input.Time)
	service.DoStuff(list_dev,set_timer_input.Hid,set_timer_input.Uid,set_timer_input.Time,ret.TimerId)

	ret_val.Data=ret
	this.ServeJSON()
	return

}

// @Title Get timer by uid
// @Summary Get timer by uid
// @Param	auth	header	string	"token"
// @Success 200
// @Failure 201 body is empty
// @router /getTimerByUid [get]
func (this *TimerController) Get() {
	ret_val:= models.Response{Status:true,Rcode:200}
	this.Data["json"]=&ret_val

	uid:= GetUidByToken(this.Ctx.Request.Header.Get("auth"))

	/*
	did:= this.Ctx.Request.Header.Get("did")
	skip_str:= this.Ctx.Request.Header.Get("skip")
	limit_str:= this.Ctx.Request.Header.Get("limit")

	limit,_:= strconv.ParseInt(limit_str,0,64)
	skip,_:= strconv.ParseInt(skip_str, 0, 64)
	*/

	result,err:=models.GetTimerDeviceByUid(uid)

	if err {
		ret_val.Rcode=100
		ret_val.Message="Get date timer fail"
		this.ServeJSON()
		return
	}

	ret_val.Data=result
	ret_val.Message="Get date timer success"
	this.ServeJSON()
	return

}

// @Title getTimerByDid
// @Summary getTimerByDid
// @Param	auth	header	string	"token"
// @Param	did	header	string	device_id
// @Success 200
// @Failure 201 body is empty
// @router /getTimerByDid [get]
func (this *TimerController) GetTimer() {
	ret_val:= models.Response{Status:true,Rcode:200}
	this.Data["json"]=&ret_val

	uid:= GetUidByToken(this.Ctx.Request.Header.Get("auth"))
	did:= this.Ctx.Request.Header.Get("did")

	/*

	skip_str:= this.Ctx.Request.Header.Get("skip")
	limit_str:= this.Ctx.Request.Header.Get("limit")

	limit,_:= strconv.ParseInt(limit_str,0,64)
	skip,_:= strconv.ParseInt(skip_str, 0, 64)
	*/

	result,err:=models.GetTimerDeviceByDid(uid,did)

	if err {
		ret_val.Rcode=100
		ret_val.Message="Get date timer fail"
		this.ServeJSON()
		return
	}

	ret_val.Data=result
	ret_val.Message="Get date timer success"
	this.ServeJSON()
	return

}


// @Title setTimer
// @Summary Set timer
// @Param	auth	header	string	"token"
// @Param	timer_id	header	string	timer_id
// @Success 200
// @Failure 201 body is empty
// @router /deleteTimer [post]
func (this *TimerController) PostDeleteTimer() {
	ret_val:= models.Response{Status:true,Rcode:200}
	this.Data["json"]=&ret_val

	timer_id:= this.Ctx.Request.Header.Get("timer_id")

	if err:=models.DeleteTimerDevice(timer_id);err{
		print("No timer id")
		ret_val.Rcode=100
		ret_val.Message="No timer_id"
		this.ServeJSON()
		return
	}

	ret_val.Data=timer_id
	service.RemoveChanByTimerId(timer_id)
	this.ServeJSON()
	return

}

// @Title setTimerByDate
// @Summary setCronTimer
// @Param	auth	header	string	"token"
// @Param	devices	formData	string devices
// @Param	status	formData	int status
// @Param	time	formData	string time
// @Param	date	formData	string time
// @Success 200
// @Failure 201
// @router /setCronTimer [post]
func (this *TimerController) PostCreateCronTimer() {
	ret_val:= models.Response{Status:true,Rcode:200}
	this.Data["json"]=&ret_val

	uid:= GetUidByToken(this.Ctx.Request.Header.Get("auth"))
	set_timer_input:=models.CronTimerInput{}

	if err:= this.ParseForm(&set_timer_input);err!=nil {
		this.Ctx.ResponseWriter.WriteHeader(100)
		ret_val.Message="Wrong format"
		ret_val.Status=false
		this.ServeJSON()
		return

	}

	set_timer_input.Uid=uid

	if set_timer_input.Cron_Name=="" || set_timer_input.Uid=="" || set_timer_input.Time==""{
		ret_val.Rcode=103
		ret_val.Message="Please fill in all fields"
		this.ServeJSON()
		return
	}

	if existed:= models.CheckCronExisted(set_timer_input.Uid,set_timer_input.Cron_Name,set_timer_input.Time); existed==true{
		ret_val.Rcode=312
		ret_val.Message="Cron existed"
		this.ServeJSON()
		return
	}

	json.Unmarshal([]byte(this.GetString("devices")),&set_timer_input.Devices)
	fmt.Printf("%+v",set_timer_input,this.GetString("devices"))

	list_dev:= models.FindListDeviceByID(set_timer_input.Devices)
	ret:=models.SaveCronDevice(list_dev,set_timer_input.Cron_Name,"H1",set_timer_input.Uid,set_timer_input.Time,set_timer_input.Date)
	service.CronDoStuff(list_dev,set_timer_input.Hid,set_timer_input.Uid,set_timer_input.Time,set_timer_input.Date,ret.TimerId)

	ret_val.Data=ret

	ret_val.Message="Create cron success"
	this.ServeJSON()
	return

}

// @Title setTimer
// @Summary Set timer
// @Param	auth	header	string	"token"
// @Param	timer_id	header	string	timer_id
// @Success 200
// @Failure 201 body is empty
// @router /deleteCronTimer [post]
func (this *TimerController) PostDeleteCronTimer() {
	ret_val:= models.Response{Status:true,Rcode:200}
	this.Data["json"]=&ret_val

	uid:= GetUidByToken(this.Ctx.Request.Header.Get("auth"))
	timer_id:= this.GetString("timer_id")

	if execute,err:= models.CheckTimerExcuteByTimerID(timer_id); err!=nil {
		print("No timer id")
		ret_val.Rcode=100
		ret_val.Message=("No timer id")
		this.ServeJSON()
		return
	} else {
		if execute==false {
			ret_val.Rcode=101
			ret_val.Message=("Timer stopped")
			this.ServeJSON()
			return
		}
	}

	if err:=models.DeleteCronTimer(uid,timer_id);err{
		print("No timer id")
		ret_val.Rcode=100
		ret_val.Message="Delete cron fail"
		this.ServeJSON()
		return
	}
	service.RemoveCronChanByTimerId(timer_id)

	ret_val.Message="Delete cron success"
	this.ServeJSON()
	return

}

// @Title getTimer
// @Summary Get timer
// @Param	auth	header	string	"token"
// @Param	hid	header	string	hid
// @Success 200
// @Failure 201 body is empty
// @router /getCronTimer [get]
func (this *TimerController) GetCron() {
	ret_val:= models.Response{Status:true,Rcode:200}
	this.Data["json"]=&ret_val

	uid:= GetUidByToken(this.Ctx.Request.Header.Get("auth"))
	hid:= this.Ctx.Request.Header.Get("hid")

	result,err:=models.GetCronByHid(uid,hid)

	if err {
		ret_val.Rcode=100
		this.ServeJSON()
		return
	}

	ret_val.Data=result
	this.ServeJSON()
	return

}

// @Title stop Cron
// @Summary stop Cron
// @Param	auth	header	string	"token"
// @Param	timer_id	header	string	timer_id
// @Success 200
// @Failure 201 body is empty
// @router /stopCronTimer [post]
func (this *TimerController) PostStopCronTimer() {
	ret_val:= models.Response{Status:true,Rcode:200}
	this.Data["json"]=&ret_val

	timer_id:= this.GetString("timer_id")

	if enable,err:= models.CheckCronEnableByTimerID(timer_id); err!=nil {
		print("No timer id")
		ret_val.Rcode=100
		ret_val.Message=("No timer id")
		this.ServeJSON()
		return
	} else {
		if enable==false {
			ret_val.Rcode=101
			ret_val.Message=("Cron stopped")
			this.ServeJSON()
			return
		}
	}

	if err:=models.StopCronTimer(timer_id);err!=nil{
		print("No timer id")
		ret_val.Rcode=100
		ret_val.Message=("No timer id")
		this.ServeJSON()
		return
	}
	service.RemoveCronChanByTimerId(timer_id)
	ret_val.Data=timer_id
	ret_val.Message="Stop cron success"
	this.ServeJSON()
	return

}

// @Title setTimer
// @Summary Set timer
// @Param	auth	header	string	"token"
// @Param	timer_id	header	string	timer_id
// @Success 200
// @Failure 201 body is empty
// @router /startCronTimer [post]
func (this *TimerController) PostStartCronTimer() {
	ret_val:= models.Response{Status:true,Rcode:200}
	this.Data["json"]=&ret_val

	timer_id:= this.GetString("timer_id")

	if enable,err:= models.CheckCronEnableByTimerID(timer_id); err!=nil {
		print("No timer id")
		ret_val.Rcode=100
		ret_val.Message=("No timer id")
		this.ServeJSON()
		return
	} else {
		if enable==true{
			ret_val.Rcode=101
			ret_val.Message=("Cron started")
			this.ServeJSON()
			return
		}
	}

	data,err:=models.FindCronByTimerID(timer_id)
	if err!=nil{
		print("No timer id")
		ret_val.Rcode=100
		ret_val.Message="No timer_id"
		this.ServeJSON()
		return
	}

	models.StartCronTimer(timer_id)
	service.CronDoStuff(data.Device,data.Hid,data.Uid,data.Time,data.Date,data.TimerId)
	ret_val.Data=timer_id
	ret_val.Message="Start cron success"
	this.ServeJSON()
	return

}
