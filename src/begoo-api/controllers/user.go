package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/micro/go-micro"
	"MIX_GRPC/src/share/pb"
	"MIX_GRPC/src/share/config"
	"context"
	"github.com/tidwall/gjson"
	_ "github.com/json-iterator/go"
	"MIX_GRPC/src/share/utils/common"
)

type UserController struct {
	beego.Controller
}


func InitUser() pb.UserServiceClient {
	cli := micro.NewService()
	cli.Init()
	UserClient := pb.NewUserServiceClient(config.Namespace+config.ServiceNameUser, cli.Client())
	return UserClient
}


func (self *UserController) Regist() {
	uname := self.GetString("uname")
	password := self.GetString("password")
	email := self.GetString("email")
	UReq := &pb.RegistAccountReq{
		Uname:uname,
		Password:password,
		Email:email,
	}
	UserClient := InitUser()
	_, err := UserClient.RegistAccount(context.TODO(), UReq)
	resuinfo := make(map[string]interface{})
	if err != nil{
		result := &common.JSONStruct{200, resuinfo, gjson.Get(err.Error(), "detail" ).Str}
		self.Data["json"] = result
	} else{
		result := &common.JSONStruct{200, resuinfo, "ok"}
		self.Data["json"] = result
	}
	self.ServeJSON()
}

func (self *UserController) Show() {
	email := self.GetString("email")
	UReq := &pb.GetUinfoByEmailReq{
		Email:email,
	}
	UserClient := InitUser()
	uinfo, err := UserClient.GetUinfoByEmail(context.TODO(), UReq)
	fmt.Println(uinfo)
	resuinfo := make(map[string]interface{})
	if err != nil{
		fmt.Println(err)
		result := &common.JSONStruct{500, resuinfo, "intererror"}
		self.Data["json"] = result
		fmt.Println(result)
	} else if uinfo == nil{
		result := &common.JSONStruct{200, resuinfo, "ok"}
		self.Data["json"] = result
		fmt.Println(result)
	} else {
		resuinfo["uid"] = uinfo.Id
		resuinfo["uname"] = uinfo.Uname
		resuinfo["email"] = uinfo.Email
		resuinfo["crtime"] = uinfo.Crtime
		result := &common.JSONStruct{2001, resuinfo, "ok"}
		self.Data["json"] = result
		fmt.Println(result)
	}
	self.ServeJSON()
}

func (self *UserController) GetAll() {
	UReq := &pb.GetAllUinfoReq{}
	UserClient := InitUser()
	uinfos, err := UserClient.GetAllUinfo(context.TODO(), UReq)
	resuinfo := make([]map[string]interface{},0)
	if err != nil{
		fmt.Println(err)
		result := &common.ListJSONStruct{500, resuinfo, "intererror"}
		self.Data["json"] = result
	} else {
		//var json_iterator= jsoniter.ConfigCompatibleWithStandardLibrary
		//b, err := json_iterator.Marshal(uinfos.Alluinfo)
		for _, r:= range uinfos.Alluinfo{
			info := make(map[string]interface{})
			info["uid"] = r.Id
			info["uname"] = r.Uname
			info["email"] = r.Email
			info["crtime"] = r.Crtime
			resuinfo = append(resuinfo, info)
		}
		result := &common.ListJSONStruct{200, resuinfo, "ok"}
		self.Data["json"] = result
	}
	self.ServeJSON()
}