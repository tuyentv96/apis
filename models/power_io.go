package model

type GetDevicePowerOutput struct {
	Uid string `json:"uid" bson:"uid" form:"uid"`
	Hid string `json:"hid" bson:"hid" form:"hid"`
	Did string `json:"did" bson:"did" form:"did"`
	Dname   string `json:"dname" bson:"dname"`
	Status  int `json:"status" bson:"status"`
	Type  int `json:"type" bson:"type"`
	Total int `json:"total" bson:"total" form:"total"`
	Devices []DevicePower `json:"power_data" bson:"power_data" form:"power_data"`
}

type GetDevicePowerOneOutput struct {
	Uid string `json:"uid" bson:"uid" form:"uid"`
	Hid string `json:"hid" bson:"hid" form:"hid"`
	Did string `json:"did" bson:"did" form:"did"`
	Dname   string `json:"dname" bson:"dname"`
	Status  int `json:"status" bson:"status"`
	Type  int `json:"type" bson:"type"`
	Total int `json:"total" bson:"total" form:"total"`
	Devices []DevicePowerOne `json:"power_data" bson:"power_data" form:"power_data"`
}

