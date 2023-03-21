package ws

import (
	"context"

	"chatim/service/platform/internal/svc"
	"chatim/service/platform/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WsLogic {
	return &WsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WsLogic) Ws(req *types.WsReq) error {
	// todo: add your logic here and delete this line

	return nil
}
