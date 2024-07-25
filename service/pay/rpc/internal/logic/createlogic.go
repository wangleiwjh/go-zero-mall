package logic

import (
	"context"
	"go-zero-mall-myself/service/order/pb/order"
	"go-zero-mall-myself/service/pay/model"
	"go-zero-mall-myself/service/user/pb/user"
	"google.golang.org/grpc/status"

	"go-zero-mall-myself/service/pay/pb/pay"
	"go-zero-mall-myself/service/pay/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *pay.CreateRequest) (*pay.CreateResponse, error) {
	// todo: add your logic here and delete this line
	// 查询用户是否存在
	_, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		Id: in.Uid,
	})
	if err != nil {
		return nil, err
	}

	// 查询订单是否存在
	_, err = l.svcCtx.OrderRpc.Detail(l.ctx, &order.DetailRequest{
		Id: in.Oid,
	})
	if err != nil {
		return nil, err
	}

	// 查询订单是否已经创建支付
	_, err = l.svcCtx.PayModel.FindOneByOid(in.Oid)
	if err == nil {
		return nil, status.Error(100, "订单已创建支付")
	}

	newPay := model.Pay{
		Uid:    in.Uid,
		Oid:    in.Oid,
		Amount: in.Amount,
		Source: 0,
		Status: 0,
	}

	res, err := l.svcCtx.PayModel.Insert(l.ctx, &newPay)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	newPay.Id, err = res.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &pay.CreateResponse{
		Id: newPay.Id,
	}, nil
}
