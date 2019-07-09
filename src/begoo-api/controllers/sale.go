package controllers

import (
	"MIX_GRPC/src/share/pb"
	"MIX_GRPC/src/share/utils/common"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/micro/go-micro"
	"github.com/tidwall/gjson"
	"strconv"
	"MIX_GRPC/src/share/config"
	"context"
)

type SaleController struct {
	beego.Controller
}

func InitSale() pb.SaleServiceClient {
	cli := micro.NewService()
	cli.Init()
	SaleClient := pb.NewSaleServiceClient(config.Namespace+config.ServiceNameSale, cli.Client())
	return SaleClient
}

func (self *SaleController) Addbuyitem() {
	email := self.GetString("email")
	itemname := self.GetString("itemname")
	price, _ := strconv.ParseFloat(self.GetString("price"), 32)
	SReq := &pb.AddSaleitemReq{
		Email:email,
		Itemname:itemname,
		Price: float32(price),
	}
	SaleClient := InitSale()
	_, err := SaleClient.AddBuyitem(context.TODO(), SReq)
	ressale := make(map[string]interface{})
	if err != nil{
		result := &common.JSONStruct{200, ressale, gjson.Get(err.Error(), "detail" ).Str}
		self.Data["json"] = result
	} else{
		result := &common.JSONStruct{200, ressale, "ok"}
		self.Data["json"] = result
	}
	self.ServeJSON()
}

func (self *SaleController) GetBuyitemsByEmail() {
	email := self.GetString("email")
	SReq := &pb.GetBuyitemsByEmailReq{
		Email:email,
	}
	SaleClient := InitSale()
	sinfos, err := SaleClient.GetBuyitemsByEmail(context.TODO(), SReq)
	ressinfo := make(map[string]interface{})
	if err != nil{
		result := &common.JSONStruct{200, ressinfo, gjson.Get(err.Error(), "detail" ).Str}
		self.Data["json"] = result
	} else {
		ressinfo["name"] = sinfos.Uname
		ressinfo["items"] = sinfos.Items
		result := &common.JSONStruct{200, ressinfo, "ok"}
		self.Data["json"] = result
	}
	self.ServeJSON()
}

func (self *SaleController) GetAllBuyitems()  {
	SReq := &pb.GetAllBuyitemsReq{
	}
	SaleClient := InitSale()
	sinfos, err := SaleClient.GetAllBuyitems(context.TODO(), SReq)
	ressinfo := make([]map[string]interface{},0)
	if err != nil{
		fmt.Println(err)
		result := &common.ListJSONStruct{500, ressinfo, "intererror"}
		self.Data["json"] = result
	} else {
		for _, r:= range sinfos.Buyitems{
			info := make(map[string]interface{})
			info["uname"] = r.Uname
			info["itemname"] = r.Itemname
			info["price"] = r.Price
			info["crtime"] = r.Crtime
			ressinfo = append(ressinfo, info)
		}
		result := &common.ListJSONStruct{200, ressinfo, "ok"}
		self.Data["json"] = result
	}
	self.ServeJSON()

}