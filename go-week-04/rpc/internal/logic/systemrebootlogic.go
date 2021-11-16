package logic

import (
	"context"
	"oburn/rpc/client/odrive"

	"oburn/rpc/internal/svc"
	"oburn/rpc/oburn"

	"github.com/tal-tech/go-zero/core/logx"
)

type SystemRebootLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSystemRebootLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SystemRebootLogic {
	return &SystemRebootLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SystemRebootLogic) SystemReboot(in *oburn.EmptyRequest) (*oburn.SystemRebootResponse, error) {
	// todo: add your logic here and delete this line
	response, err := l.svcCtx.Odrive.SystemReboot(l.ctx, &odrive.EmptyRequest{})
	responseMessage := response.GetMessage()
	return &oburn.SystemRebootResponse{
		Message: &oburn.Message{
			IsSuccess: responseMessage.GetIsSuccess(),
			Code:      responseMessage.GetCode(),
			Message:   responseMessage.GetMessage(),
		},
	}, err
}
