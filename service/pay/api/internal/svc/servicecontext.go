package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-mall-myself/service/pay/api/internal/config"
	"go-zero-mall-myself/service/pay/rpc/payclient"
)

type ServiceContext struct {
	Config config.Config
	PayRpc payclient.Pay
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		PayRpc: payclient.NewPay(zrpc.MustNewClient(c.PayRpc)),
	}
}
