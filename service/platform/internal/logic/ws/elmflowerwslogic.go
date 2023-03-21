package ws

import (
	"context"

	"chatim/service/platform/internal/svc"
	"chatim/service/platform/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ElmflowerWsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewElmflowerWsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ElmflowerWsLogic {
	return &ElmflowerWsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ElmflowerWsLogic) ElmflowerWs(req *types.WsReq) error {
	// todo: add your logic here and delete this line

	return nil
}
