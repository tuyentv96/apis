package model

import (
	"gopkg.in/mgo.v2/bson"
	"apis/models/mongo/db"

	"errors"
)

type DevicePower struct {
	Did     string `json:"did" bson:"did"`
	Hid     string `json:"hid" bson:"hid"`
	Dname   string `json:"dname" bson:"dname"`
	Power  int `json:"power" bson:"power"`
	Type  int `json:"type" bson:"type"`
}

func GetDevicePower(uid string,hid string)  ([]DevicePower,error){
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	result:= []DevicePower{}

	if err := Db.C("power").Find(bson.M{"hid": hid}).Sort("dname").All(&result); err != nil {
		print("Fail")
		return result,errors.New("No home_id found")

	}

	return result,nil
}
