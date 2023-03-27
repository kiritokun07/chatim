package message

import (
	"context"

	"chatim/service/platform/internal/svc"
	"chatim/service/platform/internal/types"

	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

type EbflowerMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEbflowerMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EbflowerMessageLogic {
	return &EbflowerMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EbflowerMessageLogic) EbflowerMessage(req *types.EbUpMsgReq) (resp *types.EbUpMsgResp, err error) {
	ebMsg, err := jsonx.MarshalToString(req)
	logx.Info("platform ebflower收到消息：" + ebMsg)
	return &types.EbUpMsgResp{
		Body: &types.EbMsgRespBody{
			Errno: 0,
			Error: "success",
			Data:  true,
		},
		Cmd:       "resp.im.message.send",
		Sign:      "51BAA29E9CE298241F52985864D23165",
		Source:    "65336",
		Ticket:    "FEBCA99A-967D-EBDC-8588-F530B3E235E7",
		Timestamp: 1452686921,
		Traceid:   "0bc1407316188887096315338e3430",
		Version:   3,
	}, nil
}
