package controllers

import (
	"github.com/astaxie/beego"
	//models "apis/models"
)

// Operations about user
type AdminController struct {
	beego.Controller
}

// @Title AddDevice
// @Summary Add device for user
// @Param	auth	header	string	"The api key"
// @Param	body		body 	models.AddDeviceInput	true		"The body content"
// @Success 200 success
// @Failure 403 code not found
// @router /addDevice [post]
func (o *AdminController) PostAddDevice() {
	o.ServeJSON()
}
