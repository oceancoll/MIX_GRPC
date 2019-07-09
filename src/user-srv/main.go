package main

import (
	"MIX_GRPC/src/user-srv/handler"
	"MIX_GRPC/src/share/config"
	"MIX_GRPC/src/share/utils/log"
	"github.com/micro/go-micro"
	"github.com/micro/cli"
	"github.com/micro/go-micro/server"
	"go.uber.org/zap"
	"MIX_GRPC/src/user-srv/db"
	"MIX_GRPC/src/share/pb"
)

func main()  {
	log.Init("user")
	logger := log.Instance()
	service := micro.NewService(
		micro.Name(config.Namespace+config.ServiceNameUser),
		micro.Version("latest"))
	service.Init(
		micro.Action(func(c *cli.Context) {
			logger.Info("Info", zap.Any("user-srv", "user-srv is start ..."))
			db.Init(config.MysqlDsn)
			pb.RegisterUserServiceHandler(service.Server(), new(handler.UserServiceHandler), server.InternalHandler(true))
		}),
		micro.AfterStop(func() error {
			logger.Info("Info", zap.Any("user-srv", "user-srv is stop ..."))
			return nil
		}),
		micro.AfterStart(func() error {
			return nil
		}),
		)
	//启动service
	if err := service.Run(); err != nil {
		logger.Panic("user-srv服务启动失败 ...")
	}
}
