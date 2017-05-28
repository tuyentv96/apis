package main

import (
	_ "apis/routers"
	"github.com/astaxie/beego"

	"github.com/dgrijalva/jwt-go"
	"net/http"
	"github.com/astaxie/beego/context"
	mqtt "apis/models/mqtt"
	"github.com/astaxie/beego/plugins/cors"
)


func AuthJWT(ctx *context.Context){

	jwtString := ctx.Input.Header("auth")
	token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret123"), nil
	})
	if err == nil && token.Valid {

		return


	} else {
		RequireAuth(ctx.ResponseWriter,ctx.Request)
		return
	}

}



func RequireAuth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(401)
	w.Write([]byte("Permission denied"))
}

func main() {
	mqtt.InitMqtt()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}


	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods: []string{"GET","POST"},
		AllowHeaders:     []string{"Origin", "auth", "Access-Control-Allow-Origin","hid","did","timer_id","limit","skip"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
	}))
	beego.InsertFilter("v1/apis/*",beego.BeforeRouter,AuthJWT)
	beego.Run()
}

