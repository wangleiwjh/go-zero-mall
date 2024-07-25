package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-mall-myself/service/order/api/internal/config"
	"go-zero-mall-myself/service/order/rpc/orderclient"
	"go-zero-mall-myself/service/product/rpc/productclient"
)

type ServiceContext struct {
	Config     config.Config
	OrderRpc   orderclient.Order
	ProductRpc productclient.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		OrderRpc:   orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
		ProductRpc: productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}
