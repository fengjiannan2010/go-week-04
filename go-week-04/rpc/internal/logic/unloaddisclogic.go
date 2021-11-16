package logic

import (
	"context"
	"github.com/tal-tech/go-zero/core/logx"
	"oburn/data"
	"oburn/rpc/client/odrive"
	"oburn/rpc/internal/svc"
	"oburn/rpc/oburn"
)

type UnLoadDiscLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnLoadDiscLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnLoadDiscLogic {
	return &UnLoadDiscLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UnLoadDiscLogic) UnLoadDisc(in *oburn.EmptyRequest) (*oburn.UnLoadDiscResponse, error) {
	// todo: add your logic here and delete this line
	configResponse, err := l.svcCtx.Odrive.GetBurnConfig(l.ctx, &odrive.EmptyRequest{})
	responseMessage := configResponse.GetMessage()
	if !responseMessage.GetIsSuccess() {
		return &oburn.UnLoadDiscResponse{
			Message: &oburn.Message{
				IsSuccess: responseMessage.GetIsSuccess(),
				Code:      responseMessage.GetCode(),
				Message:   responseMessage.GetMessage(),
			},
		}, err
	}
	//文件系统类型 FsType 0:NotSet,1:ImageBurn,2:StreamBurn,3:UdfBurn,6:ExtImage,7:UnKnown
	if configResponse.GetFsType() == data.UdfBurn.Int32() {
		unloadResponse, err := l.svcCtx.Odrive.UnLoadUdfDisc(l.ctx, &odrive.EmptyRequest{})
		responseMessage = unloadResponse.GetMessage()
		return &oburn.UnLoadDiscResponse{
			Message: &oburn.Message{
				IsSuccess: responseMessage.GetIsSuccess(),
				Code:      responseMessage.GetCode(),
				Message:   responseMessage.GetMessage(),
			},
		}, err
	}

	unloadResponse, err := l.svcCtx.Odrive.UnLoadDisc(l.ctx, &odrive.EmptyRequest{})
	responseMessage = unloadResponse.GetMessage()
	return &oburn.UnLoadDiscResponse{
		Message: &oburn.Message{
			IsSuccess: responseMessage.GetIsSuccess(),
			Code:      responseMessage.GetCode(),
			Message:   responseMessage.GetMessage(),
		},
	}, err
}
