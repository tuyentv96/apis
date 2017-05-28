package controllers

import (
	"github.com/astaxie/beego"
	models "apis/models"
	"strconv"
)

// Operations about user
type UserController struct {
	beego.Controller
}

// @Title UserGetAllDevice
// @Summary User get all device
// @Param	auth	header	string	"token"
// @Success 200 {object} models.LDevice
// @Failure 201 body is empty
// @router /getAllHome [get]
func (this *UserController) GetAllHome() {
	ret_val:= models.Response{Status:true,Rcode:200,Message:"Get home success"}
	this.Data["json"]=&ret_val
	uid:= GetUidByToken(this.Ctx.Request.Header.Get("auth"))

	if uid=="" {
		ret_val.Rcode=400
		ret_val.Message="User not found"
		ret_val.Status=false
		this.ServeJSON()
		return
	}

	ret_val.Data=models.MGetHome(uid)

	this.ServeJSON()
	return

}



// @Title UserGetAllDevice
// @Summary User get all device
// @Param	auth	header	string	"token"
// @Success 200 {object} models.LDevice
// @Failure 201 body is empty
// @router /getAllDevice [get]
func (u *UserController) GetAllDevice() {
	uid:= GetUidByToken(u.Ctx.Request.Header.Get("auth"))

	if uid=="" {
		u.Data["json"]=models.Err(400)
		u.Ctx.ResponseWriter.WriteHeader(400)
		return
	}

	u.Data["json"]=models.GetAllDeviceRsp{Rcode:200,Data:models.MGetDevice(uid),Message:"Success",Status:true}

	u.ServeJSON()
	return

}

// @Title UserGetAllDevice by hid
// @Summary User get all device by hid
// @Param	auth	header	string	"token"
// @Param	hid	header	string	"hid"
// @Success 200 {object} models.LDevice
// @Failure 201 body is empty
// @router /getAllDeviceByHid [get]
func (this *UserController) GetAllDeviceByHid() {
	ret_val:= models.Response{Status:true,Rcode:200,Message:"Get home success"}
	this.Data["json"]=&ret_val


	uid:= GetUidByToken(this.Ctx.Request.Header.Get("auth"))
	hid:= this.Ctx.Request.Header.Get("hid")

	println("Hid:",hid)
	if uid=="" {
		ret_val.Rcode=400
		ret_val.Message="User not found"
		ret_val.Status=false
		this.ServeJSON()
		return
	}

	data:=models.MGetDeviceByHid(uid,hid)

	ret_val.Data=models.MGetDeviceByHidData{Total:len(data),Devices:data}

	this.ServeJSON()
	return

}

// @Title UserGetHistory
// @Summary User get history
// @Param	auth	header	string	"The token key"
// @Param	hid	header	string	"home_id"
// @Param	limit	header	int32	"count of record"
// @Param	skip	header	int32	"skip"
// @Success 200 {object} models.HistoryInfo
// @Failure 201 fail
// @Failure 401 bad request
// @router /getHistoryHome [get]
func (u *UserController) GetHisotory() {
	hid:= u.Ctx.Request.Header.Get("hid")
	limit_str:= u.Ctx.Request.Header.Get("limit")
	skip_str:= u.Ctx.Request.Header.Get("skip")

	limit,_:= strconv.ParseInt(limit_str,0,64)
	skip,_:= strconv.ParseInt(skip_str, 0, 64)

	print(hid,"-",limit,"-",skip)

	ldevice,code,err:= models.MGetHistoryHome(hid,int(skip),int(limit))
	if err {
		u.Data["json"]=models.Err(code)
		u.Ctx.ResponseWriter.WriteHeader(code)
		u.ServeJSON()
		return
	}

	print("Len:",len(ldevice))
	u.Data["json"]=models.GetHistoryRsp{Rcode:200,Data:models.HistoryInfo{Total:len(ldevice),Ldevice:ldevice},Message:"Success",Status:true}
	u.ServeJSON()
	return

}

// @Title UserGetHistory
// @Summary User get history
// @Param	auth	header	string	"The token key"
// @Param	hid	header	string	"home_id"
// @Param	limit	header	int32	"count of record"
// @Param	skip	header	int32	"skip"
// @Success 200 {object} models.HistoryInfo
// @Failure 201 fail
// @Failure 401 bad request
// @router /getHistoryHomeByTime [get]
func (u *UserController) GetHisotoryByTime() {
	hid:= u.Ctx.Request.Header.Get("hid")
	time_start_str:= u.Ctx.Request.Header.Get("time_start")
	time_end_str:= u.Ctx.Request.Header.Get("time_end")
	limit_str:= u.Ctx.Request.Header.Get("limit")
	skip_str:= u.Ctx.Request.Header.Get("skip")

	limit,_:= strconv.ParseInt(limit_str,0,64)
	skip,_:= strconv.ParseInt(skip_str, 0, 64)
	time_start,_:= strconv.ParseInt(time_start_str, 0, 64)
	time_end,_:= strconv.ParseInt(time_end_str, 0, 64)

	print(hid,"-",limit,"-",skip)

	ldevice,code,err:= models.MGetHistoryHomeByTime(hid,int(skip),int(limit),int(time_start),int(time_end))
	if err {
		u.Data["json"]=models.Err(code)
		u.Ctx.ResponseWriter.WriteHeader(code)
		u.ServeJSON()
		return
	}

	print("Len:",len(ldevice))
	u.Data["json"]=models.GetHistoryRsp{Rcode:200,Data:models.HistoryInfo{Total:len(ldevice),Ldevice:ldevice},Message:"Success",Status:true}
	u.ServeJSON()
	return

}

// @Title getHistoryByDid
// @Summary getHistoryByDid
// @Param	auth	header	string	"The token key"
// @Param	hid	header	string	"home_id"
// @Param	did	header	string	"did"
// @Param	limit	header	int32	"count of record"
// @Param	skip	header	int32	"skip"
// @Success 200 {object} models.HistoryInfo
// @Failure 201 fail
// @Failure 401 bad request
// @router /getHistoryDevice [get]
func (this *UserController) GetHistoryByDid() {
	ret_val:= models.Response{Status:true,Rcode:200,Message:"Success"}
	this.Data["json"]=&ret_val

	did:= this.Ctx.Request.Header.Get("did")
	limit_str:= this.Ctx.Request.Header.Get("limit")
	skip_str:= this.Ctx.Request.Header.Get("skip")

	limit,_:= strconv.ParseInt(limit_str,0,64)
	skip,_:= strconv.ParseInt(skip_str, 0, 64)


	ldevice,code,err:= models.MGetHistoryDevice(did,int(skip),int(limit))
	if err {
		ret_val.Rcode=code
		ret_val.Message="Load data fail"
		this.ServeJSON()
		return
	}

	print("Len:",len(ldevice))
	ret_val.Data=models.HistoryInfo{Total:len(ldevice),Ldevice:ldevice}
	this.ServeJSON()
	return

}


// @Title getHistoryByDidByTime
// @Summary getHistoryByDidByTime
// @Param	auth	header	string	"The token key"
// @Param	did	header	string	"did"
// @Param	limit	header	int32	"count of record"
// @Param	skip	header	int32	"skip"
// @Param	time_start	header	int32	"time_start"
// @Param	time_end	header	int32	"time_end"
// @Success 200 {object} models.HistoryInfo
// @Failure 201 fail
// @Failure 401 bad request
// @router /getHistoryDeviceByTime [get]
func (this *UserController) GetHistoryByDidByTime() {
	ret_val:= models.Response{Status:true,Rcode:200,Message:"Success"}
	this.Data["json"]=&ret_val

	did:= this.Ctx.Request.Header.Get("did")
	limit_str:= this.Ctx.Request.Header.Get("limit")
	skip_str:= this.Ctx.Request.Header.Get("skip")
	time_start_str:= this.Ctx.Request.Header.Get("time_start")
	time_end_str:= this.Ctx.Request.Header.Get("time_end")

	limit,_:= strconv.ParseInt(limit_str,0,64)
	skip,_:= strconv.ParseInt(skip_str, 0, 64)
	time_start,_:= strconv.ParseInt(time_start_str, 0, 64)
	time_end,_:= strconv.ParseInt(time_end_str, 0, 64)


	ldevice,code,err:= models.MGetHistoryDeviceByTime(did,int(skip),int(limit),int(time_start),int(time_end))
	if err {
		ret_val.Rcode=code
		ret_val.Message="Load data fail"
		this.ServeJSON()
		return
	}

	print("Len:",len(ldevice))
	ret_val.Data=models.HistoryInfo{Total:len(ldevice),Ldevice:ldevice}
	this.ServeJSON()
	return

}


// @Title UserUpdatePassword
// @Summary Update password user
// @Param	auth	header	string	"The token key"
// @Param	body	body	Models.UpdatePasswordForm	"Update password input"
// @Success 200 success
// @Failure 201 update password fail
// @router /updatePassword [post]
func (c *UserController) PostUpdatePassword() {
	uid:= c.Ctx.Request.Header.Get("uid")

	c.Data["json"]=models.MGetDevice(uid)

	c.ServeJSON()

}

// @Title UserGetProfile
// @Summary Get profile user
// @Param	auth	header	string	"The token key"
// @Success 200 success
// @Failure 400 fail
// @router /getProfile [get]
func (c *UserController) GetProfile() {
	uid:= c.Ctx.Request.Header.Get("uid")

	c.Data["json"]=models.MGetDevice(uid)

	c.ServeJSON()

}

// @Title UserUpdateProfile
// @Summary Get profile user
// @Param	auth	header	string	"The token key"
// @Success 200 success
// @Failure 400 fail
// @router /updateProfile [post]
func (c *UserController) PostUpdateProfile() {
	uid:= c.Ctx.Request.Header.Get("uid")

	c.Data["json"]=models.MGetDevice(uid)

	c.ServeJSON()

}