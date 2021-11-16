package logic

import (
	"context"
	"oburn/rpc/client/odrive"

	"oburn/rpc/internal/svc"
	"oburn/rpc/oburn"

	"github.com/tal-tech/go-zero/core/logx"
)

type SetBurnSpeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetBurnSpeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetBurnSpeedLogic {
	return &SetBurnSpeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetBurnSpeedLogic) SetBurnSpeed(in *oburn.SetSpeedRequest) (*oburn.SetSpeedResponse, error) {
	// todo: add your logic here and delete this line
	response, err := l.svcCtx.Odrive.SetBurnSpeedConf(l.ctx, &odrive.SetSpeedConfRequest{DiscType: in.GetDiscType(), Speed: in.GetSpeed()})
	responseMessage := response.GetMessage()
	return &oburn.SetSpeedResponse{
		Message: &oburn.Message{
			IsSuccess: responseMessage.GetIsSuccess(),
			Code:      responseMessage.GetCode(),
			Message:   responseMessage.GetMessage(),
		},
	}, err
}
