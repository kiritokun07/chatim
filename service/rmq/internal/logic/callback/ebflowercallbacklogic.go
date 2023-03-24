package callback

import (
	"context"

	"chatim/service/rmq/internal/svc"
	"chatim/service/rmq/internal/types"

	"github.com/zeromicro/go-zero/core/jsonx"
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
	message, err := jsonx.Marshal(req)
	if err != nil {
		return err
	}
	_ = l.svcCtx.Producer.PushByTopic(context.TODO(), l.svcCtx.Config.SendTopic.EbflowerTopic, message)
	//TODO 返回给平台的消息体
	return nil
}
