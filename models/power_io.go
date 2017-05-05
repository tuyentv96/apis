package model

type GetDevicePowerOutput struct {
	Uid string `json:"uid" bson:"uid" form:"uid"`
	Hid string `json:"hid" bson:"hid" form:"hid"`
	Devices []DevicePower `json:"devices" bson:"devices" form:"devices"`
}

