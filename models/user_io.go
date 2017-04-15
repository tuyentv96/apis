package model


type UpdatePasswordForm struct {
	Code string `json:"code" bson:"code" form:"code"`
	OldPwd string `json:"old_pwd" bson:"pwd" form:"old_pwd"`
	NewPwd string `json:"new_pwd" bson:"pwd" form:"new_pwd"`
}

