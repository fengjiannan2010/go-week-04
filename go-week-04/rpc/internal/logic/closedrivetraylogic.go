package logic

import (
	"context"
	"oburn/rpc/client/odrive"

	"oburn/rpc/internal/svc"
	"oburn/rpc/oburn"

	"github.com/tal-tech/go-zero/core/logx"
)

type CloseDriveTrayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCloseDriveTrayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CloseDriveTrayLogic {
	return &CloseDriveTrayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CloseDriveTrayLogic) CloseDriveTray(in *oburn.EmptyRequest) (*oburn.CloseDriveTrayResponse, error) {
	// todo: add your logic here and delete this line
	response, err := l.svcCtx.Odrive.CloseDriveTray(l.ctx, &odrive.EmptyRequest{})
	responseMessage := response.GetMessage()
	return &oburn.CloseDriveTrayResponse{
		Message: &oburn.Message{
			IsSuccess: responseMessage.GetIsSuccess(),
			Code:      responseMessage.GetCode(),
			Message:   responseMessage.GetMessage(),
		},
	}, err
}
