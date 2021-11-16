package logic

import (
	"context"
	"oburn/rpc/client/odrive"

	"oburn/rpc/internal/svc"
	"oburn/rpc/oburn"

	"github.com/tal-tech/go-zero/core/logx"
)

type CompleteDiscLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCompleteDiscLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CompleteDiscLogic {
	return &CompleteDiscLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CompleteDiscLogic) CompleteDisc(in *oburn.EmptyRequest) (*oburn.CompleteDiscResponse, error) {
	// todo: add your logic here and delete this line
	response, err := l.svcCtx.Odrive.CompleteDisc(l.ctx, &odrive.EmptyRequest{})
	responseMessage := response.GetMessage()
	return &oburn.CompleteDiscResponse{
		Message: &oburn.Message{
			IsSuccess: responseMessage.GetIsSuccess(),
			Code:      responseMessage.GetCode(),
			Message:   responseMessage.GetMessage(),
		},
	}, err
}
