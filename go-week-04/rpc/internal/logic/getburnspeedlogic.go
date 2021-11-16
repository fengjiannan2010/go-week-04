package logic

import (
	"context"
	"oburn/rpc/client/odrive"

	"oburn/rpc/internal/svc"
	"oburn/rpc/oburn"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetBurnSpeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBurnSpeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBurnSpeedLogic {
	return &GetBurnSpeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBurnSpeedLogic) GetBurnSpeed(in *oburn.GetSpeedRequest) (*oburn.GetSpeedResponse, error) {
	// todo: add your logic here and delete this line
	response, err := l.svcCtx.Odrive.GetBurnSpeedConf(l.ctx, &odrive.GetSpeedConfRequest{DiscType: in.GetDiscType()})
	responseMessage := response.GetMessage()
	return &oburn.GetSpeedResponse{
		Message: &oburn.Message{
			IsSuccess: responseMessage.GetIsSuccess(),
			Code:      responseMessage.GetCode(),
			Message:   responseMessage.GetMessage(),
		},
		DiscType: response.GetDiscType(),
		Speed:    response.GetSpeed(),
	}, err
}
