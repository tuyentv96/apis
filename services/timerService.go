package services

import (
	"time"
	"github.com/emirpasic/gods/lists/arraylist"
	"fmt"
	mqtt "apis/models/mqtt"
	"encoding/json"
	model "apis/models"
	cron "github.com/robfig/cron"
	"strings"
	"strconv"
	"errors"
)

type Mcontrols struct {
	Uid string `json:"uid" bson:"uid" form:"uid"`
	Hid string `json:"hid" bson:"hid" form:"hid"`
	Devices []model.Device `json:"devices" bson:"devices" form:"devices"`
}

type DeviceInfo struct{
	Hid string
	Did string
	Status int
	Uid string
	Time int64
}

type TimerData struct{
	Devices []model.Device `json:"devices" bson:"devices" form:"devices"`
	Hid string `json:"hid" bson:"hid" form:"hid"`
	Uid string `json:"uid" bson:"uid" form:"uid"`
	Time int64 `json:"time" bson:"time" form:"time"`
}

type CronData struct{
	Devices []model.Device `json:"devices" bson:"devices" form:"devices"`
	Hid string `json:"hid" bson:"hid" form:"hid"`
	Uid string `json:"uid" bson:"uid" form:"uid"`
	Time string `json:"time" bson:"time" form:"time"`
	Date string `json:"date" bson:"date" form:"date"`
}

type AlarmChan struct{
	Data TimerData
	Chan chan string
	ChanId string
}

type CronChan struct{
	Data CronData
	Chan chan string
	ChanId string
}

var chans = arraylist.New()
var cron_chans = arraylist.New()

func Download(out chan string, urls string) {

	go func() {
		time.Sleep(time.Second * 6)
		out <- "Sleep Done"
	}()
}

func Process(data AlarmChan) {
	timeout:=data.Data.Time-time.Now().Unix()
	if timeout<0{
		timeout=0
	}
	print("Timeout to:",timeout)
	select {
	case <- data.Chan:
		print("Stop timer\n")
		RemoveChan(FindChanById(data.ChanId))
		print("Chans size:",chans.Size())
		return
		// Do something with html
	case <- time.After(time.Second * time.Duration(timeout) ):
		print("time out:")
		RemoveChan(FindChanById(data.ChanId))
		print("Chans size:",chans.Size())
		model.UpdateStatusTimer(data.ChanId)
		payl,_:=json.Marshal(data.Data)
		mqtt.MqttPublishMessage("Timer/MCONTROLS",payl)
		return

		// Timed out after 10 seconds!
	}
}

func DoStuff(list_dev []model.Device,hid string,uid string,timeout int64,timer_id string) string{

	// Make a channel
	dev:=TimerData{Hid:hid,Devices:list_dev,Uid:uid,Time:timeout}
	pipe := make(chan string)
	alarm:= AlarmChan{dev,pipe,timer_id}
	fmt.Printf("%+v\n",alarm)
	chans.Add(alarm)
	// Kick off goroutines to download the URLs
	//Download(pipe, "WTF")

	// Process them, with a timeout
	go Process(alarm)
	return alarm.ChanId
}

func FindChanById(id string)  int{
	foundIndex, foundValue := chans.Find(func(index int, value interface{}) bool {
		return value.(AlarmChan).ChanId == id
	})
	if foundValue!=nil {
		fmt.Print(foundIndex)
		return foundIndex
	}

	return -1
}


func RemoveChan(index int)  {
	chans.Remove(index)
}

func RemoveChanByTimerId(timer_id string)  bool{
	foundIndex, foundValue := chans.Find(func(index int, value interface{}) bool {
		return value.(AlarmChan).ChanId == timer_id
	})
	if foundValue!=nil {
		fmt.Print(" Delte Timer_id completed",foundIndex)
		foundValue.(AlarmChan).Chan<-"stop"
		return false
	}
	fmt.Print(" Timer_id not found")
	return true

}

func CronDoStuff(list_dev []model.Device,hid string,uid string,timeout string,date string,timer_id string) string{

	// Make a channel
	dev:=CronData{Hid:hid,Devices:list_dev,Uid:uid,Time:timeout,Date:date}
	pipe := make(chan string)
	alarm:= CronChan{dev,pipe,timer_id}
	fmt.Printf("%+v\n",date)
	cron_chans.Add(alarm)
	// Kick off goroutines to download the URLs


	// Process them, with a timeout
	go CronGoProcess(alarm)
	return alarm.ChanId
}


func CronDoJob(data CronChan) {
	println("Crontab:",data.ChanId)
	payl,_:=json.Marshal(data.Data)
	mqtt.MqttPublishMessage("Timer/MCONTROLS",payl)
}

func ParseTimeToHourMin(data string)  (int,int,error){
	i:=strings.Index(data,":")

	min,min_err:= strconv.ParseInt(data[i+1:],0,32)
	if min_err!=nil{
		return 0,0,errors.New("Parse min error")
	}
	hour,hour_err:= strconv.ParseInt(data[:i],0,32)
	if hour_err!=nil{
		return 0,0,errors.New("Parse hour error")
	}

	return int(hour),int(min),nil

}

func CronGoProcess(data CronChan) {
	c := cron.New()

	hour,min,err:=ParseTimeToHourMin(data.Data.Time)

	if err!=nil{
		return
	}

	sche:= fmt.Sprintf("0 %d %d * * %s",min,hour,data.Data.Date)
	fmt.Print("Cron: ",sche)
	c.AddFunc(sche, func() { go CronDoJob(data) })
	c.Start()
	select {
	case <-data.Chan:
		print("Stop timer\n")
		c.Stop()
		RemoveChan(FindChanById(data.ChanId))
		print("Chans size:", chans.Size())
		return
		// Do something with html
	}
}

func RemoveCronChan(index int)  {
	cron_chans.Remove(index)
}

func RemoveCronChanByTimerId(timer_id string)  bool{
	foundIndex, foundValue := cron_chans.Find(func(index int, value interface{}) bool {
		return value.(CronChan).ChanId == timer_id
	})
	if foundValue!=nil {
		fmt.Print(" Delte Timer_id completed",foundIndex)
		foundValue.(CronChan).Chan<-"stop"
		return false
	}
	fmt.Print(" Timer_id not found")
	return true

}