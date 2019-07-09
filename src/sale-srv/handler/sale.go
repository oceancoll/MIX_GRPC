package handler

import (
	"MIX_GRPC/src/sale-srv/db"
	"MIX_GRPC/src/share/config"
	"MIX_GRPC/src/share/pb"
	"context"
	"fmt"
	"github.com/micro/go-micro"
	_ "github.com/micro/go-micro/service/grpc"
	"go.uber.org/zap"
	"MIX_GRPC/src/share/errors"
)


type SaleServiceHandler struct {
	//UserClient pb.UserServiceClient
	logger *zap.Logger
}

func InitUser() pb.UserServiceClient {
	cli := micro.NewService()
	cli.Init()
	UserClient := pb.NewUserServiceClient(config.Namespace+config.ServiceNameUser, cli.Client())
	return UserClient
}
//新增用户购买信息
func (s *SaleServiceHandler)AddBuyitem(ctx context.Context, req *pb.AddSaleitemReq, rsp *pb.AddSaleitemRsp) error {
	email := req.Email
	itemname := req.Itemname
	price := req.Price
	UReq := &pb.GetUinfoByEmailReq{
		Email:email,
	}
	UserClient := InitUser()
	uinfo, err := UserClient.GetUinfoByEmail(context.TODO(), UReq)
	if err != nil{
		fmt.Println(err)
		return errors.ErrorSaleFailed
	}
	if uinfo == nil{
		return errors.ErrorUserNotExists
	}
	uid := uinfo.Id
	err = db.InsertSaleItem(uid, itemname, price)
	if err != nil{
		return errors.ErrorSaleFailed
	}
	return nil
}


//通过id查询某用户购买信息
func (s *SaleServiceHandler)GetBuyitemsByEmail(ctx context.Context, req *pb.GetBuyitemsByEmailReq, rsp *pb.GetBuyitemsByEmailRsp) error {
	email := req.Email
	UReq := &pb.GetUinfoByEmailReq{
		Email:email,
	}
	UserClient := InitUser()
	uinfo, err := UserClient.GetUinfoByEmail(context.TODO(), UReq)
	if err != nil{
		fmt.Println(err)
		return errors.ErrorSaleFailed
	}
	if uinfo == nil{
		return errors.ErrorUserNotExists
	}
	uid := uinfo.Id
	items, err := db.GetBuyitemsByUid(uid)
	if err != nil{
		fmt.Println(err)
		return errors.ErrorSaleFailed
	}
	finalitems := []*pb.Item{}
	for _, row := range items{
		item := pb.Item{Itemname:row.Itemname, Price: row.Price, Crtime:row.CrTime}
		finalitems = append(finalitems, &item)
	}
	rsp.Uname = uinfo.Uname
	rsp.Items = finalitems
	return nil
}

//获取所有购买信息
func (s *SaleServiceHandler)GetAllBuyitems(ctx context.Context, req *pb.GetAllBuyitemsReq, rsp *pb.GetAllBuyitemsRsp) error {
	items, err := db.SelectAllItems()
	if err != nil{
		return errors.ErrorSaleFailed
	}
	for _, row := range items{
		item := pb.BuyItem{Uname:row.Uname, Itemname:row.Itemname, Price:row.Price, Crtime:row.CrTime}
		rsp.Buyitems = append(rsp.Buyitems, &item)
	}
	return nil
}
