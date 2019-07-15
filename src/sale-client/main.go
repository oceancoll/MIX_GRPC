package main

import (
	"MIX_GRPC/src/share/config"
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"MIX_GRPC/src/share/pb"
)
func runClient(service micro.Service)  {
	salemicro := pb.NewSaleServiceClient(config.Namespace+config.ServiceNameSale, service.Client())
	//rsp, err := salemicro.AddBuyitem(context.TODO(), &pb.AddSaleitemReq{Email:"1@1.com", Itemname:"banana", Price:4.2})
	//rsp, err := salemicro.GetBuyitemsByEmail(context.TODO(), &pb.GetBuyitemsByEmailReq{Email:"1@1.com"})
	rsp, err := salemicro.GetAllBuyitems(context.TODO(), &pb.GetAllBuyitemsReq{})
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(rsp)
}

func main()  {
	service := micro.NewService()
	service.Init()
	runClient(service)
}