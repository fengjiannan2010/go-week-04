package logic

import (
	"context"
	"oburn/data"
	"oburn/rpc/client/odrive"

	"oburn/rpc/internal/svc"
	"oburn/rpc/oburn"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetBurnConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBurnConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBurnConfigLogic {
	return &GetBurnConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBurnConfigLogic) GetBurnConfig(in *oburn.EmptyRequest) (*oburn.GetBurnConfigResponse, error) {
	// todo: add your logic here and delete this line
	response, err := l.svcCtx.Odrive.GetBurnConfig(l.ctx, &odrive.EmptyRequest{})
	responseMessage := response.GetMessage()
	if responseMessage.IsSuccess {
		l.svcCtx.BurnConfig.DiscLabel = response.GetDiscLabel()
		l.svcCtx.BurnConfig.DiscPasswd = response.GetDiscPasswd()
		l.svcCtx.BurnConfig.FsType = data.DiscFsType(response.GetFsType())
	}
	return &oburn.GetBurnConfigResponse{
		Message: &oburn.Message{
			IsSuccess: responseMessage.GetIsSuccess(),
			Code:      responseMessage.GetCode(),
			Message:   responseMessage.GetMessage(),
		},
		DiscLabel:  response.GetDiscLabel(),
		DiscPasswd: response.GetDiscPasswd(),
		FsType:     response.GetFsType(),
	}, err
}
