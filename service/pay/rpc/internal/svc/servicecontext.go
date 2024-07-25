package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-mall-myself/service/order/rpc/orderclient"
	"go-zero-mall-myself/service/pay/model"
	"go-zero-mall-myself/service/pay/rpc/internal/config"
	"go-zero-mall-myself/service/user/rpc/userclient"
)

type ServiceContext struct {
	Config   config.Config
	PayModel model.PayModel
	UserRpc  userclient.User
	OrderRpc orderclient.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:   c,
		PayModel: model.NewPayModel(conn, c.CacheRedis),
		UserRpc:  userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		OrderRpc: orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
	}
}
