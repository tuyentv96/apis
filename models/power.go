package model

import (
	"gopkg.in/mgo.v2/bson"
	"apis/models/mongo/db"

	"errors"
	"time"
	"fmt"
)

type DevicePower struct {
	Did     string `json:"did" bson:"did"`
	Hid     string `json:"hid" bson:"hid"`
	Dname   string `json:"dname" bson:"dname"`
	Power  int `json:"power" bson:"power"`
	Type  int `json:"type" bson:"type"`
	Date time.Time `json:"date " bson:"date"`
	Time  int64 `json:"time" bson:"time"`
}

type DevicePowerOne struct {
	Power  int `json:"power" bson:"power"`
	Date time.Time `json:"date " bson:"date"`
	Time  int64 `json:"time" bson:"time"`
}

func GetDevicePower(uid string,hid string)  ([]DevicePower,error){
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	year, month, day := time.Now().Date()
	time_now:= time.Date(year, month, day, 0, 0, 0, 0, time.Local)

	result:= []DevicePower{}

	if err := Db.C("power").Find(bson.M{"hid": hid,"date": time_now}).Sort("dname").All(&result); err != nil {
		print("Fail")
		return result,errors.New("No home_id found")

	}

	return result,nil
}

func GetDevicePowerLimit(did string,date_skip int64,date_limit int64)  ([]DevicePowerOne,error){
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	var date_end int64=time.Now().Local().Unix()
	date_end-= (24 * 60 * 60)*date_skip
	y,m,d:=time.Unix(date_end,0).Date()
	time_end:= time.Date(y, m, d, 0, 0, 0, 0, time.Local)

	var next_day int64=time.Now().Local().Unix()
	next_day-= (24 * 60 * 60)*(date_limit+date_skip)
	y1,m1,d1:=time.Unix(next_day,0).Date()
	time_from:= time.Date(y1, m1, d1, 0, 0, 0, 0, time.Local)

	fmt.Printf("%+v","now:",date_end,"from:",time_from)

	result:= []DevicePowerOne{}

	if err := Db.C("power").Find(bson.M{"did": did,
		"date": bson.M{
		"$gt": time_from,
		"$lte": time_end,
		},
	}).Sort("dname").All(&result); err != nil {
		print("Fail")
		return result,errors.New("No home_id found")

	}

	return result,nil
}

func GetDevicePowerByTime(did string,date_start int64,date_end int64)  ([]DevicePowerOne,error){
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	result:= []DevicePowerOne{}
	println("DAte:",date_start,date_end)

	if err := Db.C("power").Find(bson.M{"did": did,
		"time": bson.M{
			"$gte": date_start,
			"$lte": date_end,
		},
	}).Sort("dname").All(&result); err != nil {
		print("Fail")
		return result,errors.New("No home_id found")

	}

	return result,nil
}


func GetHomePowerLimit(hid string,date_skip int64,date_limit int64)  ([]DevicePowerOne,error){
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	var date_end int64=time.Now().Local().Unix()
	date_end-= (24 * 60 * 60)*date_skip
	y,m,d:=time.Unix(date_end,0).Date()
	time_end:= time.Date(y, m, d, 0, 0, 0, 0, time.Local)

	var next_day int64=time.Now().Local().Unix()
	next_day-= (24 * 60 * 60)*(date_limit+date_skip)
	y1,m1,d1:=time.Unix(next_day,0).Date()
	time_from:= time.Date(y1, m1, d1, 0, 0, 0, 0, time.Local)

	fmt.Printf("%+v","now:",date_end,"from:",time_from)

	result:= []DevicePowerOne{}

	if err := Db.C("power").Find(bson.M{"hid": hid,
		"date": bson.M{
			"$gt": time_from,
			"$lte": time_end,
		},
	}).Sort("dname").All(&result); err != nil {
		print("Fail")
		return result,errors.New("No home_id found")

	}

	return result,nil
}


