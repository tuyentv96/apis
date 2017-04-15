package model

import (
	"gopkg.in/mgo.v2/bson"
	"apis/models/mongo/db"
	"fmt"
	"crypto/sha256"
)

type User struct {
		Uid     string `json:"uid" bson:"uid" form:"uid"  `
		Pwd     string `json:"pwd,omitempty" bson:"pwd" form:"pwd"  `
}

type Object struct {
	ObjectId   string
	Score      int64
	PlayerName string
}


type Userpsmdevice struct {
	Lhome []struct {
		Hid  string `json:"hid"`
		Type int    `json:"type"`
	} `json:"lhome"`

	UID   string `json:"uid"`
	Uname string `json:"uname"`
}

func (u *User) FindByUid(uid string) (int,bool){
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	if err := Db.C("users").Find(bson.M{"uid": uid}).One(&u); err != nil {
		print("Fail clgt")
		return 104,true
	}

	if u.Uid=="" {
		return 104,true
	}

	return 200,false

}

func (u *User)	CheckPwd(pwd string)  (int,bool){

	encryted_pwd:= fmt.Sprintf("%x",sha256.Sum256([]byte(pwd)))

	if u.Pwd!=encryted_pwd {
		return 410,true
	}

	return 200,false

}
func (u *User) ClearPass() {
	u.Pwd=""
}

