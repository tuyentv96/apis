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
// @router /getAllDevice [get]
func (u *UserController) GetAllDevice() {
	uid:= GetUidByToken(u.Ctx.Request.Header.Get("auth"))

	if uid=="" {
		u.Ctx.ResponseWriter.WriteHeader(400)
		return
	}

	u.Data["json"]=models.MGetDevice(uid)

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
// @router /getHistory [get]
func (u *UserController) GetHisotory() {
	hid:= u.Ctx.Request.Header.Get("hid")
	limit_str:= u.Ctx.Request.Header.Get("limit")
	skip_str:= u.Ctx.Request.Header.Get("skip")

	limit,_:= strconv.ParseInt(limit_str,0,64)
	skip,_:= strconv.ParseInt(skip_str, 0, 64)

	print(hid,"-",limit,"-",skip)

	ldevice,code,err:= models.MGetHistoryDevice(hid,int(skip),int(limit))
	if err {
		u.Data["json"]=models.Err(code)
		u.Ctx.ResponseWriter.WriteHeader(code)
		u.ServeJSON()
		return
	}

	u.Data["json"]=&models.HistoryInfo{Ldevice:ldevice}
	u.ServeJSON()
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