package callback

import (
	"context"

	"chatim/service/rmq/internal/svc"
	"chatim/service/rmq/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EbflowerCallbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEbflowerCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EbflowerCallbackLogic {
	return &EbflowerCallbackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EbflowerCallbackLogic) EbflowerCallback(req *types.EbMsg) error {
	// todo: add your logic here and delete this line

	return nil
}
