package main

import (
	"MIX_GRPC/src/sale-srv/handler"
	"MIX_GRPC/src/share/config"
	"MIX_GRPC/src/share/utils/log"
	"github.com/micro/go-micro"
	"github.com/micro/cli"
	"github.com/micro/go-micro/server"
	"go.uber.org/zap"
	"MIX_GRPC/src/sale-srv/db"
	"MIX_GRPC/src/share/pb"
)

func main()  {
	log.Init(config.ServiceNameSale)
	logger := log.Instance()
	service := micro.NewService(
		micro.Name(config.Namespace+config.ServiceNameSale),
		micro.Version("latest"))
	service.Init(
		micro.Action(func(c *cli.Context) {
			logger.Info("Info", zap.Any("sale-srv", "sale-srv is start ..."))
			db.Init(config.MysqlDsn)
			pb.RegisterSaleServiceHandler(service.Server(), new(handler.SaleServiceHandler), server.InternalHandler(true))
		}),
		micro.AfterStop(func() error {
			logger.Info("Info", zap.Any("sale-srv", "sale-srv is stop ..."))
			return nil
		}),
		micro.AfterStart(func() error {
			return nil
		}),
	)
	//启动service
	if err := service.Run(); err != nil {
		logger.Panic("sale-srv服务启动失败 ...")
	}
}
