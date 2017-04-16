package controllers

import (
	models "apis/models"
	"github.com/astaxie/beego"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/astaxie/beego/context"
	"encoding/json"
	"net/http"
)

type jwtCustomClaims struct {
	Uid  string `json:"uid"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}


// Operations about Auth
type AuthController struct {
	beego.Controller
}


// @Title UserLogin
// @Summary User Login
// @Param	uid	formData	string apikey
// @Param	pwd	formData	string apikey
// @Success 200 {object} models.LoginRsp
// @Failure 100 wrong format
// @Failure 410 wrong password
// @router /login [post]
func (c *AuthController) Post() {
	login_form:=models.LoginForm{}
	if err:= c.ParseForm(&login_form);err!=nil {
		c.Ctx.ResponseWriter.WriteHeader(100)
		c.ServeJSON()
		return

	}
	print(login_form.Uid)
	print(login_form.Uid,login_form.Pwd)

	uid:= login_form.Uid
	pwd:= login_form.Pwd

	user:= models.User{}

	if code,err:= user.FindByUid(uid); err{
		// Wrong user
		print(code)
		c.Data["json"]=models.Err(402)
		c.Ctx.ResponseWriter.WriteHeader(402)
		c.ServeJSON()
		return
	}

	if code, err := user.CheckPwd(pwd); err {
		// Wrong password
		c.Data["json"]=models.Err(code)
		c.Ctx.ResponseWriter.WriteHeader(410)
		c.ServeJSON()
		return
	}

	user.ClearPass()

	claims := &jwtCustomClaims{
		user.Uid,
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret123"))
	if err != nil {
		c.ServeJSON()
		return
	}
	user_info:=models.UserInfo{}
	user_info.Uid=user.Uid
	user_info.Token=t

	c.Data["json"]=models.LoginRsp{Rcode:200,UserInfo: user_info,Message: "Login successful",Status:true}
	c.ServeJSON()
	return
}



// @Title TokenVerify
// @Summary Token verfiy
// @Param	auth	header 	string	true	"The token id"
// @Success 200 success
// @Failure 100 wrong format
// @Failure 401 token is expired
// @router /tokenVerify [post]
func (c *AuthController) PostTokenVerify() {
	jwtString := c.Ctx.Request.Header.Get("auth")
	token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret123"), nil
	})

	if err == nil && token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		uid := fmt.Sprintf("%v",claims["uid"])
		fmt.Printf(uid)
		c.Ctx.ResponseWriter.WriteHeader(200)
		c.Data["json"]=models.Message{Rcode:200,Message: "success",Status:true}
		c.ServeJSON()
		return

	} else {
		c.Ctx.ResponseWriter.WriteHeader(401)
		c.Data["json"]=models.Err(401)
		c.ServeJSON()
		return
	}

	return
}


func Auth(c *context.Context) {
	beego.Debug("checking......")
}

func RequireAuth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(401)
	w.Write([]byte("Permission denied\n"))
}

// @Title GetCode
// @Summary Get code
// @Param	body		body 	models.GetCodeForm	true		"The body content"
// @Success 200 success
// @Failure 100 wrong format
// @router /forgetPassword/getCode [post]
func (a *AuthController) PostGetCode() {
	input:= models.GetCodeForm{}
	if err:=json.Unmarshal(a.Ctx.Input.RequestBody, &input);err!=nil {
		a.Data["json"]=models.Err(100)
		a.Ctx.ResponseWriter.WriteHeader(100)
		a.ServeJSON()
	}
	print("uid:",input.Uid,a.GetString("uid"))
	a.Data["json"]=input
	a.ServeJSON()
}


// @Title VerifyCode
// @Summary Verify Code
// @Param	body		body 	models.VerifyCodeForm	true		"The object content"
// @Success 200 success
// @Failure 403 code not found
// @router /forgetPassword/verifyCode [post]
func (o *AuthController) PostVerifyCode() {
	var ob models.LoginForm
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	print("dm",ob.Uid)
	o.ServeJSON()
}

// @Title Logout
// @Summary Logout
// @Param	auth	header 	string	true	"The token id"
// @Success 200 success
// @Failure 401 Token is expired
// @router /logout [post]
func (this *AuthController) PostLogout() {
	this.Data["json"]=models.Message{Rcode:200,Message:"Logout successful",Status:true}
	this.ServeJSON()
}

func GetUidByToken(jwt_string string)  string{
	token, err := jwt.Parse(jwt_string, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret123"), nil
	})

	if err == nil && token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		uid := fmt.Sprintf("%v",claims["uid"])
		return uid
	}
	return ""

}

