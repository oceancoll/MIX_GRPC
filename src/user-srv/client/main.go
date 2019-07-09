package main

import (
	"MIX_GRPC/src/share/config"
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"MIX_GRPC/src/share/pb"
)
func runClient(service micro.Service)  {
	usermicro := pb.NewUserServiceClient(config.Namespace+config.ServiceNameUser, service.Client())
	//rsp, err := usermicro.RegistAccount(context.TODO(), &pb.RegistAccountReq{Uname:"xiaohong", Password:"123456", Email:"2@2.com"})
	//rsp, err := usermicro.GetUinfoByEmail(context.TODO(), &pb.GetUinfoByEmailReq{Email:"1@3.com"})
	rsp, err := usermicro.GetAllUinfo(context.TODO(), &pb.GetAllUinfoReq{})
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
