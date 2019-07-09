package errors

import (
	"MIX_GRPC/src/share/config"
	"github.com/micro/go-micro/errors"
)

const errorCodeSaleSuccess  = 200

var (
	ErrorSaleSuccess = errors.New(
		config.ServiceNameSale,"操作成功",errorCodeSaleSuccess,
	)

	ErrorSaleFailed = errors.New(
		config.ServiceNameSale,"操作异常",errorCodeSaleSuccess,
	)
)