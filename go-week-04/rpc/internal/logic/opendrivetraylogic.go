package logic

import (
	"context"
	"oburn/rpc/client/odrive"

	"oburn/rpc/internal/svc"
	"oburn/rpc/oburn"

	"github.com/tal-tech/go-zero/core/logx"
)

type OpenDriveTrayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOpenDriveTrayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OpenDriveTrayLogic {
	return &OpenDriveTrayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OpenDriveTrayLogic) OpenDriveTray(in *oburn.OpenDriveTrayRequest) (*oburn.OpenDriveTrayResponse, error) {
	// todo: add your logic here and delete this line
	response, err := l.svcCtx.Odrive.OpenDriveTray(l.ctx, &odrive.OpenDriveTrayRequest{
		DisableUnlock: in.GetDisableUnlock(),
	})
	responseMessage := response.GetMessage()
	return &oburn.OpenDriveTrayResponse{
		Message: &oburn.Message{
			IsSuccess: responseMessage.GetIsSuccess(),
			Code:      responseMessage.GetCode(),
			Message:   responseMessage.GetMessage(),
		},
	}, err
}
