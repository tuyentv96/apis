package model

type SetTimerInput struct {
	Uid string `json:"uid" bson:"uid" form:"uid"`
	Hid string `json:"hid" bson:"hid" form:"hid"`
	Devices []DeviceInfo `json:"devices" bson:"devices" form:"devices"`
	Time int64 `json:"time" bson:"time" form:"time"`
}

type Mcontrols struct {
	Uid string `json:"uid"`
	Hid string `json:"hid"`
	Device []DeviceInfo  `json:"devices"`
}

type DeviceInfo struct {
	Did string `json:"did" bson:"did" form:"did"`
	Status int `json:"status" bson:"status" form:"status"`
}

type CronTimerInput struct {
	Uid string `json:"uid" bson:"uid" form:"uid"`
	Hid string `json:"hid" bson:"hid" form:"hid"`
	Devices []DeviceInfo `json:"devices" bson:"devices" form:"devices"`
	Time string `json:"time" bson:"time" form:"time"`
	Date string `json:"date" bson:"date" form:"date"`
}
