package logic

import (
	"context"
	"go-zero-mall-myself/service/pay/pb/pay"

	"go-zero-mall-myself/service/pay/api/internal/svc"
	"go-zero-mall-myself/service/pay/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateRequest) (resp *types.CreateResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.PayRpc.Create(l.ctx, &pay.CreateRequest{
		Uid:    req.Uid,
		Oid:    req.Oid,
		Amount: req.Amount,
	})
	if err != nil {
		return nil, err
	}

	return &types.CreateResponse{
		Id: res.Id,
	}, nil
}
