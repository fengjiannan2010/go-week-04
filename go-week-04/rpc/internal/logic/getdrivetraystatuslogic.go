package logic

import (
	"context"
	"oburn/rpc/client/odrive"

	"oburn/rpc/internal/svc"
	"oburn/rpc/oburn"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetDriveTrayStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDriveTrayStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDriveTrayStatusLogic {
	return &GetDriveTrayStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDriveTrayStatusLogic) GetDriveTrayStatus(in *oburn.EmptyRequest) (*oburn.GetDriveTrayStatusResponse, error) {
	// todo: add your logic here and delete this line
	response, err := l.svcCtx.Odrive.GetDriveTrayStatus(l.ctx, &odrive.EmptyRequest{})
	responseMessage := response.GetMessage()
	return &oburn.GetDriveTrayStatusResponse{
		Message: &oburn.Message{
			IsSuccess: responseMessage.GetIsSuccess(),
			Code:      responseMessage.GetCode(),
			Message:   responseMessage.GetMessage(),
		},
		Status: response.GetStatus(),
	}, err
}
