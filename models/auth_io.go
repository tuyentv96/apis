package model

// LoginInfo definiton.
type LoginForm struct {
	Uid string `json:"uid" bson:"uid" form:"uid"`
	Pwd string `json:"pwd" bson:"pwd" form:"pwd"`
}
type LoginRsp struct {
	Message string `json:"message"`
	UserInfo UserInfo `json:"user"`
	Token string `json:"token"`
}

type UserInfo struct {
	Uid string `json:"uid"`
	//UName string `json:"uname"`
}
type GetCodeForm struct {
	Uid string `json:"uid" bson:"uid" form:"uid"`
}

type VerifyCodeForm struct {
	Uid string `json:"uid" bson:"uid" form:"uid"`
	Code string `json:"code" bson:"code" form:"code"`
}

type Message struct {
	Message string `json:"message"`
}

