package logic

import (
	"context"
	"oburn/rpc/client/odrive"

	"oburn/rpc/internal/svc"
	"oburn/rpc/oburn"

	"github.com/tal-tech/go-zero/core/logx"
)

type SystemPoweroffLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSystemPoweroffLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SystemPoweroffLogic {
	return &SystemPoweroffLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SystemPoweroffLogic) SystemPoweroff(in *oburn.EmptyRequest) (*oburn.SystemPoweroffResponse, error) {
	// todo: add your logic here and delete this line
	response, err := l.svcCtx.Odrive.SystemPoweroff(l.ctx, &odrive.EmptyRequest{})
	responseMessage := response.GetMessage()
	return &oburn.SystemPoweroffResponse{
		Message: &oburn.Message{
			IsSuccess: responseMessage.GetIsSuccess(),
			Code:      responseMessage.GetCode(),
			Message:   responseMessage.GetMessage(),
		},
	}, err
}
